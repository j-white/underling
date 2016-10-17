package underlinglib

import (
	"strings"
)

const (
	SERVICE_UNKNOWN = 0
	SERVICE_AVAILABLE = 1
	SERVICE_UNAVAILABLE = 2
	SERVICE_UNRESPONSIVE = 3
	SERVICE_UNKNOWN_NAME = "Unknown"
	SERVICE_AVAILABLE_NAME = "Up"
	SERVICE_UNAVAILABLE_NAME = "Down"
	SERVICE_UNRESPONSIVE_NAME = "Unresponsive"
)

type PollerRpcModule struct {
}

func (poller PollerRpcModule) GetId() (id string) {
	return "Poller"
}

func (poller PollerRpcModule) HandleRequest(requestBody string) (responseBody string) {
	request := PollerRequestDTO{}
	UnmarshalFromXml(strings.NewReader(requestBody), &request)
	response := MonitorExec(request)
	responseBody, _ = MarshalToXml(response)
	return responseBody
}

func MonitorExec(request PollerRequestDTO) PollerResponseDTO {
	// TODO: Can we dynamically register these in Go?
	// In Java one could use the ServiceLoader
	status := PollerStatusDTO{}
	switch request.ClassName {
	case "org.opennms.netmgt.poller.monitors.IcmpMonitor":
		status = IcmpMonitor(request)
	}
	return PollerResponseDTO{Status: status}
}

func IcmpMonitor(request PollerRequestDTO) PollerStatusDTO {
	p := NewPinger()
	rtt, err := p.Ping(request.Address, DefaultPingRetries, DefaultPingTimeout)
	if err != nil {
		return PollerStatusDTO{Code: SERVICE_UNKNOWN, Name: SERVICE_UNKNOWN_NAME, Reason: err.Error()}
	} else if rtt > 0 {
		return PollerStatusDTO{Code: SERVICE_AVAILABLE, Name: SERVICE_AVAILABLE_NAME}
	} else {
		return PollerStatusDTO{Code: SERVICE_UNAVAILABLE, Name: SERVICE_UNAVAILABLE_NAME, Reason: "Did not respond before timeout."}
	}
}
