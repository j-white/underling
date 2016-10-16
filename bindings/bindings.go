package bindings

import (
	"github.com/j-white/underling/underlinglib"
)

var underling *underlinglib.Underling

func StartUnderling(id string, location string, mq string) {
	if underling != nil {
		return
	}

	// Convert the flat configuration to the UnderlingConfig
	underlingConf := new(underlinglib.UnderlingConfig)
	underlingConf.Minion.Id = id
	underlingConf.Minion.Location = location
	underlingConf.OpenNMS.Mq = mq

	underling = new(underlinglib.Underling)
	underling.Start(*underlingConf)
}

func StopUnderling() {
	underling.Stop()
	underling = nil
}
