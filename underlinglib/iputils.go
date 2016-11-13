package underlinglib

import (
	"encoding/binary"
	"net"
	"strings"
)

func resolve(address string) (*net.IPAddr, error) {
	netProto := "ip4:icmp"
	if strings.Index(address, ":") != -1 {
		netProto = "ip6:ipv6-icmp"
	}
	return net.ResolveIPAddr(netProto, address)
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func IPAddressRange(begin string, end string) (func() *net.IPAddr, error) {
	firstip, err := resolve(begin)
	if err != nil {
		return nil, err
	}
	lastip, err := resolve(end)
	if err != nil {
		return nil, err
	}

	first := ip2int(firstip.IP)
	last := ip2int(lastip.IP)

	current := first - 1
	return func() *net.IPAddr {
		current += 1
		if current > last {
			return nil
		} else {
			currentip, _ := net.ResolveIPAddr(firstip.Network(), int2ip(current).String())
			return currentip
		}
	}, nil
}
