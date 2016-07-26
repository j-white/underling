package underlinglib

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/soniah/gosnmp"
	"net"
	"strconv"
	"strings"
)

func marshalBase128Int(out *bytes.Buffer, n int64) (err error) {
	if n == 0 {
		err = out.WriteByte(0)
		return
	}

	l := 0
	for i := n; i > 0; i >>= 7 {
		l++
	}

	for i := l - 1; i >= 0; i-- {
		o := byte(n >> uint(i*7))
		o &= 0x7f
		if i != 0 {
			o |= 0x80
		}
		err = out.WriteByte(o)
		if err != nil {
			return
		}
	}

	return nil
}

func marshalObjectIdentifier(oid []int) (ret []byte, err error) {
	out := new(bytes.Buffer)
	if len(oid) < 2 || oid[0] > 6 || oid[1] >= 40 {
		return nil, errors.New("invalid object identifier")
	}

	err = out.WriteByte(byte(oid[0]*40 + oid[1]))
	if err != nil {
		return
	}
	for i := 2; i < len(oid); i++ {
		err = marshalBase128Int(out, int64(oid[i]))
		if err != nil {
			return
		}
	}

	ret = out.Bytes()
	return
}

func marshalOID(oid string) ([]byte, error) {
	var err error

	// Encode the oid
	oid = strings.Trim(oid, ".")
	oidParts := strings.Split(oid, ".")
	oidBytes := make([]int, len(oidParts))

	// Convert the string OID to an array of integers
	for i := 0; i < len(oidParts); i++ {
		oidBytes[i], err = strconv.Atoi(oidParts[i])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse OID: %s\n", err.Error())
		}
	}

	mOid, err := marshalObjectIdentifier(oidBytes)

	if err != nil {
		return nil, fmt.Errorf("Unable to marshal OID: %s\n", err.Error())
	}

	return mOid, err
}

/**
Converts the value to a byte-array that can be used to initialize a
java.math.BigInteger via the (byte[]) constructor.
*/
func toJavaBigIntegerBytes(value uint32) []byte {
	// Convert the integer to a byte-array
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, value)
	return bytesToJavaBigIntegerBytes(bytes)
}

func bytesToJavaBigIntegerBytes(valueBytes []byte) []byte {
	var bytes []byte
	// Find the first byte with a non-zero value, and trim the slice
	offset := 0
	for ; offset < len(valueBytes)-1; offset++ {
		if valueBytes[offset] != 0 {
			break
		}
	}

	if len(valueBytes) < 1 {
		bytes = []byte{byte(0)}
	} else {
		bytes = valueBytes[offset:]
	}

	// If the left-most bit of the first byte is 1, prepend another byte for the sign
	if bytes[0]>>7 == 1 {
		bytes = append([]byte{byte(0)}, bytes...)
	}

	return bytes
}

func getResultForPDU(pdu gosnmp.SnmpPDU, base string) SNMPResultDTO {
	var valueBytes []byte
	switch pdu.Type {
	case gosnmp.Counter32:
		fallthrough
	case gosnmp.Counter64:
		fallthrough
	case gosnmp.Uinteger32:
		fallthrough
	case gosnmp.Integer:
		fallthrough
	case gosnmp.TimeTicks:
		valueBytes = bytesToJavaBigIntegerBytes(gosnmp.ToBigInt(pdu.Value).Bytes())
	case gosnmp.OctetString:
		valueBytes = pdu.Value.([]byte)
	case gosnmp.ObjectIdentifier:
		valueBytes = []byte(pdu.Value.(string))
	case gosnmp.IPAddress:
		ip := net.ParseIP(pdu.Value.(string))
		ip4 := ip.To4()
		if ip4 != nil {
			valueBytes = ip4
		} else {
			valueBytes = ip.To16()
		}
	default:
		valueBytes = make([]byte, 0)
	}

	result := SNMPResultDTO{
		Base:     base,
		Instance: pdu.Name[len(base):],
		Value: SNMPValueDTO{
			Type:  int(pdu.Type),
			Value: base64.StdEncoding.EncodeToString(valueBytes),
		},
	}
	return result
}

func getOidToWalk(base string, instance string) string {
	var effectiveOid string
	if len(instance) > 0 {
		// Append the instance to the OID
		effectiveOid = base + instance
		// And remove the last byte
		effectiveOid = effectiveOid[:len(effectiveOid)-2]
	} else {
		// Use the OID "as-is"
		effectiveOid = base
	}
	return effectiveOid
}
