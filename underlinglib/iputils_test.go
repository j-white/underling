package underlinglib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPAddressRange(t *testing.T) {
	// A single IPv4 address
	VerifyExpectedIPAddressRange(t, "127.0.0.1", "127.0.0.1", []string{"127.0.0.1"})
	// A simple IPv4 range
	VerifyExpectedIPAddressRange(t, "127.0.0.1", "127.0.0.3", []string{"127.0.0.1", "127.0.0.2", "127.0.0.3"})
	// A single IPv6 address
	//VerifyExpectedIPAddressRange(t, "::1", "::1", []string{"::1"})
}

func VerifyExpectedIPAddressRange(t *testing.T, begin string, end string, expected []string) {
	gen, _ := IPAddressRange(begin, end)

	var ips []string
	for addr := gen(); addr != nil; addr = gen() {
		ips = append(ips, addr.IP.String())
	}

	assert.Equal(t, expected, ips)
}
