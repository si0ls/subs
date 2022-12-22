package stl

import (
	"encoding/xml"
	"io"
)

func (gsi *GSIBlock) EncodeXML(w io.Writer) error {
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	return enc.Encode(gsi)
}

func (gsi *GSIBlock) DecodeXML(r io.Reader) error {
	dec := xml.NewDecoder(r)
	return dec.Decode(gsi)
}
