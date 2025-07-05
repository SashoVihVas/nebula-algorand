package libp2p

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/textproto"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dennis-tra/nebula-crawler/config"
	"github.com/dennis-tra/nebula-crawler/core"
	"github.com/dennis-tra/nebula-crawler/db"
	"github.com/dennis-tra/nebula-crawler/kubo"
	"github.com/dennis-tra/nebula-crawler/utils"
	"github.com/tinylib/msgp/msgp"

	"github.com/libp2p/go-libp2p"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/record"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	"github.com/libp2p/go-msgio"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// Handshake implementation based on go-algorand source code.
const (
	ProtocolVersionHeader       = "X-Algorand-Version"
	ProtocolAcceptVersionHeader = "X-Algorand-Accept-Version"
	TelemetryIDHeader           = "X-Algorand-TelId"
	GenesisHeader               = "X-Algorand-Genesis"
	NodeRandomHeader            = "X-Algorand-NodeRandom"
	AddressHeader               = "X-Algorand-Location"
	InstanceNameHeader          = "X-Algorand-InstanceName"
	PeerFeaturesHeader          = "X-Algorand-Peer-Features"
	PeerFeatureProposalCompression = "ppzstd"
	PeerFeatureVoteVpackCompression = "avvpack"

	maxHeaderKeys   = 64
	maxHeaderValues = 16
)

type peerMetadataProvider interface {
	TelemetryGUID() string
	InstanceName() string
	GenesisID() string
	PublicAddress() string
	RandomID() string
	SupportedProtoVersions() []string
}

type peerMetaInfo struct {
	telemetryID  string
	instanceName string
	version      string
	features     string
}

type peerMetaValues []string

type peerMetaHeaders map[string]peerMetaValues

func peerMetaHeadersToHTTPHeaders(headers peerMetaHeaders) http.Header {
	httpHeaders := make(http.Header, len(headers))
	for k, v := range headers {
		httpHeaders[k] = v
	}
	return httpHeaders
}

func peerMetaHeadersFromHTTPHeaders(headers http.Header) peerMetaHeaders {
	pmh := make(peerMetaHeaders, len(headers))
	for k, v := range headers {
		pmh[k] = v
	}
	return pmh
}

func setHeaders(header http.Header, netProtoVer string, meta peerMetadataProvider) {
	header.Set(TelemetryIDHeader, meta.TelemetryGUID())
	header.Set(InstanceNameHeader, meta.InstanceName())
	if pa := meta.PublicAddress(); pa != "" {
		header.Set(AddressHeader, pa)
	}
	if rid := meta.RandomID(); rid != "" {
		header.Set(NodeRandomHeader, rid)
	}
	header.Set(GenesisHeader, meta.GenesisID())

	features := []string{PeerFeatureProposalCompression}

	header.Set(PeerFeaturesHeader, strings.Join(features, ","))

	if netProtoVer != "" {
		header.Set(ProtocolVersionHeader, netProtoVer)
	}
	for _, v := range meta.SupportedProtoVersions() {
		header.Add(ProtocolAcceptVersionHeader, v)
	}
}

func checkProtocolVersionMatch(otherHeaders http.Header, ourSupportedProtocolVersions []string) (string, string) {
	otherAcceptedVersions := otherHeaders[textproto.CanonicalMIMEHeaderKey(ProtocolAcceptVersionHeader)]
	for _, otherAcceptedVersion := range otherAcceptedVersions {
		// --- Start of suggested change ---
		for _, ourProto := range ourSupportedProtocolVersions {
			// Check if our full protocol ID (e.g., "/algorand-ws/2.2.0")
			// ends with the node's version (e.g., "/2.2" or "/2.2.0")
			if strings.HasSuffix(ourProto, "/"+otherAcceptedVersion) || strings.HasSuffix(ourProto, "/"+otherAcceptedVersion+".0") {
				return otherAcceptedVersion, ""
			}
		}
		// --- End of suggested change ---
	}

	otherVersion := otherHeaders.Get(ProtocolVersionHeader)
	// --- Start of suggested change ---
	for _, ourProto := range ourSupportedProtocolVersions {
		if strings.HasSuffix(ourProto, "/"+otherVersion) || strings.HasSuffix(ourProto, "/"+otherVersion+".0") {
			return otherVersion, otherVersion
		}
	}
	// --- End of suggested change ---

	return "", otherVersion
}

type peerFeatureFlag int

const (
	pfCompressedProposal peerFeatureFlag = 1 << iota
	pfCompressedVoteVpack
)

