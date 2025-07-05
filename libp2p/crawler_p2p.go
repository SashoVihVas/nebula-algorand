package libp2p

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"

	"github.com/dennis-tra/nebula-crawler/config"
	"github.com/dennis-tra/nebula-crawler/core"
	"github.com/dennis-tra/nebula-crawler/db"
	pgmodels "github.com/dennis-tra/nebula-crawler/db/models/pg"
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
		defer func() {
			// Free connection resources
			if err := c.host.Network().ClosePeer(pi.ID()); err != nil {
				log.WithError(err).WithField("remoteID", pi.ID().ShortString()).Warnln("Could not close connection to peer")
			}
		}()

		result := P2PResult{
			RoutingTable: &core.RoutingTable[PeerInfo]{PeerID: pi.ID()},
		}

		// register the given peer (before connecting) to receive
		// the identify result on the returned channel
		identifyChan := c.host.RegisterIdentify(pi.ID())

		var conn network.Conn
		result.ConnectStartTime = time.Now()
		conn, result.ConnectError = c.connect(ctx, pi.AddrInfo) // use filtered addr list
		result.ConnectEndTime = time.Now()

		// If we could successfully connect to the peer we actually crawl it.
		if result.ConnectError == nil {
			// keep track of the multi address over which we successfully connected
			result.ConnectMaddr = conn.RemoteMultiaddr()

			// keep track of the transport of the open connection
			result.Transport = conn.ConnState().Transport

			// Wait for the stream to be ready
			var stream network.Stream
			for i := 0; i < 10; i++ { // Poll for up to 5 seconds
				if val, ok := c.driver.activeStreams.Load(pi.ID()); ok {
					stream = val.(network.Stream)
					break
				}
				time.Sleep(500 * time.Millisecond)
			}

			if stream == nil {
				result.CrawlError = fmt.Errorf("timed out waiting for stream to be ready")
			} else {
				// Fetch all neighbors
				result.RoutingTable, result.CrawlError = c.drainBuckets(ctx, pi.AddrInfo)
			}

			if result.CrawlError != nil {
				result.CrawlErrorStr = db.NetError(result.CrawlError)
			}

			// wait for the Identify exchange to complete (no-op if already done)
			// the internal timeout is set to 30 s. When crawling we only allow 5s.
			timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			select {
			case <-timeoutCtx.Done():
				// identification timed out.
			case identify, more := <-identifyChan:
				// identification may have succeeded.
				if !more {
					break
				}

				result.Agent = identify.AgentVersion
				result.ListenMaddrs = identify.ListenAddrs
				result.Protocols = make([]string, len(identify.Protocols))
				for i := range identify.Protocols {
					result.Protocols[i] = string(identify.Protocols[i])
				}
			}

			if c.cfg.GossipSubPX {
				// give the other peer a chance to open a stream to and prune us
				streams := openInboundGossipSubStreams(c.host, pi.ID())

				// The maximum time to wait for the gossipsub px to complete
				maxGossipSubWait := c.cfg.DialTimeout

				// the minimum time to wait for the gossipsub px to start
				minGossipSubWait := 2 * time.Second

				// the time since we're connected to the peer
				elapsed := time.Since(result.ConnectEndTime)

				// the remaining time until the maximum wait time is reached
				remainingWait := maxGossipSubWait - elapsed

				// the interval to check the open gossipsub streams
				interval := 250 * time.Millisecond

				// if 1) we are supposed to wait a little longer for the
				// gossipsub exchange (remainingWait > 0) and 2) we either have
				// at least one open gossipsubstream or haven't waited long enough
				// for such a stream to be there yet then we will enter the for
				// loop that waits the calculated remaining time before exiting
				// or until we don't have any open gossipsub streams anymore by
				// checking every 250ms if there are still any.

				if remainingWait > 0 && (streams != 0 || elapsed < minGossipSubWait) {

					// if we don't have an open stream yet and the check
					// interval is way below the minimum wait time we increase
					// the initial ticker delay
					initialTickerDelay := interval
					if streams == 0 && minGossipSubWait-elapsed > interval {
						initialTickerDelay = minGossipSubWait - elapsed
					}

					timer := time.NewTimer(remainingWait)
					ticker := time.NewTicker(initialTickerDelay)

					defer timer.Stop()
					defer ticker.Stop()

					for {
						select {
						case <-ctx.Done():
							// exit for loop because the context was cancelled
						case <-timer.C:
							// exit for loop because the maximum wait time was reached
						case <-ticker.C:
							ticker.Reset(interval)
							if openInboundGossipSubStreams(c.host, pi.ID()) != 0 {
								continue
							}
							// exit for loop because we don't have any open
							// streams despite waiting minGossipSubWait
						}
						break
					}
				}
			}
		} else {
			// if there was a connection error, parse it to a known one
			result.ConnectErrorStr = db.NetError(result.ConnectError)
		}

		// deregister peer from identify messages
		c.host.DeregisterIdentify(pi.ID())

		// send the result back and close channel
		select {
		case resultCh <- result:
		case <-ctx.Done():
		}
	}()

	return resultCh
}

