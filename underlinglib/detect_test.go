package underlinglib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIcmpDetect(t *testing.T) {
	request := DetectorRequestDTO{Address: "127.0.0.1"}
	response := IcmpDetect(request)
	assert.Equal(t, true, response.Detected)
}
