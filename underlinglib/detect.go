package underlinglib

import (
	"strings"
)

type DetectorRpcModule struct {
}

func (detect DetectorRpcModule) GetId() (id string) {
	return "Detect"
}

func (detect DetectorRpcModule) HandleRequest(requestBody string) (responseBody string) {
	request := DetectorRequestDTO{}
	UnmarshalFromXml(strings.NewReader(requestBody), &request)
	response := DetectExec(request)
	responseBody, _ = MarshalToXml(response)
	return responseBody
}

func DetectExec(request DetectorRequestDTO) DetectorResponseDTO {
	// TODO: Can we dynamically register these in Go?
	// In Java one could use the ServiceLoader
	switch request.ClassName {
	case "org.opennms.netmgt.provision.detector.snmp.SnmpDetector":
		return SnmpDetect(request)
	case "org.opennms.netmgt.provision.detector.icmp.IcmpDetector":
		return IcmpDetect(request)
	default:
		return DetectorResponseDTO{Detected: false}
	}
}

func IcmpDetect(request DetectorRequestDTO) DetectorResponseDTO {
	p := NewPinger()
	rtt, err := p.Ping(request.Address, DefaultPingRetries, DefaultPingTimeout)
	if err != nil {
		return DetectorResponseDTO{Detected: false, FailureMessage: err.Error()}
	} else if rtt > 0 {
		return DetectorResponseDTO{Detected: true}
	} else {
		return DetectorResponseDTO{Detected: false}
	}
}

func SnmpDetect(request DetectorRequestDTO) DetectorResponseDTO {
	// TODO: Implement the detector
	return DetectorResponseDTO{Detected: true}
}
