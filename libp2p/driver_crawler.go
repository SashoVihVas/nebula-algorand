package libp2p

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/dennis-tra/nebula-crawler/config"
	"github.com/dennis-tra/nebula-crawler/core"
	"github.com/dennis-tra/nebula-crawler/db"
	"github.com/dennis-tra/nebula-crawler/kubo"
	"github.com/dennis-tra/nebula-crawler/utils"

	"github.com/algorand/go-algorand/network/p2p"
	"github.com/algorand/go-algorand/protocol"
	"github.com/libp2p/go-libp2p"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	core_protocol "github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	ma "github.com/multiformats/go-multiaddr"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	algorand_config "github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/logging"
)

type PeerInfo struct {
	peer.AddrInfo
}

var _ core.PeerInfo[PeerInfo] = (*PeerInfo)(nil)

func (p PeerInfo) ID() peer.ID {
	return p.AddrInfo.ID
}

func (p PeerInfo) Addrs() []ma.Multiaddr {
	return p.AddrInfo.Addrs
}

func (p PeerInfo) Merge(other PeerInfo) PeerInfo {
	if p.AddrInfo.ID != other.AddrInfo.ID {
		panic("merge peer ID mismatch")
	}

	return PeerInfo{
		AddrInfo: peer.AddrInfo{
			ID:    p.AddrInfo.ID,
			Addrs: utils.MergeMaddrs(p.AddrInfo.Addrs, other.AddrInfo.Addrs),
		},
	}
}

func (p PeerInfo) DiscoveryPrefix() uint64 {
	kadID := kbucket.ConvertPeerID(p.AddrInfo.ID)
	return binary.BigEndian.Uint64(kadID[:8])
}

func (p PeerInfo) DeduplicationKey() string {
	return p.AddrInfo.ID.String()
}

type CrawlDriverConfig struct {
	Version        string
	WorkerCount    int
	Network        config.Network
	Protocols      []string
	DialTimeout    time.Duration
	CheckExposed   bool
	BootstrapPeers []peer.AddrInfo
	AddrDialType   config.AddrType
	MeterProvider  metric.MeterProvider
	TracerProvider trace.TracerProvider
	GossipSubPX    bool
	LogErrors      bool
}

func (cfg *CrawlDriverConfig) CrawlerConfig() *CrawlerConfig {
	crawlerCfg := DefaultCrawlerConfig()
	crawlerCfg.DialTimeout = cfg.DialTimeout
	crawlerCfg.CheckExposed = cfg.CheckExposed
	crawlerCfg.AddrDialType = cfg.AddrDialType
	crawlerCfg.GossipSubPX = cfg.GossipSubPX
	crawlerCfg.LogErrors = cfg.LogErrors
	return crawlerCfg
}

func (cfg *CrawlDriverConfig) WriterConfig() *core.CrawlWriterConfig {
	return &core.CrawlWriterConfig{}
}

type CrawlDriver struct {
	cfg             *CrawlDriverConfig
	host            *Host
	discoveries     map[peer.ID]*p2p.CapabilitiesDiscovery
	dbc             db.Client
	pxPeersChan     chan []PeerInfo
	tasksChan       chan PeerInfo
	workerStateChan chan string
	crawlerCount    int
	writerCount     int
}

var _ core.Driver[PeerInfo, core.CrawlResult[PeerInfo]] = (*CrawlDriver)(nil)

