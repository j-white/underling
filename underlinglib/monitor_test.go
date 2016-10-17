package underlinglib

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIcmpMonitor(t *testing.T) {
	if os.Getenv("CIRCLECI") != "" {
		t.Skip("skipping test; ICMP not supported in CirclCI.")
	}

	request := PollerRequestDTO{Address: "127.0.0.1"}
	status := IcmpMonitor(request)
	assert.Equal(t, 1, status.Code)
	assert.Equal(t, "Up", status.Name)
}
