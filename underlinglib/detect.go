package underlinglib

func IcmpDetect(request DetectorRequestDTO) DetectorResponseDTO {
	return DetectorResponseDTO{Detected: true}
}

func SnmpDetect(request DetectorRequestDTO) DetectorResponseDTO {
	return DetectorResponseDTO{Detected: true}
}

func Detect(request DetectorRequestDTO) DetectorResponseDTO {
	switch request.ClassName {
	case "org.opennms.netmgt.provision.detector.snmp.SnmpDetector":
		return SnmpDetect(request)
	case "org.opennms.netmgt.provision.detector.icmp.IcmpDetector":
		return IcmpDetect(request)
	default:
		return DetectorResponseDTO{Detected: false, FailureMessage: "Unsupported detector class " + request.ClassName}
	}
}