// connect establishes a connection to the given peer.
func (c *Crawler) connect(ctx context.Context, pi peer.AddrInfo) (network.Conn, error) {
	if len(pi.Addrs) == 0 {
		return nil, db.ErrNoPublicIP
	}

	// init an exponential backoff
	bo := backoff.NewExponentialBackOff()
	bo.InitialInterval = time.Second
	bo.MaxInterval = 10 * time.Second
	bo.MaxElapsedTime = time.Minute
	bo.Clock = c.cfg.Clock

	// keep track of retries for debug logging
	retry := 0

	for {
		logEntry := log.WithFields(log.Fields{
			"timeout":  c.cfg.DialTimeout.String(),
			"remoteID": pi.ID.String(),
			"retry":    retry,
			"maddrs":   pi.Addrs,
		})
		logEntry.Debugln("Connecting to peer", pi.ID.ShortString())

		// save addresses into the peer store temporarily
		c.host.Peerstore().AddAddrs(pi.ID, pi.Addrs, peerstore.TempAddrTTL)

		timeoutCtx, cancel := context.WithTimeout(ctx, c.cfg.DialTimeout)
		conn, err := c.host.Network().DialPeer(timeoutCtx, pi.ID)
		cancel()

		// yay, it worked! Or has it? The caller checks the connectedness again.
		if err == nil {
			log.Infof("Successfully connected to peer %s", pi.ID.ShortString())
			return conn, nil
		}

		log.WithError(err).Errorf("Failed to connect to peer %s. Full error: %v", pi.ID.ShortString(), err)

		switch true {
		case strings.Contains(err.Error(), db.ErrorStr[pgmodels.NetErrorConnectionRefused]):
			// Might be transient because the remote doesn't want us to connect.
			// Try again, but reduce the maximum elapsed time because it's still
			// unlikely to succeed
			bo.MaxElapsedTime = 2 * c.cfg.DialTimeout
		case strings.Contains(err.Error(), db.ErrorStr[pgmodels.NetErrorConnectionGated]):
			// Hints at a configuration issue and should not happen, but if it
			// does it could be transient. Try again anyway, but at least log a warning.
			logEntry.WithError(err).Warnln("Connection gated!")
		case strings.Contains(err.Error(), db.ErrorStr[pgmodels.NetErrorCantAssignRequestedAddress]):
			// Transient error due to local UDP issues. Try again!
		case strings.Contains(err.Error(), "dial backoff"):
			// should not happen because we disabled backoff checks with our
			// go-libp2p fork. Try again anyway, but at least log a warning.
			logEntry.WithError(err).Warnln("Dial backoff!")
		case strings.Contains(err.Error(), "RESOURCE_LIMIT_EXCEEDED (201)"): // thrown by a circuit relay
			// We already have too many open connections over a relay. Try again!
		default:
			logEntry.WithError(err).Debugln("Failed connecting to peer", pi.ID.ShortString())
			return nil, err
		}

		sleepDur := bo.NextBackOff()
		if sleepDur == backoff.Stop {
			logEntry.WithError(err).Debugln("Exceeded retries connecting to peer", pi.ID.ShortString())
			return nil, err
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(sleepDur):
			retry += 1
			continue
		}

	}
}

