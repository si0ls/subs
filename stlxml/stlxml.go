package stlxml

import (
	"encoding/xml"
	"io"

	"github.com/si0ls/subs/stl"
)

// STLXML is the XML representation of STL file.
type STLXML struct {
	XMLName xml.Name `xml:"StlXml"`
	GSI     GSIXML   `xml:"HEAD>GSI"`
	TTI     []TTIXML `xml:"BODY>TTICONTAINER>TTI"`
}

// New returns a new stlxml.STLXML.
func New() *STLXML {
	return &STLXML{}
}

// Decode reads and decodes the STLXML file from r.
func (stlXML *STLXML) Decode(r io.Reader) error {
	dec := xml.NewDecoder(r)
	return dec.Decode(stlXML)
}

// Encode encodes and writes the STLXML file to w.
func (stlXML *STLXML) Encode(w io.Writer) error {
	return stlXML.EncodeIndent(w, "", "  ")
}

// ENcodeIndent encodes and writes the STLXML file to w with indent.
func (stlXML *STLXML) EncodeIndent(w io.Writer, prefix, indent string) error {
	enc := xml.NewEncoder(w)
	enc.Indent(prefix, indent)
	return enc.Encode(stlXML)
}

// FromSTL converts a stl.File to a stlxml.STLXML.
func (stlXML *STLXML) FromSTL(stlFile stl.File) {
	stlXML.GSI.FromSTL(*stlFile.GSI)
	stlXML.TTI = make([]TTIXML, len(stlFile.TTI))
	for i, tti := range stlFile.TTI {
		stlXML.TTI[i].FromSTL(*tti, stlFile.GSI.CCT)
	}
}

// ToSTL converts a stlxml.STLXML to a stl.File.
func (stlXML *STLXML) ToSTL(cct stl.CharacterCodeTable) stl.File {
	var file stl.File
	gsi := stlXML.GSI.ToSTL()
	file.GSI = &gsi
	file.TTI = make([]*stl.TTIBlock, len(stlXML.TTI))
	for i, ttiXML := range stlXML.TTI {
		tti := ttiXML.ToSTL(cct)
		file.TTI[i] = &tti
	}
	return file
}

// Validate validates the STLXML file.
func (stlXML *STLXML) Validate() []error {
	panic("not implemented")
}
