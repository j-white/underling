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
	return DetectorResponseDTO{Detected: true}
}

func SnmpDetect(request DetectorRequestDTO) DetectorResponseDTO {
	return DetectorResponseDTO{Detected: true}
}
