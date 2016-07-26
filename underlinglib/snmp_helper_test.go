package underlinglib

import (
	"github.com/soniah/gosnmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToJavaBigIntegerBytes(t *testing.T) {
	assert.Equal(t, []byte{0}, toJavaBigIntegerBytes(0))
	assert.Equal(t, []byte{1}, toJavaBigIntegerBytes(1))
	assert.Equal(t, []byte{127}, toJavaBigIntegerBytes(127))
	assert.Equal(t, []byte{0, 128}, toJavaBigIntegerBytes(128))
	assert.Equal(t, []byte{127, 255}, toJavaBigIntegerBytes(32767))
	assert.Equal(t, []byte{0, 255, 255, 255, 255}, toJavaBigIntegerBytes(4294967295))
}

func TestPDUToResultBaseAndInstanceHandling(t *testing.T) {
	verifyBaseAndInstance(t, ".1.3.6.1.2.1.1.3.0", ".1.3.6.1.2.1.1.3", ".0")
	verifyBaseAndInstance(t, ".1.3.6.1.2.1.1.3.0", ".1.3.6.1.2.1.1", ".3.0")
	verifyBaseAndInstance(t, ".1.3.6.1.2.1.1.3.0", ".1.3.6.1.2.1", ".1.3.0")
	// TODO: Verify error conditions i.e. OID does not start with base
}

func verifyBaseAndInstance(t *testing.T, oid string, base string, resultInstance string) {
	pdu := gosnmp.SnmpPDU{
		Name:  oid,
		Type:  gosnmp.TimeTicks,
		Value: 0,
	}
	expectedResult := SNMPResultDTO{
		Base:     base,
		Instance: resultInstance,
		Value: SNMPValueDTO{
			Type:  67,
			Value: "AA==",
		},
	}
	assert.Equal(t, expectedResult, getResultForPDU(pdu, base))
}

func TestPDUToResultTypeAndValueHandling(t *testing.T) {
	verifyTypeAndResult(t, gosnmp.ObjectIdentifier, ".1.3.6.1.2.1.92", 6, "LjEuMy42LjEuMi4xLjky")
	verifyTypeAndResult(t, gosnmp.TimeTicks, 0, 67, "AA==")
	verifyTypeAndResult(t, gosnmp.TimeTicks, 1073741824, 67, "QAAAAA==")
	verifyTypeAndResult(t, gosnmp.TimeTicks, 2147483648, 67, "AIAAAAA=")
	verifyTypeAndResult(t, gosnmp.Counter64, 30480, 70, "dxA=")
	verifyTypeAndResult(t, gosnmp.IPAddress, "172.23.1.6", 64, "rBcBBg==")
	verifyTypeAndResult(t, gosnmp.OctetString, []byte("test"), 4, "dGVzdA==")
}

func verifyTypeAndResult(t *testing.T, pduType gosnmp.Asn1BER, pduValue interface{}, resultType int, resultValue string) {
	pdu := gosnmp.SnmpPDU{
		Name:  ".1.3.6.1.2.1.1.3.0",
		Type:  pduType,
		Value: pduValue,
	}
	expectedResult := SNMPResultDTO{
		Base:     ".1.3.6.1.2.1.1.3",
		Instance: ".0",
		Value: SNMPValueDTO{
			Type:  resultType,
			Value: resultValue,
		},
	}
	assert.Equal(t, expectedResult, getResultForPDU(pdu, ".1.3.6.1.2.1.1.3"))
}

func TestGetOIDToWalk(t *testing.T) {
	assert.Equal(t, ".1.3.6.1.2.1.1.3", getOidToWalk(".1.3.6.1.2.1.1.3", ""))
	assert.Equal(t, ".1.3.6.1.2.1.1.3", getOidToWalk(".1.3.6.1.2.1.1", ".3.0"))
	assert.Equal(t, ".1.3.6.1.2.1.1.3", getOidToWalk(".1.3.6.1.2.1", ".1.3.0"))
}
