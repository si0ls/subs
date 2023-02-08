package stlxml

import (
	"encoding/xml"
	"time"

	"github.com/si0ls/subs/stl"
)

func unmarshalXMLInt(i *int, d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	var err error
	*i, err = decodeInt(s)
	return err
}

func marshalXMLInt(i int, e *xml.Encoder, start xml.StartElement, length int) error {
	return e.EncodeElement(encodeInt(i, length), start)
}

func unmarshalXMLByte(b *byte, d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	var err error
	*b, err = decodeByte(s)
	return err
}

func marshalXMLByte(b byte, e *xml.Encoder, start xml.StartElement, length int) error {
	return e.EncodeElement(encodeByte(b, length), start)
}

func unmarshalXMLHexInt(i *int, d *xml.Decoder, start xml.StartElement, upper bool) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	var err error
	*i, err = decodeHexInt(s, upper)
	return err
}

func marshalXMLHexInt(i int, e *xml.Encoder, start xml.StartElement, length int, upper bool) error {
	return e.EncodeElement(encodeHexInt(i, length, upper), start)
}

func unmarshalXMLHexByte(i *byte, d *xml.Decoder, start xml.StartElement, upper bool) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	var err error
	*i, err = decodeHexByte(s, upper)
	return err
}

func marshalXMLHexByte(b byte, e *xml.Encoder, start xml.StartElement, length int, upper bool) error {
	return e.EncodeElement(encodeHexByte(b, length, upper), start)
}

func unmarshalXMLTimecode(t *stl.Timecode, d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	var err error
	*t, err = decodeTimecode(s)
	return err
}

func marshalXMLTimecode(t stl.Timecode, e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(encodeTimecode(t), start)
}

func unmarshalXMLTime(t *time.Time, d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	var err error
	*t, err = decodeTime(s)
	return err
}

func marshalXMLTime(t time.Time, e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(encodeTime(t), start)
}

func unmarshalXMLString(str *string, d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	*str = decodeString(s)
	return nil
}

func marshalXMLString(s string, e *xml.Encoder, start xml.StartElement, length int) error {
	return e.EncodeElement(encodeString(s, length), start)
}

func unmarshalXMLBytes(b *[]byte, d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	*b = []byte(s)
	return nil
}

func marshalXMLBytes(b []byte, e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(string(b), start)
}
