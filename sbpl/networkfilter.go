package sbpl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/syumai/sbx/internal/sliceutil"
)

func NewNetworkFilterAddress(host string, port string) (*NetworkFilterAddress, error) {
	if host != "*" && host != "localhost" {
		return nil, fmt.Errorf("invalid host: %s", host)
	}
	portNum, _ := strconv.Atoi(port)
	if port != "*" && (portNum < 0 || portNum > 65535) {
		return nil, fmt.Errorf("invalid port: %s", port)
	}
	return &NetworkFilterAddress{Host: host, Port: port}, nil
}

type NetworkFilterAddress struct {
	Host string
	Port string
}

func (a *NetworkFilterAddress) String() string {
	return fmt.Sprintf(`"%s:%s"`, a.Host, a.Port)
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
