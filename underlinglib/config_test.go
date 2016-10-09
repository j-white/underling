package underlinglib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseConfigFromYaml(t *testing.T) {
	var yaml = `
minion:
  location: UNDERLING
  id: 73d292b8-9db7-48be-8c84-c736b97fc4e7
opennms:
  url: https://127.0.0.1:8980/opennms
  mq: localhost:61613
  id: OpenNMS
  username: admin
  password: admin
underling:
  detectors:
    fail_for_unknown_detectors: false
`

	config, err := GetConfig([]byte(yaml))
	fmt.Print(err)
	assert.Equal(t, "UNDERLING", config.Minion.Location)
	assert.Equal(t, "73d292b8-9db7-48be-8c84-c736b97fc4e7", config.Minion.Id)
	assert.Equal(t, "https://127.0.0.1:8980/opennms", config.OpenNMS.Url)
	assert.Equal(t, "localhost:61613", config.OpenNMS.Mq)
	assert.Equal(t, "OpenNMS", config.OpenNMS.Id)
	assert.Equal(t, "admin", config.OpenNMS.Username)
	assert.Equal(t, "admin", config.OpenNMS.Password)
}
