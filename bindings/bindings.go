package bindings

import (
	"github.com/j-white/underling/underlinglib"
)

var underling *underlinglib.Underling

type Config struct {
	Location string
	Id string
	Mq string
}

func StartUnderling(conf Config) {
	if (underling != nil) {
		return
	}
	
	// Convert the flat configuration to the UnderlingConfig
	underlingConf := new(underlinglib.UnderlingConfig)
	underlingConf.Minion.Id = conf.Id
	underlingConf.Minion.Location = conf.Location
	underlingConf.OpenNMS.Mq = conf.Mq
	
	underling = new(underlinglib.Underling)
	underling.Start(*underlingConf)
}

func StopUnderling() {
	underling.Stop()
	underling = nil
}
