package libp2p

import (
	"context"
	"fmt"
	"os"
	//"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	//"github.com/libp2p/go-libp2p/core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
	//"go.uber.org/atomic"
	//"golang.org/x/sync/errgroup"
	crawlerconfig "github.com/dennis-tra/nebula-crawler/config"

	"github.com/dennis-tra/nebula-crawler/core"
	"github.com/dennis-tra/nebula-crawler/db"
	//pgmodels "github.com/dennis-tra/nebula-crawler/db/models/pg"
	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/logging"
	algonet "github.com/algorand/go-algorand/network"
	algonetp2p "github.com/algorand/go-algorand/network/p2p"
	dhtcrawler "github.com/libp2p/go-libp2p-kad-dht/crawler"
	identify "github.com/libp2p/go-libp2p/p2p/protocol/identify"
)

type P2PResult struct {
	RoutingTable *core.RoutingTable[PeerInfo]

	// All multi addresses that the remote peer claims to listen on
	// this can be different from the ones that we received from another peer
	// e.g., they could miss quic-v1 addresses if the reporting peer doesn't
	// know about that protocol.
	ListenMaddrs []ma.Multiaddr

	// The agent version of the crawled peer
	Agent string

	// The protocols the peer supports
	Protocols []string

	// Any error that has occurred when connecting to the peer
	ConnectError error

	// The above error transferred to a known error
	ConnectErrorStr string

	// Any error that has occurred during fetching neighbor information
	CrawlError error

	// The above error transferred to a known error
	CrawlErrorStr string

	// When was the connection attempt made
	ConnectStartTime time.Time

	// When have we established a successful connection
	ConnectEndTime time.Time

	// The multiaddress of the successful connection
	ConnectMaddr ma.Multiaddr

	// the transport of a successful connection
	Transport string
}

// crawlP2P establishes a connection and crawls neighbor info from a peer.
// It returns a channel that streams the crawling results asynchronously.
// The method retrieves routing table, listen addresses, protocols, and agent.
// Connection attempts and errors are tracked for debugging or analysis.
// It supports context cancellation for graceful operation termination.
func (c *Crawler) crawlP2P(ctx context.Context, pi PeerInfo) <-chan P2PResult {
	resultCh := make(chan P2PResult)

	go func() {
		defer close(resultCh)

		result := P2PResult{
			RoutingTable: &core.RoutingTable[PeerInfo]{PeerID: pi.ID()},
		}

		// Establish Algorand P2P network connection
		result.ConnectStartTime = time.Now()
		net, cleanup, err := c.connect(ctx, pi.AddrInfo)
		result.ConnectEndTime = time.Now()
		if cleanup != nil {
			defer cleanup()
		}

		if err == nil {
			// Crawl via DHT RPC using the live StreamHost
			result.RoutingTable, result.CrawlError = c.queryConnectedPeer(ctx, pi.AddrInfo, net)
			if h := net.HttpServer.StreamHost; h != nil {
				// brief wait window for identify to complete
				deadline := time.Now().Add(2 * time.Second)
				for {
					if av, err := h.Peerstore().Get(pi.ID(), "AgentVersion"); err == nil {
						if s, ok := av.(string); ok {
							result.Agent = s
						}
					}
					if prots, err := h.Peerstore().GetProtocols(pi.ID()); err == nil && len(prots) > 0 {
						result.Protocols = make([]string, 0, len(prots))
						for _, p := range prots {
							result.Protocols = append(result.Protocols, string(p))
						}
					}
					if result.Agent != "" || time.Now().After(deadline) {
						break
					}
					time.Sleep(100 * time.Millisecond)
				}
			}
			if result.CrawlError != nil {
				result.CrawlErrorStr = db.NetError(result.CrawlError)
			}
		} else {
			result.ConnectError = err
			result.ConnectErrorStr = db.NetError(err)
		}

		// send the result back and close channel
		select {
		case resultCh <- result:
		case <-ctx.Done():
		}
	}()

	return resultCh
}