const versionPeerFeatures = "2.2"

var versionPeerFeaturesNum [2]int64

func init() {
	var err error
	versionPeerFeaturesNum[0], versionPeerFeaturesNum[1], err = versionToMajorMinor(versionPeerFeatures)
	if err != nil {
		panic(fmt.Sprintf("failed to parse version %v: %s", versionPeerFeatures, err.Error()))
	}
}

func versionToMajorMinor(version string) (int64, int64, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("version %s does not have two components", version)
	}
	major, err := strconv.ParseInt(parts[0], 10, 8)
	if err != nil {
		return 0, 0, err
	}
	minor, err := strconv.ParseInt(parts[1], 10, 8)
	if err != nil {
		return 0, 0, err
	}
	return major, minor, nil
}

func decodePeerFeatures(version string, announcedFeatures string) peerFeatureFlag {
	major, minor, err := versionToMajorMinor(version)
	if err != nil {
		return 0
	}
	if major < versionPeerFeaturesNum[0] || (major == versionPeerFeaturesNum[0] && minor < versionPeerFeaturesNum[1]) {
		return 0
	}
	var features peerFeatureFlag
	parts := strings.Split(announcedFeatures, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == PeerFeatureProposalCompression {
			features |= pfCompressedProposal
		}
		if part == PeerFeatureVoteVpackCompression {
			features |= pfCompressedVoteVpack
		}
	}
	return features
}

func readPeerMetaHeaders(stream io.ReadWriter, p2pPeer peer.ID, netProtoSupportedVersions []string) (peerMetaInfo, error) {
	var msgLenBytes [2]byte
	rn, err := stream.Read(msgLenBytes[:])
	if rn != 2 || err != nil {
		err0 := fmt.Errorf("error reading response message length from peer %s: %w", p2pPeer, err)
		return peerMetaInfo{}, err0
	}
	log.WithField("Received message's length:", msgLenBytes).Info("TRALALALA")

	msgLen := binary.BigEndian.Uint16(msgLenBytes[:])
	msgBytes := make([]byte, msgLen)
	rn, err = stream.Read(msgBytes[:])
	if rn != int(msgLen) || err != nil {
		err0 := fmt.Errorf("error reading response message from peer %s: %w, expected: %d, read: %d", p2pPeer, err, msgLen, rn)
		return peerMetaInfo{}, err0
	}
	log.Infof("Received message from %s: %x", p2pPeer, msgBytes)

	var responseHeaders peerMetaHeaders
	_, err = responseHeaders.UnmarshalMsg(msgBytes[:])
	if err != nil {
		err0 := fmt.Errorf("error unmarshaling response message from peer %s: %w", p2pPeer, err)
		return peerMetaInfo{}, err0
	}
	headers := peerMetaHeadersToHTTPHeaders(responseHeaders)
	matchingVersion, _ := checkProtocolVersionMatch(headers, netProtoSupportedVersions)
	if matchingVersion == "" {
		err0 := fmt.Errorf("peer %s does not support any of the supported protocol versions: %v", p2pPeer, netProtoSupportedVersions)
		return peerMetaInfo{}, err0
	}
	return peerMetaInfo{
		telemetryID:  headers.Get(TelemetryIDHeader),
		instanceName: headers.Get(InstanceNameHeader),
		version:      matchingVersion,
		features:     headers.Get(PeerFeaturesHeader),
	}, nil
}

func writePeerMetaHeaders(stream io.ReadWriter, p2pPeer peer.ID, networkProtoVersion string, pmp peerMetadataProvider) error {
	header := make(http.Header)
	setHeaders(header, networkProtoVersion, pmp)
	meta := peerMetaHeadersFromHTTPHeaders(header)
	data, err := meta.MarshalMsg(nil)
	if err != nil {
		return fmt.Errorf("error marshaling peer meta headers: %w", err)
	}
	length := len(data)
	if length > math.MaxUint16 {
		msg := fmt.Sprintf("error writing initial message, too large: %v, peer %s", header, p2pPeer)
		panic(msg)
	}
	metaMsg := make([]byte, 2+length)
	binary.BigEndian.PutUint16(metaMsg, uint16(length))
	copy(metaMsg[2:], data)
	log.WithField("Received message:", metaMsg).Info("LALALA")
	_, err = stream.Write(metaMsg)
	if err != nil {
		err0 := fmt.Errorf("error sending initial message: %w", err)
		return err0
	}
	return nil
}