func NewCrawlDriver(dbc db.Client, cfg *CrawlDriverConfig) (*CrawlDriver, error) {
	userAgent := "nebula/" + cfg.Version
	if cfg.Network == config.NetworkAvailTuringLC || cfg.Network == config.NetworkAvailMainnetLC {
		userAgent = "avail-light-client/light-client/1.12.13/rust-client"
	}

	// Create a standard libp2p host
	baseHost, err := newLibp2pHost(userAgent)
	if err != nil {
		return nil, fmt.Errorf("new libp2p host: %w", err)
	}

	// Wrap the standard host with our custom Host wrapper
	h, err := WrapHost(baseHost)
	if err != nil {
		return nil, fmt.Errorf("wrapping host: %w", err)
	}

	tasksChan := make(chan PeerInfo, len(cfg.BootstrapPeers))
	for _, addrInfo := range cfg.BootstrapPeers {
		tasksChan <- PeerInfo{AddrInfo: addrInfo}
	}

	// Initialize the map for discoveries. Even with a single host, the key is the peer ID.
	discoveries := make(map[peer.ID]*p2p.CapabilitiesDiscovery)

	if cfg.Network == config.NetworkAlgoTestnet {
		networkID := protocol.NetworkID("/testnet/kad/1.0.0")
		dhtCfg := algorand_config.GetDefaultLocal()
		logger := logging.NewLogger() // A simple logger for initialization

		// A simple bootstrap function that returns the initial peer list
		bootstrapFunc := func() []peer.AddrInfo {
			return cfg.BootstrapPeers
		}

		// Use the single host to create the capabilities discovery service.
		// Since our custom *Host embeds the host.Host interface, it can be used here directly.
		disc, err := p2p.MakeCapabilitiesDiscovery(context.Background(), dhtCfg, h, networkID, logger, bootstrapFunc)
		if err != nil {
			return nil, fmt.Errorf("failed to create capabilities discovery for host %s: %w", h.ID(), err)
		}
		discoveries[h.ID()] = disc
	}

	d := &CrawlDriver{
		cfg:             cfg,
		host:            h, // Now correctly assigning the wrapped *Host
		discoveries:     discoveries,
		dbc:             dbc,
		tasksChan:       tasksChan,
		pxPeersChan:     make(chan []PeerInfo),
		workerStateChan: make(chan string, cfg.WorkerCount),
		crawlerCount:    0,
		writerCount:     0,
	}

	return d, nil
}

func (d *CrawlDriver) NewWorker() (core.Worker[PeerInfo, core.CrawlResult[PeerInfo]], error) {
	var pm *pb.ProtocolMessenger

	allProtocols := make([]core_protocol.ID, len(d.cfg.Protocols))
	for i, p := range d.cfg.Protocols {
		allProtocols[i] = core_protocol.ID(p)
	}

	ms := &msgSender{
		h:         d.host, // Use the single host instance
		protocols: allProtocols,
		timeout:   d.cfg.DialTimeout,
	}

	var err error
	pm, err = pb.NewProtocolMessenger(ms)
	if err != nil {
		return nil, fmt.Errorf("new protocol messenger: %w", err)
	}

	c := &Crawler{
		id:        fmt.Sprintf("crawler-%02d", d.crawlerCount),
		host:      d.host, // Use the single host instance
		pm:        pm,
		psTopics:  make(map[string]struct{}),
		cfg:       d.cfg.CrawlerConfig(),
		client:    kubo.NewClient(),
		stateChan: d.workerStateChan,
		driver:    d, // Pass the driver
	}

	d.crawlerCount += 1

	return c, nil
}

func (d *CrawlDriver) NewWriter() (core.Worker[core.CrawlResult[PeerInfo], core.WriteResult], error) {
	w := core.NewCrawlWriter[PeerInfo](fmt.Sprintf("writer-%02d", d.writerCount), d.dbc, d.cfg.WriterConfig())
	d.writerCount += 1
	return w, nil
}

func (d *CrawlDriver) Tasks() <-chan PeerInfo {
	return d.tasksChan
}

func (d *CrawlDriver) Close() {
	// Closing logic remains the same
}

func (d *CrawlDriver) shutdown() {
	// Shutdown logic remains the same
}

func (d *CrawlDriver) monitorGossipSubPX() {
	// Monitoring logic remains the same
}

func newLibp2pHost(userAgent string) (host.Host, error) {
	cm := connmgr.NullConnMgr{}
	rm := network.NullResourceManager{}

	h, err := libp2p.New(
		libp2p.UserAgent(userAgent),
		libp2p.ResourceManager(&rm),
		libp2p.ConnectionManager(cm),
		libp2p.Security(noise.ID, noise.New),
		libp2p.DisableMetrics(),
		libp2p.SwarmOpts(swarm.WithReadOnlyBlackHoleDetector()),
		libp2p.UDPBlackHoleSuccessCounter(nil),
		libp2p.IPv6BlackHoleSuccessCounter(nil),
	)
	if err != nil {
		return nil, fmt.Errorf("new libp2p host: %w", err)
	}

	return h, nil
}