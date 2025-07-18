package utils

import (
	"fmt"
	"net"
	"strings"
	"syscall"
	"testing"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/stretchr/testify/require"
)

// MaddrsToAddrs maps a slice of multi addresses to their string representation.
func MaddrsToAddrs(maddrs []ma.Multiaddr) []string {
	addrs := make([]string, len(maddrs))
	for i, maddr := range maddrs {
		addrs[i] = maddr.String()
	}
	return addrs
}

// EllipsizeMaddr shortens the certhashes
func EllipsizeMaddr(input string, maxLength int) string {
	parts := strings.Split(input, "/")
	for i := 0; i < len(parts); i++ {
		// If the part is "certhash", the next part (if it exists) is the one to shorten
		if parts[i] == "certhash" && i+1 < len(parts) {
			parts[i+1] = ellipsize(parts[i+1], maxLength)
		}
	}
	return strings.Join(parts, "/")
}

// ellipsize truncates a string by keeping the first and last parts intact and replacing the middle with "..."
func ellipsize(input string, maxLength int) string {
	if len(input) <= maxLength {
		return input // No need to truncate
	}

	partLength := (maxLength - 2) / 2
	if partLength < 1 {
		partLength = 1
	}

	start := input[:partLength]
	end := input[len(input)-partLength:]

	return fmt.Sprintf("%s..%s", start, end)
}

// AddrsToMaddrs maps a slice of addresses to their multiaddress representation.
func AddrsToMaddrs(addrs []string) ([]ma.Multiaddr, error) {
	maddrs := make([]ma.Multiaddr, len(addrs))
	for i, addr := range addrs {
		maddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}
		maddrs[i] = maddr
	}

	return maddrs, nil
}

// MustMultiaddr returns parses the given multi address string and stops the
// test with an error if that fails.
func MustMultiaddr(t testing.TB, maddrStr string) ma.Multiaddr {
	maddr, err := ma.NewMultiaddr(maddrStr)
	require.NoError(t, err)
	return maddr
}

// AddrInfoFilterPrivateMaddrs strips private multiaddrs from the given peer
// address information. It returns two new AddrInfo structs. The first contains
// only non-private multi addresses and the second return value contains only
// private multi addresses.
func AddrInfoFilterPrivateMaddrs(pi peer.AddrInfo) (peer.AddrInfo, peer.AddrInfo) {
	keep := peer.AddrInfo{
		ID:    pi.ID,
		Addrs: []ma.Multiaddr{},
	}

	drop := peer.AddrInfo{
		ID:    pi.ID,
		Addrs: []ma.Multiaddr{},
	}

	// Just keep public multi addresses
	for _, maddr := range pi.Addrs {
		if manet.IsPrivateAddr(maddr) {
			drop.Addrs = append(drop.Addrs, maddr)
			continue
		}
		keep.Addrs = append(keep.Addrs, maddr)
	}

	return keep, drop
}

// AddrInfoFilterPublicMaddrs strips public multiaddrs from the given peer
// address information. It returns two new AddrInfo structs. The first contains
// only non-public multi addresses and the second return value contains only
// public multi addresses.
func AddrInfoFilterPublicMaddrs(pi peer.AddrInfo) (peer.AddrInfo, peer.AddrInfo) {
	keep := peer.AddrInfo{
		ID:    pi.ID,
		Addrs: []ma.Multiaddr{},
	}

	drop := peer.AddrInfo{
		ID:    pi.ID,
		Addrs: []ma.Multiaddr{},
	}

	// Just keep public multi addresses
	for _, maddr := range pi.Addrs {
		if manet.IsPublicAddr(maddr) {
			drop.Addrs = append(drop.Addrs, maddr)
			continue
		}
		keep.Addrs = append(keep.Addrs, maddr)
	}

	return keep, drop
}

// FilterPrivateMaddrs strips private multiaddrs from the given peer address information.
func FilterPrivateMaddrs(maddrs []ma.Multiaddr) []ma.Multiaddr {
	var filtered []ma.Multiaddr
	for _, maddr := range maddrs {
		if manet.IsPrivateAddr(maddr) {
			continue
		}
		filtered = append(filtered, maddr)
	}
	return filtered
}

// FilterPublicMaddrs strips public multiaddrs from the given peer address information.
func FilterPublicMaddrs(maddrs []ma.Multiaddr) []ma.Multiaddr {
	var filtered []ma.Multiaddr
	for _, maddr := range maddrs {
		if manet.IsPublicAddr(maddr) {
			continue
		}
		filtered = append(filtered, maddr)
	}
	return filtered
}

// MergeMaddrs takes two slices of multi addresses and merges them into a single
// one.
func MergeMaddrs(maddrSet1 []ma.Multiaddr, maddrSet2 []ma.Multiaddr) []ma.Multiaddr {
	maddrSetOut := make(map[string]ma.Multiaddr, len(maddrSet1))
	for _, maddr := range maddrSet1 {
		if _, found := maddrSetOut[string(maddr.Bytes())]; found {
			continue
		}
		maddrSetOut[string(maddr.Bytes())] = maddr
	}

	for _, maddr := range maddrSet2 {
		if _, found := maddrSetOut[string(maddr.Bytes())]; found {
			continue
		}
		maddrSetOut[string(maddr.Bytes())] = maddr
	}

	maddrsOut := make([]ma.Multiaddr, 0, len(maddrSetOut))
	for _, maddr := range maddrSetOut {
		maddrsOut = append(maddrsOut, maddr)
	}

	return maddrsOut
}

// IsResourceLimitExceeded returns true if the given error represents an error
// related to a limit of the local resource manager.
func IsResourceLimitExceeded(err error) bool {
	return err != nil && strings.HasSuffix(err.Error(), network.ErrResourceLimitExceeded.Error())
}

// nebulaIdentityScheme is an always-valid ID scheme. When a new [enode.Node] is
// constructed, the Verify method won't check the signature, and we just assume
// the record is valid. However, the NodeAddr method returns the correct node
// identifier.
type nebulaIdentityScheme struct{}

// Verify doesn't check the signature or anything. It assumes all records to be
// valid.
func (nebulaIdentityScheme) Verify(r *enr.Record, sig []byte) error {
	return nil
}

// NodeAddr returns the node's ID. The logic is copied from the [enode.V4ID]
// implementation.
func (nebulaIdentityScheme) NodeAddr(r *enr.Record) []byte {
	var pubkey enode.Secp256k1
	err := r.Load(&pubkey)
	if err != nil {
		return nil
	}
	buf := make([]byte, 64)
	math.ReadBits(pubkey.X, buf[:32])
	math.ReadBits(pubkey.Y, buf[32:])
	return crypto.Keccak256(buf)
}

// GetUDPBufferSize reads the receive and send buffer sizes from the system
func GetUDPBufferSize(conn *net.UDPConn) (rcvbuf int, sndbuf int, err error) {
	rawConn, err := conn.SyscallConn()
	if err != nil {
		return 0, 0, err
	}

	var (
		rcverr error
		snderr error
	)
	err = rawConn.Control(func(fd uintptr) {
		rcvbuf, rcverr = syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF)
		sndbuf, snderr = syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_SNDBUF)
	})
	if rcverr != nil {
		err = rcverr
	} else if snderr != nil {
		err = snderr
	}

	return
}
