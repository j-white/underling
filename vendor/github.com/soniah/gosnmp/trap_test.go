// Copyright 2012-2016 The GoSNMP Authors. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

package gosnmp

import (
	"log"
	"os" //"io/ioutil"
	//"testing"
)

var testsUnmarshalTrap = []struct {
	in  func() []byte
	out *SnmpPacket
}{
	{genericV3Trap,
		&SnmpPacket{
			Version:   Version3,
			PDUType:   SNMPv2Trap,
			RequestID: 190378322,
			MsgFlags:  AuthNoPriv,
			SecurityParameters: &UsmSecurityParameters{
				UserName:                 "myuser",
				AuthenticationProtocol:   MD5,
				AuthenticationPassphrase: "mypassword",
				Logger: log.New(os.Stdout, "", 0),
			},
		},
	},
}

/*func TestUnmarshalTrap(t *testing.T) {
	Default.Logger = log.New(os.Stdout, "", 0)

SANITY:
	for i, test := range testsUnmarshalTrap {

		Default.SecurityParameters = test.out.SecurityParameters.Copy()

		var buf = test.in()
		var res = Default.unmarshalTrap(buf)
		if res == nil {
			t.Errorf("#%d, UnmarshalTrap returned nil", i)
			continue SANITY
		}

		// test enough fields fields to ensure unmarshalling was sucessful.
		// full unmarshal testing is performed in TestUnmarshal
		if res.Version != test.out.Version {
			t.Errorf("#%d Version result: %v, test: %v", i, res.Version, test.out.Version)
		}
		if res.RequestID != test.out.RequestID {
			t.Errorf("#%d RequestID result: %v, test: %v", i, res.RequestID, test.out.RequestID)
		}
	}
}
*/
func genericV3Trap() []byte {
	return []byte{
		0x30, 0x81, 0xd7, 0x02, 0x01, 0x03, 0x30, 0x11, 0x02, 0x04, 0x62, 0xaf,
		0x5a, 0x8e, 0x02, 0x03, 0x00, 0xff, 0xe3, 0x04, 0x01, 0x01, 0x02, 0x01,
		0x03, 0x04, 0x33, 0x30, 0x31, 0x04, 0x11, 0x80, 0x00, 0x1f, 0x88, 0x80,
		0x77, 0xdf, 0xe4, 0x4f, 0xaa, 0x70, 0x02, 0x58, 0x00, 0x00, 0x00, 0x00,
		0x02, 0x01, 0x0f, 0x02, 0x01, 0x00, 0x04, 0x06, 0x6d, 0x79, 0x75, 0x73,
		0x65, 0x72, 0x04, 0x0c, 0xd8, 0xb6, 0x9c, 0xb8, 0x22, 0x91, 0xfc, 0x65,
		0xb6, 0x84, 0xcb, 0xfe, 0x04, 0x00, 0x30, 0x81, 0x89, 0x04, 0x11, 0x80,
		0x00, 0x1f, 0x88, 0x80, 0x77, 0xdf, 0xe4, 0x4f, 0xaa, 0x70, 0x02, 0x58,
		0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0xa7, 0x72, 0x02, 0x04, 0x39, 0x19,
		0x9c, 0x61, 0x02, 0x01, 0x00, 0x02, 0x01, 0x00, 0x30, 0x64, 0x30, 0x0f,
		0x06, 0x08, 0x2b, 0x06, 0x01, 0x02, 0x01, 0x01, 0x03, 0x00, 0x43, 0x03,
		0x15, 0x2f, 0xec, 0x30, 0x14, 0x06, 0x0a, 0x2b, 0x06, 0x01, 0x06, 0x03,
		0x01, 0x01, 0x04, 0x01, 0x00, 0x06, 0x06, 0x2b, 0x06, 0x01, 0x02, 0x01,
		0x01, 0x30, 0x16, 0x06, 0x08, 0x2b, 0x06, 0x01, 0x02, 0x01, 0x01, 0x01,
		0x00, 0x04, 0x0a, 0x72, 0x65, 0x64, 0x20, 0x6c, 0x61, 0x70, 0x74, 0x6f,
		0x70, 0x30, 0x0d, 0x06, 0x08, 0x2b, 0x06, 0x01, 0x02, 0x01, 0x01, 0x07,
		0x00, 0x02, 0x01, 0x05, 0x30, 0x14, 0x06, 0x07, 0x2b, 0x06, 0x01, 0x02,
		0x01, 0x01, 0x02, 0x06, 0x09, 0x2b, 0x06, 0x01, 0x04, 0x01, 0x02, 0x03,
		0x04, 0x05}
}
