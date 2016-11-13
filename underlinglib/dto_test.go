package underlinglib

import (
	"encoding/xml"
	"testing"
)

func TestMarshalUnmarshalSNMPRequestDTO(t *testing.T) {
	expectedXml := `<snmp-request location="dc2" description="some random oids">
  <agent>
    <address>192.168.0.2</address>
    <readCommunity>public</readCommunity>
    <version>1</version>
  </agent>
  <get correlation-id="42">
    <oid>.1.3.6.1.2.1.3.1.3.0</oid>
  </get>
  <walk correlation-id="43" max-repetitions="4">
    <oid>.1.3.6.1.2.1.4.34.1.3</oid>
    <oid>.1.3.6.1.2.1.4.34.1.5</oid>
    <oid>.1.3.6.1.2.1.4.34.1.4</oid>
  </walk>
  <walk correlation-id="44" instance=".0">
    <oid>.1.3.6.1.2.1.3.1.3</oid>
  </walk>
</snmp-request>`

	agent := SNMPAgentDTO{
		XMLName:       xml.Name{Space: "", Local: "agent"},
		Address:       "192.168.0.2",
		ReadCommunity: "public",
		Version:       1,
	}

	get := SNMPGetRequestDTO{
		XMLName:       xml.Name{Space: "", Local: "get"},
		CorrelationID: "42",
		OIDs:          []string{".1.3.6.1.2.1.3.1.3.0"},
	}

	walk_with_many_oids := SNMPWalkRequestDTO{
		XMLName:        xml.Name{Space: "", Local: "walk"},
		CorrelationID:  "43",
		MaxRepetitions: 4,
		OIDs:           []string{".1.3.6.1.2.1.4.34.1.3", ".1.3.6.1.2.1.4.34.1.5", ".1.3.6.1.2.1.4.34.1.4"},
	}

	walk_with_instance := SNMPWalkRequestDTO{
		XMLName:       xml.Name{Space: "", Local: "walk"},
		CorrelationID: "44",
		Instance:      ".0",
		OIDs:          []string{".1.3.6.1.2.1.3.1.3"},
	}

	expectedRequest := SNMPRequestDTO{
		XMLName:     xml.Name{Space: "", Local: "snmp-request"},
		Location:    "dc2",
		Description: "some random oids",
		Agent:       agent,
		Gets:        []SNMPGetRequestDTO{get},
		Walks:       []SNMPWalkRequestDTO{walk_with_many_oids, walk_with_instance},
	}

	MarshalUnmarshal(t, expectedRequest, expectedXml, &(SNMPRequestDTO{}))
}

func TestMarshalUnmarshalSNMPResponseDTO(t *testing.T) {
	expectedXml := `<snmp-response>
  <response correlation-id="42">
    <result>
      <base>.1.3.6.1.2</base>
      <instance>1.3.6.1.2.1.4.34.1.3.1.2.3.4</instance>
      <value type="70">Cg==</value>
    </result>
    <result>
      <base>.1.3.6.1.2</base>
      <instance>1.3.6.1.2.1.4.34.1.3.1.2.3.5</instance>
      <value type="69">!Cg==</value>
    </result>
  </response>
</snmp-response>`

	first_result := SNMPResultDTO{
		XMLName:  xml.Name{Space: "", Local: "result"},
		Base:     ".1.3.6.1.2",
		Instance: "1.3.6.1.2.1.4.34.1.3.1.2.3.4",
		Value: SNMPValueDTO{
			XMLName: xml.Name{Space: "", Local: "value"},
			Type:    70,
			Value:   "Cg==",
		},
	}

	second_result := SNMPResultDTO{
		XMLName:  xml.Name{Space: "", Local: "result"},
		Base:     ".1.3.6.1.2",
		Instance: "1.3.6.1.2.1.4.34.1.3.1.2.3.5",
		Value: SNMPValueDTO{
			XMLName: xml.Name{Space: "", Local: "value"},
			Type:    69,
			Value:   "!Cg==",
		},
	}

	response := SNMPResponseDTO{
		XMLName:       xml.Name{Space: "", Local: "response"},
		CorrelationID: "42",
		Results:       []SNMPResultDTO{first_result, second_result},
	}

	expectedResponse := SNMPMultiResponseDTO{
		XMLName:   xml.Name{Space: "", Local: "snmp-response"},
		Responses: []SNMPResponseDTO{response},
	}

	MarshalUnmarshal(t, expectedResponse, expectedXml, &(SNMPMultiResponseDTO{}))
}

