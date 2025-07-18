package db

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

type CrawlState string

const (
	CrawlStateStarted   CrawlState = "started"
	CrawlStateSucceeded CrawlState = "succeeded"
	CrawlStateCancelled CrawlState = "cancelled"
	CrawlStateFailed    CrawlState = "failed"
)

type VisitType string

const (
	VisitTypeDial  VisitType = "dial"
	VisitTypeCrawl VisitType = "crawl"
)

func (v VisitType) String() string {
	return string(v)
}

type SealCrawlArgs struct {
	Crawled    int
	Dialable   int
	Undialable int
	Remaining  int
	State      CrawlState
}

type VisitArgs struct {
	PeerID           peer.ID
	DiscoveryPrefix  uint64
	AgentVersion     string
	Protocols        []string
	DialMaddrs       []ma.Multiaddr
	FilteredMaddrs   []ma.Multiaddr
	ExtraMaddrs      []ma.Multiaddr
	ListenMaddrs     []ma.Multiaddr
	DialErrors       []string
	ConnectMaddr     ma.Multiaddr
	DialDuration     time.Duration
	ConnectDuration  time.Duration
	CrawlDuration    time.Duration
	VisitStartedAt   time.Time
	VisitEndedAt     time.Time
	ConnectErrorStr  string
	CrawlErrorStr    string
	VisitType        VisitType
	Neighbors        []peer.ID
	NeighborPrefixes []uint64
	ErrorBits        uint16
	Properties       json.RawMessage
}

type Client interface {
	io.Closer
	// InitCrawl initializes a new crawl instance in the database.
	// The clients are responsible for tracking the crawl's ID and associate
	// later database queries with it. This is necessary because different
	// database engines have different types of IDs. ClickHouse commonly uses string
	// IDs and Postgres uses integers. Making the [Client] interface generic
	// on that ID would complicate the code a lot, so we require Clients to
	// keep state. This is added complexity traded for code clarity. It's a trade-
	// off and IMO this is less bad.
	InitCrawl(ctx context.Context, version string) error

	// SealCrawl marks the crawl (that the Client tracks internally) as done.
	SealCrawl(ctx context.Context, args *SealCrawlArgs) error

	// QueryBootstrapPeers fetches peers from the database that can be used
	// for bootstrapping into the network. The result will contain from zero up
	// to limit entries.
	QueryBootstrapPeers(ctx context.Context, limit int) ([]peer.AddrInfo, error)

	// InsertVisit TODO
	InsertVisit(ctx context.Context, args *VisitArgs) error

	// InsertCrawlProperties TODO
	InsertCrawlProperties(ctx context.Context, properties map[string]map[string]int) error

	// SelectPeersToProbe TODO
	SelectPeersToProbe(ctx context.Context) ([]peer.AddrInfo, error)

	// Flush instructs the client to write all cached data to the database.
	// Client implementations may cache and batch inserts. Flush tells the
	// client to insert everything that's pending.
	Flush(ctx context.Context) error
}

var (
	_ Client = (*PostgresClient)(nil)
	_ Client = (*NoopClient)(nil)
	_ Client = (*JSONClient)(nil)
	_ Client = (*ClickHouseClient)(nil)
)