type PeerInfo struct {
	peer.AddrInfo
}

func (d *CrawlDriver) algorandStreamHandler(stream network.Stream) {
	log.WithField("remotePeer", stream.Conn().RemotePeer()).Info("New Algorand stream")

	var err error
	var pmi peerMetaInfo
	if stream.Stat().Direction == network.DirOutbound {
		log.WithField("remotePeer", stream.Conn().RemotePeer()).Info("WE ARE INBOUND")
		err = writePeerMetaHeaders(stream, stream.Conn().RemotePeer(), d.protocolVersion, d)
		if err != nil {
			log.WithError(err).Warn("error reading peer meta headers")
			_ = stream.Reset()
			return
		}
		pmi, err = readPeerMetaHeaders(stream, stream.Conn().RemotePeer(), d.cfg.Protocols)
		if err != nil {
			log.WithError(err).Warn("error writing peer meta headers")
			_ = stream.Reset()
			return
		}
	} else {
		log.WithField("remotePeer", stream.Conn().RemotePeer()).Info("WE ARE OUTBOUND")
		err = writePeerMetaHeaders(stream, stream.Conn().RemotePeer(), d.protocolVersion, d)
		if err != nil {
			log.WithError(err).Warn("error writing peer meta headers")
			_ = stream.Reset()
			return
		}
		pmi, err = readPeerMetaHeaders(stream, stream.Conn().RemotePeer(), d.cfg.Protocols)
		if err != nil {
			log.WithError(err).Warn("error reading peer meta headers")
			_ = stream.Reset()
			return
		}
	}
log.WithField("pmi_version", pmi.version).WithField("pmi_features", pmi.features).Info("Algorand handshake successful")
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
	cfg   *CrawlDriverConfig
	hosts map[peer.ID]*Host
	dbc   db.Client

	pxPeersChan     chan []PeerInfo
	tasksChan       chan PeerInfo
	workerStateChan chan string
	crawlerCount    int
	writerCount     int
	protocolVersion string
}

func (d *CrawlDriver) TelemetryGUID() string {
	return "nebula-crawler"
}

func (d *CrawlDriver) InstanceName() string {
	return "nebula-instance"
}

func (d *CrawlDriver) GenesisID() string {
	return "testnet-v1.0"
}

func (d *CrawlDriver) PublicAddress() string {
	return ""
}

func (d *CrawlDriver) RandomID() string {
	return ""
}

func (d *CrawlDriver) SupportedProtoVersions() []string {
	return d.cfg.Protocols
}


var _ core.Driver[PeerInfo, core.CrawlResult[PeerInfo]] = (*CrawlDriver)(nil)

func NewCrawlDriver(dbc db.Client, cfg *CrawlDriverConfig) (*CrawlDriver, error) {
	userAgent := "nebula/" + cfg.Version
	if cfg.Network == config.NetworkAvailTuringLC || cfg.Network == config.NetworkAvailMainnetLC {
		userAgent = "avail-light-client/light-client/1.12.13/rust-client"
	}

	hosts := make(map[peer.ID]*Host, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		h, err := newLibp2pHost(userAgent)
		if err != nil {
			return nil, fmt.Errorf("new libp2p host: %w", err)
		}
		hosts[h.ID()] = h
	}

	tasksChan := make(chan PeerInfo, len(cfg.BootstrapPeers))
	for _, addrInfo := range cfg.BootstrapPeers {
		addrInfo := addrInfo
		tasksChan <- PeerInfo{AddrInfo: addrInfo}
	}

	d := &CrawlDriver{
		cfg:             cfg,
		hosts:           hosts,
		dbc:             dbc,
		tasksChan:       tasksChan,
		pxPeersChan:     make(chan []PeerInfo),
		workerStateChan: make(chan string, cfg.WorkerCount),
		crawlerCount:    0,
		writerCount:     0,
		protocolVersion: "2.2",
	}

	if cfg.Network == config.NetworkAlgoTestnet {
		for _, h := range hosts {
			h.SetStreamHandler("/algorand-ws/2.2.0", d.algorandStreamHandler)
		}
	}

	if cfg.GossipSubPX {
		go d.monitorGossipSubPX()

		for _, h := range d.hosts {
			for _, protID := range d.cfg.Protocols {
				h.SetStreamHandler(protocol.ID(protID), d.handleGossipSubStream)
			}
		}
	} else {
		close(tasksChan)
	}

	return d, nil
}