func (c *Crawler) drainBuckets(ctx context.Context, pi peer.AddrInfo) (*core.RoutingTable[PeerInfo], error) {
	rt, err := kbucket.NewRoutingTable(20, kbucket.ConvertPeerID(pi.ID), time.Hour, nil, time.Hour, nil)
	if err != nil {
		return nil, err
	}

	allNeighborsLk := sync.RWMutex{}
	allNeighbors := map[peer.ID]peer.AddrInfo{}

	errorBits := atomic.NewUint32(0)

	errg := errgroup.Group{}
	for i := uint(0); i <= 15; i++ { // 15 is maximum
		count := i // Copy value
		errg.Go(func() error {
			var neighbors []*peer.AddrInfo
			var err error
			if c.driver.cfg.Network == config.NetworkAlgoTestnet {
				neighbors, err = c.drainBucketAlgorand(ctx, rt, pi.ID, count)
			} else {
				neighbors, err = c.drainBucketDefault(ctx, rt, pi.ID, count)
			}

			if err != nil {
				errorBits.Add(1 << count)
				return err
			}

			allNeighborsLk.Lock()
			for _, n := range neighbors {
				allNeighbors[n.ID] = *n
			}
			allNeighborsLk.Unlock()

			return nil
		})
	}
	err = errg.Wait()

	routingTable := &core.RoutingTable[PeerInfo]{
		PeerID:    pi.ID,
		Neighbors: []PeerInfo{},
		ErrorBits: uint16(errorBits.Load()),
		Error:     err,
	}

	for _, n := range allNeighbors {
		routingTable.Neighbors = append(routingTable.Neighbors, PeerInfo{AddrInfo: n})
	}

	return routingTable, err
}

func (c *Crawler) drainBucketDefault(ctx context.Context, rt *kbucket.RoutingTable, pid peer.ID, bucket uint) ([]*peer.AddrInfo, error) {
	rpi, err := rt.GenRandPeerID(bucket)
	if err != nil {
		log.WithError(err).WithField("enr", pid.ShortString()).WithField("bucket", bucket).Warnln("Failed generating random peer ID")
		return nil, fmt.Errorf("generating random peer ID with CPL %d: %w", bucket, err)
	}

	var neighbors []*peer.AddrInfo
	for retry := 0; retry < 2; retry++ {
		neighbors, err = c.pm.GetClosestPeers(ctx, pid, rpi)
		if err == nil {
			// getting closest peers was successful!
			return neighbors, nil
		}
		// ... (error handling as in original file)
	}
	return nil, fmt.Errorf("getting closest peer with CPL %d: %w", bucket, err)
}

