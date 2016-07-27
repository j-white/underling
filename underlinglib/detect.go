package underlinglib

import (
	"github.com/tatsushid/go-fastping"
	"net"
	"strings"
	"time"
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
	switch request.ClassName {
	case "org.opennms.netmgt.provision.detector.snmp.SnmpDetector":
		return SnmpDetect(request)
	case "org.opennms.netmgt.provision.detector.icmp.IcmpDetector":
		return IcmpDetect(request)
	default:
		return DetectorResponseDTO{Detected: false, FailureMessage: "Unsupported detector class " + request.ClassName}
	}
}

func IcmpDetect(request DetectorRequestDTO) DetectorResponseDTO {
	channel := make(chan DetectorResponseDTO)

	p := fastping.NewPinger()
	p.Network("udp") // Use UDP sockets
	ra, err := net.ResolveIPAddr("ip4:icmp", request.Address)
	if err != nil {
		return DetectorResponseDTO{Detected: false, FailureMessage: err.Error()}
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		channel <- DetectorResponseDTO{Detected: true}
	}
	p.OnIdle = func() {
		channel <- DetectorResponseDTO{Detected: false}
	}
	err = p.Run()
	if err != nil {
		return DetectorResponseDTO{Detected: false, FailureMessage: err.Error()}
	}

	return <-channel
}

func SnmpDetect(request DetectorRequestDTO) DetectorResponseDTO {
	return DetectorResponseDTO{Detected: true}
}
