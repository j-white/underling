package underlinglib

import (
	"bytes"
	"encoding/xml"
	"io"
)

func UnmarshalFromXml(reader io.Reader, v interface{}) error {
	if err := xml.NewDecoder(reader).Decode(v); err != nil {
		return err
	}
	return nil
}

func MarshalToXml(v interface{}) (string, error) {
	buf := new(bytes.Buffer)
	encoder := xml.NewEncoder(buf)
	encoder.Indent("", "  ")
	if err := encoder.Encode(v); err != nil {
		return "", err
	}

	return buf.String(), nil
}
