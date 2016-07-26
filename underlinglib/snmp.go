package underlinglib

import (
	"fmt"
	"github.com/soniah/gosnmp"
	"time"
)

func dowalk(request SNMPRequestDTO, walk SNMPWalkRequestDTO) SNMPResponseDTO {
	response := SNMPResponseDTO{CorrelationID: walk.CorrelationID}

	// TODO: Pull all of the fields from the agent
	client := gosnmp.GoSNMP{
		Target:    request.Agent.Address,
		Port:      161,
		Community: request.Agent.ReadCommunity,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Retries:   3,
	}

	if err := client.Connect(); err != nil {
		fmt.Printf("Connect Error: %v\n", err)
		return response
	}
	defer client.Conn.Close()

	for _, oid := range walk.OIDs {
		effectiveOid := getOidToWalk(oid, walk.Instance)
		err := client.BulkWalk(effectiveOid, func(pdu gosnmp.SnmpPDU) error {
			response.Results = append(response.Results, getResultForPDU(pdu, oid))
			return nil
		})
		if err != nil {
			fmt.Printf("Walk Error: %v\n", err)
			return response
		}
	}

	return response
}

func Exec(request SNMPRequestDTO) SNMPMultiResponseDTO {
	multiResponse := SNMPMultiResponseDTO{}

	for _, walk := range request.Walks {
		multiResponse.Responses = append(multiResponse.Responses, dowalk(request, walk))
	}

	// TODO: Handle GETs too
	return multiResponse
}
