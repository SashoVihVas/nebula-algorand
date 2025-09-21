package libp2p

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p/core/protocol"
	"os"
	//"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	//"github.com/libp2p/go-libp2p/core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
	//"go.uber.org/atomic"
	//"golang.org/x/sync/errgroup"

	"github.com/dennis-tra/nebula-crawler/core"
	"github.com/dennis-tra/nebula-crawler/db"
	//pgmodels "github.com/dennis-tra/nebula-crawler/db/models/pg"
	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/logging"
	algonet "github.com/algorand/go-algorand/network"
	algonetp2p "github.com/algorand/go-algorand/network/p2p"
	aprotocol "github.com/algorand/go-algorand/protocol"
	dhtcrawler "github.com/libp2p/go-libp2p-kad-dht/crawler"
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
		// No need for the defer function to close the peer connection,
		// as the 'net' object from 'connect' handles its own lifecycle.

		result := P2PResult{
			RoutingTable: &core.RoutingTable[PeerInfo]{PeerID: pi.ID()},
		}

		// The original code used c.host to register for identify results.
		// Since the new logic uses a separate network stack from the `connect` function,
		// the original identify logic may no longer be applicable in the same way.
		// For this adjustment, we focus on the DHT query part as requested.
		// If identify information is still needed, it would require a deeper integration
		// with the algonet stack.

		var net *algonet.P2PNetwork // Adjusted type to match 'connect' return
		result.ConnectStartTime = time.Now()
		net, result.ConnectError = c.connect(ctx, pi.AddrInfo) // use filtered addr list
		result.ConnectEndTime = time.Now()

		// If we could successfully connect to the peer we actually crawl it.
		if result.ConnectError == nil {
			// Note: The original logic for getting ConnectMaddr, Transport, Agent, Protocols
			// was tied to the libp2p `network.Conn` and `identify` service.
			// The new `connect` function returns an `algonet.P2PNetwork` object.
			// We proceed with the core requirement: draining buckets with the new logic.

			// Fetch all neighbors using the new logic
			result.RoutingTable, result.CrawlError = c.queryConnectedPeer(ctx, pi.AddrInfo, net)
			if result.CrawlError != nil {
				result.CrawlErrorStr = db.NetError(result.CrawlError)
			}

			// The original GossipSub and Identify logic is omitted as it was
			// dependent on the original libp2p host and connection object.
			// The primary goal is to replace the neighbor discovery.
		} else {
			// if there was a connection error, parse it to a known one
			result.ConnectErrorStr = db.NetError(result.ConnectError)
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
func (c *Crawler) connect(ctx context.Context, pi peer.AddrInfo) (*algonet.P2PNetwork, error) {

	if len(pi.Addrs) == 0 {
		return nil, fmt.Errorf("no addresses provided for peer %s", pi.ID)
	}

	// Construct the full multiaddresses for the target peer.
	// e.g., /ip4/1.2.3.4/tcp/1234 + /p2p/QmPeerID...
	// becomes /ip4/1.2.3.4/tcp/1234/p2p/QmPeerID...
	p2pAddr, err := ma.NewMultiaddr("/p2p/" + pi.ID.String())
	if err != nil {
		return nil, fmt.Errorf("invalid p2p multiaddr for peer %s: %w", pi.ID, err)
	}

	peerAddresses := make([]string, 0, len(pi.Addrs))
	for _, addr := range pi.Addrs {
		fullAddr := addr.Encapsulate(p2pAddr)
		peerAddresses = append(peerAddresses, fullAddr.String())
	}

	log := logging.NewLogger()
	log.SetLevel(logging.Info)

	cfg := config.GetDefaultLocal()
	cfg.DNSBootstrapID = "" // Should be commented out if you want to use DNS bootstrap
	//cfg.EnableDHTProviders = true // Enable DHT providers to find peers
	cfg.EnableP2P = true // Enable P2P networking
	cfg.NetAddress = "0.0.0.0:4161"
	cfg.GossipFanout = 1
	dataDir, err := os.MkdirTemp("", "algorand-crawler")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(dataDir)

	genesisInfo := algonet.GenesisInfo{
		GenesisID: "testnet-v1.0",
		NetworkID: "testnet",
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
		log.Fatalf("Failed to create P2P network: %v", err)
	}

	handler := &SimpleMessageHandler{}

	net.RegisterHandlers([]algonet.TaggedMessageHandler{
		{Tag: aprotocol.TxnTag, MessageHandler: handler},
	})

	if err := net.Start(); err != nil {
		log.Fatalf("Failed to start network: %v", err)
	}
	defer net.Stop()

	log.Infof("Network started. Peer ID: %s", net.PeerID())
	log.Info("Waiting to connect to peer...")
	for {
		if len(net.GetPeers(algonet.PeersConnectedOut)) > 0 {
			log.Info("Connected to at least one peer.")
			break
		}
		time.Sleep(2 * time.Second)
	}
	return net, err
}

func (c *Crawler) queryConnectedPeer(ctx context.Context, pi peer.AddrInfo, net *algonet.P2PNetwork) (*core.RoutingTable[PeerInfo], error) {
	host := net.HttpServer.StreamHost

	// Define the custom protocol ID for Algorand's DHT
	algorandProtocolID := protocol.ID("/algorand/kad/testnet/kad/1.0.0")

	// Now, create the crawler using the DHT from the connected Algorand node
	dhtCrawler, err := dhtcrawler.NewDefaultCrawler(host, dhtcrawler.WithProtocols([]protocol.ID{algorandProtocolID}))
	if err != nil {
		// If crawler creation fails, return the error wrapped in the target struct format.
		rt := &core.RoutingTable[PeerInfo]{
			PeerID: pi.ID,
			Error:  fmt.Errorf("failed to create dht crawler: %w", err),
		}
		return rt, rt.Error
	}

	// Query the connected peer (pi.ID) for peers closest to its own ID.
	// This is a simple, effective way to ask "who do you know?"
	gossipPeers, err := dhtCrawler.DhtRPC.GetClosestPeers(ctx, pi.ID, pi.ID)
	if err != nil {
		// If the query fails, populate the error fields and return.
		rt := &core.RoutingTable[PeerInfo]{
			PeerID:    pi.ID,
			Neighbors: []PeerInfo{},
			Error:     err,
		}
		return rt, err
	}

	// On success, format the gossipPeers into the target routing table structure.
	routingTable := &core.RoutingTable[PeerInfo]{
		PeerID:    pi.ID,
		Neighbors: make([]PeerInfo, 0, len(gossipPeers)), // Pre-allocate slice capacity
	}

	// Convert each `*peer.AddrInfo` from gossipPeers into the `PeerInfo` type.
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
	// We can choose to broadcast the message further, ignore it, or respond.
	// For this example, we'll just ignore it.
	return algonet.OutgoingMessage{Action: algonet.Ignore}
}
