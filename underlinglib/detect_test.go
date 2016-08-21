package underlinglib

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIcmpDetect(t *testing.T) {
	if os.Getenv("CIRCLECI") != "" {
		t.Skip("skipping test; ICMP not supported in CirclCI.")
	}

	request := DetectorRequestDTO{Address: "127.0.0.1"}
	response := IcmpDetect(request)
	assert.Equal(t, "", response.FailureMessage)
	assert.Equal(t, true, response.Detected)
}