func (d *CrawlDriver) NewWorker() (core.Worker[PeerInfo, core.CrawlResult[PeerInfo]], error) {
	hostsList := make([]string, 0, len(d.hosts))
	for _, h := range d.hosts {
		hostsList = append(hostsList, string(h.ID()))
	}
	sort.Strings(hostsList)
	hostID := peer.ID(hostsList[d.crawlerCount%len(d.hosts)])

	allProtocols := []protocol.ID{
		protocol.ID("/algorand-ws/2.2.0"),
		protocol.ID("/algorand/kad/testnet-v1.0"),
	}

	ms := &msgSender{
		h:         d.hosts[hostID].Host,
		protocols: allProtocols, // Use the combined list here
		timeout:   d.cfg.DialTimeout,
	}

	pm, err := pb.NewProtocolMessenger(ms)
	if err != nil {
		return nil, fmt.Errorf("new protocol messenger: %w", err)
	}

	c := &Crawler{
		id:        fmt.Sprintf("crawler-%02d", d.crawlerCount),
		host:      d.hosts[hostID],
		pm:        pm,
		psTopics:  make(map[string]struct{}),
		cfg:       d.cfg.CrawlerConfig(),
		client:    kubo.NewClient(),
		stateChan: d.workerStateChan,
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
	shutdown := make(chan struct{})
	go func() {
		d.shutdown()
		close(shutdown)
	}()

	select {
	case <-shutdown:
	case <-time.After(10 * time.Second):
		log.Warnln("shutdown timed out")
	}
}

func (d *CrawlDriver) shutdown() {
	var wgHostClose sync.WaitGroup
	for _, h := range d.hosts {
		wgHostClose.Add(1)
		go func(h *Host) {
			defer wgHostClose.Done()
			if err := h.Close(); err != nil {
				log.WithError(err).Warnln("failed to close host")
			}
		}(h)
	}

	var wgTasksClose sync.WaitGroup
	wgTasksClose.Add(1)
	go func() {
		for range d.tasksChan {
		}
		wgTasksClose.Done()
	}()

	wgHostClose.Wait()
	close(d.pxPeersChan)
	wgTasksClose.Wait()
}

func (d *CrawlDriver) monitorGossipSubPX() {
	defer close(d.tasksChan)

	timeout := time.Second
	timer := time.NewTimer(math.MaxInt64)

	busyWorkers := 0
LOOP:
	for {
		select {
		case <-timer.C:
			log.Infof("All workers idle for %s. Stop monitoring gossipsub streams.", timeout)
			break LOOP
		case state := <-d.workerStateChan:
			switch state {
			case "busy":
				busyWorkers += 1
			case "idle":
				busyWorkers -= 1
			}
			if busyWorkers == 0 {
				timer.Reset(timeout)
			} else {
				timer.Reset(math.MaxInt64)
			}
		case pxPeers, more := <-d.pxPeersChan:
			if !more {
				break LOOP
			}

			log.Infof("Discovered %d peers via gossipsub\n", len(pxPeers))
			for _, pxPeer := range pxPeers {
				d.tasksChan <- pxPeer
			}
		}
	}
}

func newLibp2pHost(userAgent string) (*Host, error) {
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

	return WrapHost(h)
}

func (d *CrawlDriver) handleGossipSubStream(incomingStream network.Stream) {
	defer func() {
		if err := incomingStream.Reset(); err != nil {
			log.WithError(err).Warnln("Failed to reset incoming stream")
		}
	}()

	remoteID := incomingStream.Conn().RemotePeer()
	localID := incomingStream.Conn().LocalPeer()

	helloRPC, err := readRPC(incomingStream)
	if err != nil {
		return
	}

	outgoingStream, err := d.hosts[localID].NewStream(context.Background(), remoteID, pubsub.GossipSubDefaultProtocols...)
	if err != nil {
		return
	}
	defer func() {
		if err := outgoingStream.Reset(); err != nil {
			log.WithError(err).Warnln("Failed to reset outgoing stream")
		}
	}()

	if err = writeRPC(outgoingStream, helloRPC); err != nil {
		return
	}

	graftRPC := &pubsub_pb.RPC{
		Control: &pubsub_pb.ControlMessage{
			Graft: make([]*pubsub_pb.ControlGraft, len(helloRPC.GetSubscriptions())),
		},
	}
	for i, sub := range helloRPC.GetSubscriptions() {
		graftRPC.Control.Graft[i] = &pubsub_pb.ControlGraft{
			TopicID: sub.Topicid,
		}
	}

	if err = writeRPC(outgoingStream, graftRPC); err != nil {
		return
	}

	for {
		rpc, err := readRPC(incomingStream)
		if err != nil {
			return
		}

		pxPeers := make([]PeerInfo, 0)
		for _, prune := range rpc.GetControl().GetPrune() {
			for _, p := range prune.GetPeers() {
				addrInfo, err := parseSignedPeerRecord(p.SignedPeerRecord)
				if err != nil {
					log.WithError(err).Debugln("failed to parse signed peer record")
					continue
				}
				pxPeers = append(pxPeers, PeerInfo{AddrInfo: *addrInfo})
			}
		}

		if len(pxPeers) > 0 {
			d.pxPeersChan <- pxPeers
			return
		}
	}
}

func readRPC(s network.Stream) (*pubsub_pb.RPC, error) {
	data, err := msgio.NewVarintReader(s).ReadMsg()
	if err != nil {
		return nil, fmt.Errorf("failed to read message: %w", err)
	}

	rpc := &pubsub_pb.RPC{}
	if err = rpc.Unmarshal(data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal rpc message: %w", err)
	}

	return rpc, nil
}

func writeRPC(s network.Stream, rpc *pubsub_pb.RPC) error {
	data, err := rpc.Marshal()
	if err != nil {
		return fmt.Errorf("failed to marshal rpc message: %w", err)
	}

	if err = msgio.NewVarintWriter(s).WriteMsg(data); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	return nil
}

func parseSignedPeerRecord(signedPeerRecord []byte) (*peer.AddrInfo, error) {
	envelope, err := record.UnmarshalEnvelope(signedPeerRecord)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal signed peer record: %w", err)
	}

	pid, err := peer.IDFromPublicKey(envelope.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive peer ID: %s", err)
	}
	r, err := envelope.Record()
	if err != nil {
		return nil, fmt.Errorf("failed to obtain record: %w", err)
	}

	rec, ok := r.(*peer.PeerRecord)
	if !ok {
		return nil, fmt.Errorf("not a peer record")
	}

	addrInfo := peer.AddrInfo{
		ID:    pid,
		Addrs: rec.Addrs,
	}

	return &addrInfo, nil
}

