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
