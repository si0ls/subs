package stl

import (
	"encoding/xml"
	"io"
)

func (tti *TTIBlock) EncodeXML(w io.Writer) error {
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	return enc.Encode(tti)
}

func (tti *TTIBlock) DecodeXML(r io.Reader) error {
	dec := xml.NewDecoder(r)
	return dec.Decode(tti)
}

/*
func (tti *TTIBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "TTIBlock"
	if err := e.EncodeElement(tti.SGN, xml.StartElement{Name: xml.Name{Local: "SGN"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(tti.SN, xml.StartElement{Name: xml.Name{Local: "SN"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(tti.EBN, xml.StartElement{Name: xml.Name{Local: "EBN"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(tti.CS, xml.StartElement{Name: xml.Name{Local: "CS"}}); err != nil {
		return err
	}
	return nil
}
*/
