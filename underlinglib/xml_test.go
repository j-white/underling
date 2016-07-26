package underlinglib

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func MarshalUnmarshal(t *testing.T, val interface{}, xmls string, ref interface{}) {
	expectedXml, expectedValue := xmls, val

	// Unmarshal
	err := UnmarshalFromXml(strings.NewReader(expectedXml), ref)
	assert.Nil(t, err)
	assert.Equal(t, expectedValue, reflect.ValueOf(ref).Elem().Interface())

	// Marshal
	actualXml, err := MarshalToXml(ref)
	assert.Nil(t, err)
	assert.Equal(t, expectedXml, actualXml)
}