func openInboundGossipSubStreams(h host.Host, pid peer.ID) int {
	openStreams := 0
	for _, conn := range h.Network().ConnsToPeer(pid) {
		for _, stream := range conn.GetStreams() {
			if stream.Stat().Direction != network.DirInbound {
				continue
			}
			switch stream.Protocol() {
			case pubsub.GossipSubID_v10, pubsub.GossipSubID_v11, pubsub.GossipSubID_v12, pubsub.FloodSubID:
				openStreams++
			}
		}
	}
	return openStreams
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *peerMetaHeaders) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0004 uint32 // Keep as uint32
	zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if (*z) == nil {
		(*z) = make(peerMetaHeaders, int(zb0004)) // Cast to int for make()
	} else if len(*z) > 0 {
		for key := range *z {
			delete(*z, key)
		}
	}
	for zb0004 > 0 {
		var zb0001 string
		var zb0002 peerMetaValues
		zb0004--
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		var zb0006 uint32 // Keep as uint32
		zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		if cap(zb0002) >= int(zb0006) {
			zb0002 = (zb0002)[:int(zb0006)] // Cast to int for slicing
		} else {
			zb0002 = make(peerMetaValues, int(zb0006)) // Cast to int for make()
		}
		for zb0003 := range zb0002 {
			zb0002[zb0003], bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, zb0001, zb0003)
				return
			}
		}
		(*z)[zb0001] = zb0002
	}
	o = bts
	return
}

// MarshalMsg implements msgp.Marshaler
func (z peerMetaHeaders) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, uint32(len(z)))
	keys := make([]string, 0, len(z))
	for k := range z {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		o = msgp.AppendString(o, k)
		o = msgp.AppendArrayHeader(o, uint32(len(z[k])))
		for za0002 := range z[k] {
			o = msgp.AppendString(o, z[k][za0002])
		}
	}
	return o, nil
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z peerMetaHeaders) Msgsize() (s int) {
	s = msgp.MapHeaderSize
	if z != nil {
		for k, v := range z {
			_ = k
			_ = v
			s += msgp.StringPrefixSize + len(k) + msgp.ArrayHeaderSize
			for za0002 := range v {
				s += msgp.StringPrefixSize + len(v[za0002])
			}
		}
	}
	return
}