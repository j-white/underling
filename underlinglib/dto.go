package underlinglib

import (
	"encoding/xml"
)

type SNMPAgentDTO struct {
	XMLName       xml.Name `xml:"agent"`
	Address       string   `xml:"address"`
	ReadCommunity string   `xml:"readCommunity"`
	Version       int      `xml:"version"`
}

type SNMPGetRequestDTO struct {
	XMLName       xml.Name `xml:"get"`
	CorrelationID string   `xml:"correlation-id,attr"`
	OIDs          []string `xml:"oid"`
}

type SNMPWalkRequestDTO struct {
	XMLName        xml.Name `xml:"walk"`
	CorrelationID  string   `xml:"correlation-id,attr"`
	MaxRepetitions int      `xml:"max-repetitions,attr,omitempty"`
	Instance       string   `xml:"instance,attr,omitempty"`
	OIDs           []string `xml:"oid"`
}

type SNMPRequestDTO struct {
	XMLName     xml.Name             `xml:"snmp-request"`
	Location    string               `xml:"location,attr"`
	Description string               `xml:"description,attr"`
	Agent       SNMPAgentDTO         `xml:"agent"`
	Gets        []SNMPGetRequestDTO  `xml:"get"`
	Walks       []SNMPWalkRequestDTO `xml:"walk"`
}

type SNMPValueDTO struct {
	XMLName xml.Name `xml:"value"`
	Type    int      `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type SNMPResultDTO struct {
	XMLName  xml.Name     `xml:"result"`
	Base     string       `xml:"base"`
	Instance string       `xml:"instance"`
	Value    SNMPValueDTO `xml:"value"`
}

type SNMPResponseDTO struct {
	XMLName       xml.Name        `xml:"response"`
	CorrelationID string          `xml:"correlation-id,attr"`
	Results       []SNMPResultDTO `xml:"result"`
}

type SNMPMultiResponseDTO struct {
	XMLName   xml.Name          `xml:"snmp-response"`
	Responses []SNMPResponseDTO `xml:"response"`
}

type DetectorAttributeDTO struct {
	XMLName xml.Name `xml:"detector-attribute"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:",chardata"`
}

type RuntimeAttributeDTO struct {
	XMLName xml.Name `xml:"runtime-attribute"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:",chardata"`
}

type AttributeDTO struct {
	XMLName xml.Name `xml:"attribute"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:",chardata"`
}

type DetectorRequestDTO struct {
	XMLName            xml.Name               `xml:"detector-request"`
	Location           string                 `xml:"location,attr"`
	ClassName          string                 `xml:"class-name,attr"`
	Address            string                 `xml:"address,attr"`
	DetectorAttributes []DetectorAttributeDTO `xml:"detector-attribute"`
	RuntimeAttributes  []RuntimeAttributeDTO  `xml:"runtime-attribute"`
}

type DetectorResponseDTO struct {
	XMLName        xml.Name       `xml:"detector-response"`
	Detected       bool           `xml:"detected,attr"`
	FailureMessage string         `xml:"failure-message,attr,omitempty"`
	Attributes     []AttributeDTO `xml:"attribute"`
}

type MinionIdentityDTO struct {
	XMLName  xml.Name `xml:"minion"`
	Id       string   `xml:"id"`
	Location string   `xml:"location"`
}

type PollerAttributeDTO struct {
	XMLName  xml.Name `xml:"attribute"`
	Key      string   `xml:"key,attr"`
	Value    string   `xml:"value,attr"`
	Contents string   `xml:",chardata"`
}

type PollerRequestDTO struct {
	XMLName    xml.Name             `xml:"poller-request"`
	Location   string               `xml:"location,attr"`
	ClassName  string               `xml:"class-name,attr"`
	Address    string               `xml:"address,attr"`
	Attributes []PollerAttributeDTO `xml:"attribute"`
}

type PollerStatusDTO struct {
	XMLName      xml.Name `xml:"poll-status"`
	Code         int      `xml:"code,attr"`
	Name         string   `xml:"name,attr,omitempty"`
	ResponseTime float64  `xml:"response-time,attr,omitempty"`
	Reason       string   `xml:"reason,attr,omitempty"`
	Time         string   `xml:"time,attr"`
}

type PollerResponseDTO struct {
	XMLName xml.Name        `xml:"poller-response"`
	Status  PollerStatusDTO `xml:"poll-status"`
}