func TestMarshalUnmarshalDetectorRequestDTO(t *testing.T) {
	expectedXml := `<detector-request location="MINION" class-name="org.opennms.netmgt.provision.detector.icmp.IcmpDetector" address="127.0.0.1">
  <detector-attribute key="port">8980</detector-attribute>
  <runtime-attribute key="password">foo</runtime-attribute>
</detector-request>`

	detectorAttribute := DetectorAttributeDTO{
		XMLName: xml.Name{Space: "", Local: "detector-attribute"},
		Key:     "port",
		Value:   "8980",
	}

	runtimeAttribute := RuntimeAttributeDTO{
		XMLName: xml.Name{Space: "", Local: "runtime-attribute"},
		Key:     "password",
		Value:   "foo",
	}

	expectedRequest := DetectorRequestDTO{
		XMLName:            xml.Name{Space: "", Local: "detector-request"},
		Location:           "MINION",
		ClassName:          "org.opennms.netmgt.provision.detector.icmp.IcmpDetector",
		Address:            "127.0.0.1",
		DetectorAttributes: []DetectorAttributeDTO{detectorAttribute},
		RuntimeAttributes:  []RuntimeAttributeDTO{runtimeAttribute},
	}

	MarshalUnmarshal(t, expectedRequest, expectedXml, &(DetectorRequestDTO{}))
}

func TestMarshalUnmarshalDetectorResponseDTO(t *testing.T) {
	expectedXml := `<detector-response detected="true" failure-message="classCast exception">
  <attribute key="vendor">MOO</attribute>
</detector-response>`

	attribute := AttributeDTO{
		XMLName: xml.Name{Space: "", Local: "attribute"},
		Key:     "vendor",
		Value:   "MOO",
	}

	expectedResponse := DetectorResponseDTO{
		XMLName:        xml.Name{Space: "", Local: "detector-response"},
		Detected:       true,
		FailureMessage: "classCast exception",
		Attributes:     []AttributeDTO{attribute},
	}

	MarshalUnmarshal(t, expectedResponse, expectedXml, &(DetectorResponseDTO{}))
}

func TestMarshalUnmarshalMinionIdentityDTO(t *testing.T) {
	expectedXml := `<minion>
  <id>73d292b8-9db7-48be-8c84-c736b97fc4e7</id>
  <location>UNDERLING</location>
</minion>`

	expectedIdentity := MinionIdentityDTO{
		XMLName:  xml.Name{Space: "", Local: "minion"},
		Id:       "73d292b8-9db7-48be-8c84-c736b97fc4e7",
		Location: "UNDERLING",
	}

	MarshalUnmarshal(t, expectedIdentity, expectedXml, &(MinionIdentityDTO{}))
}

func TestMarshalUnmarshalPollerRequestDTO(t *testing.T) {
	expectedXml := `<poller-request location="MINION" class-name="org.opennms.netmgt.poller.monitors.IcmpMonitor" address="127.0.0.1">
  <attribute key="port" value="18980"></attribute>
</poller-request>`

	attribute := PollerAttributeDTO{
		XMLName: xml.Name{Space: "", Local: "attribute"},
		Key:     "port",
		Value:   "18980",
	}

	expectedRequest := PollerRequestDTO{
		XMLName:    xml.Name{Space: "", Local: "poller-request"},
		Location:   "MINION",
		ClassName:  "org.opennms.netmgt.poller.monitors.IcmpMonitor",
		Address:    "127.0.0.1",
		Attributes: []PollerAttributeDTO{attribute},
	}

	MarshalUnmarshal(t, expectedRequest, expectedXml, &(PollerRequestDTO{}))
}

func TestMarshalUnmarshalPollerResponseDTO(t *testing.T) {
	expectedXml := `<poller-response>
  <poll-status code="1" name="Up" time="1969-12-31T19:00:00-05:00"></poll-status>
</poller-response>`

	status := PollerStatusDTO{
		XMLName: xml.Name{Space: "", Local: "poll-status"},
		Code:    1,
		Name:    "Up",
		Time:    "1969-12-31T19:00:00-05:00",
	}

	expectedResponse := PollerResponseDTO{
		XMLName: xml.Name{Space: "", Local: "poller-response"},
		Status:  status,
	}

	MarshalUnmarshal(t, expectedResponse, expectedXml, &(PollerResponseDTO{}))
}

