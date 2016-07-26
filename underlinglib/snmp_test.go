package underlinglib

import (
//	"github.com/stretchr/testify/assert"
//	"testing"
)

/*
func TestSNMPExec(t *testing.T) {

	request := SNMPRequestDTO{
		Agent: SNMPAgentDTO{
			Address:       "127.0.0.1",
			ReadCommunity: "public",
			Version:       2,
		},
		Walks: []SNMPWalkRequestDTO{SNMPWalkRequestDTO{
			CorrelationID: "42",
			OIDs:          []string{".1.3.6.1.4.1.8072.1.7.2.1.1.3.6.115.110.109.112"},
		}},
	}

	expectedResponse := SNMPMultiResponseDTO{
		Responses: []SNMPResponseDTO{SNMPResponseDTO{
			CorrelationID: "42",
			Results: []SNMPResultDTO{SNMPResultDTO{
				Base:     ".1.3.6.1.4.1.8072.1.7.2.1.1.3.6.115.110.109.112",
				Instance: ".100",
				Value: SNMPValueDTO{
					Type:  2,
					Value: "BA==",
				},
			}},
		}},
	}

	response := Exec(request)

	assert.Equal(t, expectedResponse, response)

}
*/
