package underlinglib

import (
	"strings"
)

type PingRpcModule struct {
}

func (ping PingRpcModule) GetId() (id string) {
	return "PING"
}

func (ping PingRpcModule) HandleRequest(requestBody string) (responseBody string) {
	request := PingRequestDTO{}
	UnmarshalFromXml(strings.NewReader(requestBody), &request)
	response := DoPing(request)
	responseBody, _ = MarshalToXml(response)
	return responseBody
}

func DoPing(request PingRequestDTO) PingResponseDTO {
	p := NewPinger()
	rtt, err := p.Ping(request.Address, request.Retries, request.Timeout)
	if err != nil {
		// Error
		return PingResponseDTO{}
	} else if rtt > 0 {
		// OK
		return PingResponseDTO{RTT: float64(rtt.Nanoseconds())}
	} else {
		// Timeout
		return PingResponseDTO{}
	}
}

type PingSweepRpcModule struct {
}

func (ping PingSweepRpcModule) GetId() (id string) {
	return "PING-SWEEP"
}

func (ping PingSweepRpcModule) HandleRequest(requestBody string) (responseBody string) {
	request := PingSweepRequestDTO{}
	UnmarshalFromXml(strings.NewReader(requestBody), &request)
	response := DoPingSweep(request)
	responseBody, _ = MarshalToXml(response)
	return responseBody
}

func DoPingSweep(request PingSweepRequestDTO) PingSweepResponseDTO {
	response := PingSweepResponseDTO{}
	for _, ipRange := range request.Ranges {
		gen, _ := IPAddressRange(ipRange.Begin, ipRange.End)
		for addr := gen(); addr != nil; addr = gen() {
			p := NewPinger()
			rtt, err := p.Ping(addr.IP.String(), ipRange.Retries, ipRange.Timeout)
			if err != nil {
				// Error
			} else if rtt > 0 {
				// OK
				response.Results = append(response.Results, PingSweepResultDTO{Address: addr.IP.String(), RTT: float64(rtt.Nanoseconds())})
			} else {
				// Timeout
			}
		}
	}
	return response
}
