package underlinglib

import (
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
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

type response struct {
	addr *net.IPAddr
	rtt  time.Duration
}

func IcmpDetect(request DetectorRequestDTO) DetectorResponseDTO {
	p := fastping.NewPinger()
	p.Network("udp")

	netProto := "ip4:icmp"
	if strings.Index(request.Address, ":") != -1 {
		netProto = "ip6:ipv6-icmp"
	}
	ra, err := net.ResolveIPAddr(netProto, request.Address)
	if err != nil {
		return DetectorResponseDTO{Detected: false, FailureMessage: err.Error()}
	}
	p.AddIPAddr(ra)

	onRecv, onIdle := make(chan *response), make(chan bool)
	p.OnRecv = func(addr *net.IPAddr, t time.Duration) {
		onRecv <- &response{addr: addr, rtt: t}
	}
	p.OnIdle = func() {
		onIdle <- true
	}

	p.MaxRTT = time.Second
	p.RunLoop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	result := DetectorResponseDTO{Detected: false, FailureMessage: "oups"}
loop:
	for {
		select {
		case <-c:
			result = DetectorResponseDTO{Detected: false, FailureMessage: "Interrupted."}
			break loop
		case <-onRecv:
			result = DetectorResponseDTO{Detected: true}
			break loop
		case <-onIdle:
			result = DetectorResponseDTO{Detected: false}
			break loop
		case <-p.Done():
			if err != nil {
				result = DetectorResponseDTO{Detected: false, FailureMessage: err.Error()}
			} else {
				result = DetectorResponseDTO{Detected: false, FailureMessage: "Unkown error."}
			}
			break loop
		}
	}
	signal.Stop(c)
	p.Stop()
	return result
}

func SnmpDetect(request DetectorRequestDTO) DetectorResponseDTO {
	return DetectorResponseDTO{Detected: true}
}
