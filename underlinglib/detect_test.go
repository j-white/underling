package underlinglib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIcmpDetect(t *testing.T) {
	request := DetectorRequestDTO{Address: "203.0.113.1"}
	response := IcmpDetect(request)
	assert.Equal(t, false, response.Detected)
	assert.Equal(t, "", response.FailureMessage)
}