func (c *Crawler) drainBucketAlgorand(ctx context.Context, rt *kbucket.RoutingTable, pid peer.ID, bucket uint) ([]*peer.AddrInfo, error) {
	rpi, err := rt.GenRandPeerID(bucket)
	if err != nil {
		log.WithError(err).WithField("enr", pid.ShortString()).WithField("bucket", bucket).Warnln("Failed generating random peer ID")
		return nil, fmt.Errorf("generating random peer ID with CPL %d: %w", bucket, err)
	}

	var lastErr error
	for retry := 0; retry < 2; retry++ {
		neighbors, err := c.getClosestPeersAlgorand(ctx, pid, rpi)
		if err == nil {
			// Success, even if neighbors is empty
			return neighbors, nil
		}
		lastErr = err
		// Basic retry, can be improved with more specific error handling
		log.WithError(err).WithFields(log.Fields{
			"peer":   pid.ShortString(),
			"bucket": bucket,
			"retry":  retry,
		}).Warnln("Failed getting closest peers, retrying...")
		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("getting closest peer with CPL %d for Algorand: %w", bucket, lastErr)
}

const KademliaMessageTag = "kd"

func (c *Crawler) getClosestPeersAlgorand(ctx context.Context, targetPeer peer.ID, key peer.ID) ([]*peer.AddrInfo, error) {
	c.algorandQueryMutex.Lock()
	defer c.algorandQueryMutex.Unlock()

	val, ok := c.driver.activeStreams.Load(targetPeer)
	if !ok {
		return nil, fmt.Errorf("no active stream for peer %s", targetPeer)
	}
	stream := val.(network.Stream)

	log.WithFields(log.Fields{
		"targetPeer": targetPeer.ShortString(),
		"key":        key.ShortString(),
	}).Debugln("Sending FIND_NODE message to Algorand peer")

	// 1. Construct the Kademlia FIND_NODE message.
	msg := pb.NewMessage(pb.Message_FIND_NODE, []byte(key), 0)

	// 2. Marshal the protobuf message.
	payload, err := proto.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal FIND_NODE message: %w", err)
	}

	// 3. Prepend the Algorand message tag.
	taggedPayload := append([]byte(KademliaMessageTag), payload...)

	// 4. Create and write the 4-byte length prefix.
	lenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBuf, uint32(len(taggedPayload)))
	if _, err := stream.Write(lenBuf); err != nil {
		c.driver.activeStreams.Delete(targetPeer)
		return nil, fmt.Errorf("writing length prefix: %w", err)
	}

	// 5. Write the tagged payload.
	if _, err := stream.Write(taggedPayload); err != nil {
		c.driver.activeStreams.Delete(targetPeer)
		return nil, fmt.Errorf("writing tagged payload: %w", err)
	}

	// --- Reading the Response ---

	// 1. Read the 4-byte length prefix.
	respLenBuf := make([]byte, 4)
	if _, err := io.ReadFull(stream, respLenBuf); err != nil {
		c.driver.activeStreams.Delete(targetPeer)
		return nil, fmt.Errorf("reading response length prefix: %w", err)
	}
	respLen := binary.BigEndian.Uint32(respLenBuf)
	if respLen > network.MessageSizeMax {
		c.driver.activeStreams.Delete(targetPeer)
		return nil, fmt.Errorf("response message too large: %d", respLen)
	}

	// 2. Read the full tagged response.
	taggedRespPayload := make([]byte, respLen)
	if _, err := io.ReadFull(stream, taggedRespPayload); err != nil {
		c.driver.activeStreams.Delete(targetPeer)
		return nil, fmt.Errorf("reading tagged response payload: %w", err)
	}

	// 3. Check the tag.
	if len(taggedRespPayload) < 2 {
		return nil, fmt.Errorf("response too short to contain a tag")
	}
	tag := string(taggedRespPayload[:2])
	if tag != KademliaMessageTag {
		return nil, fmt.Errorf("unexpected response tag: got '%s', want '%s'", tag, KademliaMessageTag)
	}

	// 4. Unmarshal the protobuf payload, *stripping the tag*.
	respMsg := new(pb.Message)
	if err := proto.Unmarshal(taggedRespPayload[2:], respMsg); err != nil {
		c.driver.activeStreams.Delete(targetPeer)
		return nil, fmt.Errorf("unmarshaling CLOSER_PEERS response: %w", err)
	}

	// --- Your existing logic ---
	log.WithFields(log.Fields{
		"targetPeer":   targetPeer.ShortString(),
		"responseType": respMsg.GetType(),
		"closerPeers":  len(respMsg.GetCloserPeers()),
	}).Debugln("Received response from Algorand peer")

	if respMsg.GetType() != pb.Message_FIND_NODE {
		return nil, fmt.Errorf("unexpected response type: %s", respMsg.GetType())
	}

	var addrInfos []*peer.AddrInfo
	for _, p := range respMsg.GetCloserPeers() {
		pid, err := peer.IDFromBytes(p.Id)
		if err != nil {
			log.WithError(err).Warnln("Failed to decode peer ID from protobuf")
			continue
		}

		maddrs := make([]ma.Multiaddr, len(p.Addrs))
		for i, addrBytes := range p.Addrs {
			maddr, err := ma.NewMultiaddrBytes(addrBytes)
			if err != nil {
				log.WithError(err).Warnln("Failed to parse multiaddr from protobuf")
				continue
			}
			maddrs[i] = maddr
		}

		addrInfos = append(addrInfos, &peer.AddrInfo{ID: pid, Addrs: maddrs})
	}

	return addrInfos, nil
}