func TestMarshalUnmarshalDnsRequestDTO(t *testing.T) {
	expectedXml := `<dns-lookup-request location="MINION" host-request="localhost" query-type="LOOKUP"></dns-lookup-request>`

	expectedRequest := DnsRequestDTO{
		XMLName:     xml.Name{Space: "", Local: "dns-lookup-request"},
		Location:    "MINION",
		HostRequest: "localhost",
		QueryType:   "LOOKUP",
	}

	MarshalUnmarshal(t, expectedRequest, expectedXml, &(DnsRequestDTO{}))
}

func TestMarshalUnmarshalDnsResponseDTO(t *testing.T) {
	expectedXml := `<dns-lookup-response host-response="127.0.0.1" failure-message="UnknownHost"></dns-lookup-response>`

	expectedResponse := DnsResponseDTO{
		XMLName:        xml.Name{Space: "", Local: "dns-lookup-response"},
		HostResponse:   "127.0.0.1",
		FailureMessage: "UnknownHost",
	}

	MarshalUnmarshal(t, expectedResponse, expectedXml, &(DnsResponseDTO{}))
}

func TestMarshalUnmarshalPingRequestDTO(t *testing.T) {
	expectedXml := `<ping-request location="MINION" retries="10" timeout="5000">
  <address>127.0.0.1</address>
  <packet-size>20</packet-size>
</ping-request>`

	expectedRequest := PingRequestDTO{
		XMLName:    xml.Name{Space: "", Local: "ping-request"},
		Location:   "MINION",
		Retries:    10,
		Timeout:    5000,
		Address:    "127.0.0.1",
		PacketSize: 20,
	}

	MarshalUnmarshal(t, expectedRequest, expectedXml, &(PingRequestDTO{}))
}

func TestMarshalUnmarshalPingResponseDTO(t *testing.T) {
	expectedXml := `<ping-response>
  <rtt>10.253</rtt>
</ping-response>`

	expectedResponse := PingResponseDTO{
		XMLName: xml.Name{Space: "", Local: "ping-response"},
		RTT:     10.253,
	}

	MarshalUnmarshal(t, expectedResponse, expectedXml, &(PingResponseDTO{}))
}

func TestMarshalUnmarshalPingSweepRequestDTO(t *testing.T) {
	expectedXml := `<ping-sweep-request location="MINION" packets-per-second="9.5" packet-size="64">
  <ip-range begin="127.0.0.1" end="127.0.0.5" retries="2" timeout="1000"></ip-range>
</ping-sweep-request>`

	pingSweepRange := PingSweepRangeDTO{
		XMLName: xml.Name{Space: "", Local: "ip-range"},
		Begin:   "127.0.0.1",
		End:     "127.0.0.5",
		Retries: 2,
		Timeout: 1000,
	}

	expectedRequest := PingSweepRequestDTO{
		XMLName:          xml.Name{Space: "", Local: "ping-sweep-request"},
		Location:         "MINION",
		PacketSize:       64,
		PacketsPerSecond: 9.5,
		Ranges:           []PingSweepRangeDTO{pingSweepRange},
	}

	MarshalUnmarshal(t, expectedRequest, expectedXml, &(PingSweepRequestDTO{}))
}

func TestMarshalUnmarshalPingSweepResponseDTO(t *testing.T) {
	expectedXml := `<ping-sweep-response>
  <pinger-result>
    <address>127.0.0.1</address>
    <rtt>0.243</rtt>
  </pinger-result>
</ping-sweep-response>`

	pingSweepResult := PingSweepResultDTO{
		XMLName: xml.Name{Space: "", Local: "pinger-result"},
		Address: "127.0.0.1",
		RTT:     0.243,
	}

	expectedResponse := PingSweepResponseDTO{
		XMLName: xml.Name{Space: "", Local: "ping-sweep-response"},
		Results: []PingSweepResultDTO{pingSweepResult},
	}

	MarshalUnmarshal(t, expectedResponse, expectedXml, &(PingSweepResponseDTO{}))
}
