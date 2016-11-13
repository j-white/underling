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

type DnsRequestDTO struct {
	XMLName     xml.Name `xml:"dns-lookup-request"`
	Location    string   `xml:"location,attr"`
	HostRequest string   `xml:"host-request,attr"`
	QueryType   string   `xml:"query-type,attr"`
}

type DnsResponseDTO struct {
	XMLName        xml.Name `xml:"dns-lookup-response"`
	HostResponse   string   `xml:"host-response,attr"`
	FailureMessage string   `xml:"failure-message,attr,omitempty"`
}

type PingRequestDTO struct {
	XMLName    xml.Name `xml:"ping-request"`
	Location   string   `xml:"location,attr"`
	Retries    int      `xml:"retries,attr"`
	Timeout    int      `xml:"timeout,attr"`
	Address    string   `xml:"address"`
	PacketSize int      `xml:"packet-size"`
}

type PingResponseDTO struct {
	XMLName xml.Name `xml:"ping-response"`
	RTT     float64  `xml:"rtt"`
}

type PingSweepRangeDTO struct {
	XMLName xml.Name `xml:"ip-range"`
	Begin   string   `xml:"begin,attr"`
	End     string   `xml:"end,attr"`
	Retries int      `xml:"retries,attr"`
	Timeout int      `xml:"timeout,attr"`
}

type PingSweepRequestDTO struct {
	XMLName          xml.Name            `xml:"ping-sweep-request"`
	Location         string              `xml:"location,attr"`
	PacketsPerSecond float64             `xml:"packets-per-second,attr"`
	PacketSize       int                 `xml:"packet-size,attr"`
	Ranges           []PingSweepRangeDTO `xml:"ip-range"`
}

type PingSweepResultDTO struct {
	XMLName xml.Name `xml:"pinger-result"`
	Address string   `xml:"address"`
	RTT     float64  `xml:"rtt"`
}

type PingSweepResponseDTO struct {
	XMLName xml.Name             `xml:"ping-sweep-response"`
	Results []PingSweepResultDTO `xml:"pinger-result"`
}
