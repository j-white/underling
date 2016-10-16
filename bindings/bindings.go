package bindings

import (
	"github.com/j-white/underling/underlinglib"
)

var underling *underlinglib.Underling

func StartUnderling(mq string, opennmsId string, location string, minionId string) {
	if underling != nil {
		return
	}

	// Convert the flat configuration to the UnderlingConfig
	underlingConf := new(underlinglib.UnderlingConfig)
	underlingConf.OpenNMS.Mq = mq
	underlingConf.OpenNMS.Id = opennmsId
	underlingConf.Minion.Location = location
	underlingConf.Minion.Id = minionId

	underling = new(underlinglib.Underling)
	underling.Start(*underlingConf)
}

func StopUnderling() {
	underling.Stop()
	underling = nil
}
