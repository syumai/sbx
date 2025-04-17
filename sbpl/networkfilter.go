package sbpl

import (
	"fmt"
	"strings"

	"github.com/syumai/sbx/internal/sliceutil"
)

func NewNetworkFilterAddress(host string, port int) (*NetworkFilterAddress, error) {
	if host != "*" && host != "localhost" {
		return nil, fmt.Errorf("invalid host: %s", host)
	}
	if port < 0 || port > 65535 {
		return nil, fmt.Errorf("invalid port: %d", port)
	}
	return &NetworkFilterAddress{Host: host, Port: port}, nil
}

type NetworkFilterAddress struct {
	Host string
	Port int
}

func (a *NetworkFilterAddress) String() string {
	return fmt.Sprintf(`"%s:%d"`, a.Host, a.Port)
}

func NewNetworkFilter(isLocal bool, protocol NetworkFilterProtocol, addresses []*NetworkFilterAddress) *NetworkFilter {
	return &NetworkFilter{
		IsLocal:   isLocal,
		Protocol:  protocol,
		Addresses: addresses,
	}
}

type NetworkFilter struct {
	IsLocal   bool
	Protocol  NetworkFilterProtocol
	Addresses []*NetworkFilterAddress
}

type NetworkFilterProtocol int

const (
	NetworkFilterProtocolUnknown NetworkFilterProtocol = iota
	NetworkFilterProtocolIP
	NetworkFilterProtocolTCP
	NetworkFilterProtocolUDP
)

func (p NetworkFilterProtocol) String() string {
	switch p {
	case NetworkFilterProtocolIP:
		return "ip"
	case NetworkFilterProtocolTCP:
		return "tcp"
	case NetworkFilterProtocolUDP:
		return "udp"
	}
	panic(fmt.Sprintf("unexpected network filter protocol: %d", p))
}

func (f *NetworkFilter) String() string {
	localOrRemote := "local"
	if !f.IsLocal {
		localOrRemote = "remote"
	}
	addresses := strings.Join(sliceutil.MapStringer(f.Addresses), ", ")
	return fmt.Sprintf("(%s %s %s)", localOrRemote, f.Protocol, addresses)
}
