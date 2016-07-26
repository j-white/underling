package underlinglib

import (
	"gopkg.in/yaml.v2"
)

type UnderlingConfig struct {
	Minion struct {
		Location string
		Id       string
	}
	OpenNMS struct {
		Url      string
		Mq       string
		Username string
		Password string
	}
	Underling struct {
		Detectors struct {
			FailForUnknownDetectors bool
		}
	}
}

func GetConfig(data []byte) (UnderlingConfig, error) {
	config := UnderlingConfig{}
	err := yaml.Unmarshal(data, &config)
	return config, err
}