// connect establishes a connection to the given peer.
func (c *Crawler) connect(ctx context.Context, pi peer.AddrInfo) (*algonet.P2PNetwork, func(), error) {
	if len(pi.Addrs) == 0 {
		return nil, nil, fmt.Errorf("no addresses provided for peer %s", pi.ID)
	}

	// Build full p2p multiaddrs for the target peer.
	p2pAddr, err := ma.NewMultiaddr("/p2p/" + pi.ID.String())
	if err != nil {
		return nil, nil, fmt.Errorf("invalid p2p multiaddr for peer %s: %w", pi.ID, err)
	}
	peerAddresses := make([]string, 0, len(pi.Addrs))
	for _, addr := range pi.Addrs {
		full := addr.Encapsulate(p2pAddr)
		peerAddresses = append(peerAddresses, full.String())
	}

	log := logging.NewLogger()
	log.SetLevel(logging.Info)

	cfg := config.GetDefaultLocal()
	cfg.DNSBootstrapID = ""
	cfg.EnableP2P = true
	cfg.NetAddress = "0.0.0.0:0"
	cfg.GossipFanout = 1

	dataDir, err := os.MkdirTemp("", "algorand-crawler")
	if err != nil {
		return nil, nil, fmt.Errorf("create temp dir: %w", err)
	}

	var networkName = c.cfg.Network

	var genesisInfo algonet.GenesisInfo
	switch networkName {
	case crawlerconfig.NetworkAlgoMainnet:
		genesisInfo = algonet.GenesisInfo{
			GenesisID: "mainnet-v1.0",
			NetworkID: "mainnet",
		}
	case crawlerconfig.NetworkAlgoTestnet:
		genesisInfo = algonet.GenesisInfo{
			GenesisID: "testnet-v1.0",
			NetworkID: "testnet",
		}
	case crawlerconfig.NetworkAlgoSuppranet:
		genesisInfo = algonet.GenesisInfo{
			GenesisID: "suppranet-v1.0",
			NetworkID: "suppranet",
		}
	default:
		_ = os.RemoveAll(dataDir)
		return nil, nil, fmt.Errorf("unsupported or non-algorand network specified for p2p crawl: %s", networkName)
	}
	nodeInfo := &nopeNodeInfo{}
	var identityOpts *algonet.IdentityOpts

	net, err := algonet.NewP2PNetwork(
		log,
		cfg,
		dataDir,
		peerAddresses,
		genesisInfo,
		nodeInfo,
		identityOpts,
		nil,
	)
	if err != nil {
		_ = os.RemoveAll(dataDir)
		return nil, nil, fmt.Errorf("create P2P network: %w", err)
	}

	if err := net.Start(); err != nil {
		net.Stop()
		_ = os.RemoveAll(dataDir)
		return nil, nil, fmt.Errorf("start P2P network: %w", err)
	}
	if h := net.HttpServer.StreamHost; h != nil {
		_, _ = identify.NewIDService(h)
	}

	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		if len(net.GetPeers(algonet.PeersConnectedOut)) > 0 {
			break
		}
		select {
		case <-ctx.Done():
			net.Stop()
			_ = os.RemoveAll(dataDir)
			return nil, nil, ctx.Err()
		case <-time.After(200 * time.Millisecond):
		}
	}

	cleanup := func() {
		net.Stop()
		_ = os.RemoveAll(dataDir)
	}

	log.Infof("Network started. Peer ID: %s", net.PeerID())
	return net, cleanup, nil
}

func (c *Crawler) queryConnectedPeer(ctx context.Context, pi peer.AddrInfo, net *algonet.P2PNetwork) (*core.RoutingTable[PeerInfo], error) {
	host := net.HttpServer.StreamHost

	dhtCrawler, err := dhtcrawler.NewDefaultCrawler(host, dhtcrawler.WithProtocols(c.cfg.Protocols))
	if err != nil {
		rt := &core.RoutingTable[PeerInfo]{
			PeerID: pi.ID,
			Error:  fmt.Errorf("failed to create dht crawler: %w", err),
		}
		return rt, rt.Error
	}

	gossipPeers, err := dhtCrawler.DhtRPC.GetClosestPeers(ctx, pi.ID, pi.ID)
	if err != nil {
		rt := &core.RoutingTable[PeerInfo]{
			PeerID:    pi.ID,
			Neighbors: []PeerInfo{},
			Error:     err,
		}
		return rt, err
	}

	routingTable := &core.RoutingTable[PeerInfo]{
		PeerID:    pi.ID,
		Neighbors: make([]PeerInfo, 0, len(gossipPeers)), // Pre-allocate slice capacity
	}

	for _, discoveredPeer := range gossipPeers {
		if discoveredPeer == nil {
			continue
		}
		routingTable.Neighbors = append(routingTable.Neighbors, PeerInfo{AddrInfo: *discoveredPeer})
	}

	return routingTable, nil
}

type nopeNodeInfo struct{}

func (n *nopeNodeInfo) IsParticipating() bool {
	return false
}

func (n *nopeNodeInfo) Capabilities() []algonetp2p.Capability {
	return nil
}

type SimpleMessageHandler struct{}

func (h *SimpleMessageHandler) Handle(msg algonet.IncomingMessage) algonet.OutgoingMessage {
	fmt.Printf("Received message with tag '%s' from peer: %s\n", msg.Tag, string(msg.Data))
	return algonet.OutgoingMessage{Action: algonet.Ignore}
}
