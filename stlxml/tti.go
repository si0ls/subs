package stlxml

import (
	"encoding/xml"

	"github.com/si0ls/subs/stl"
)

// TTIXML is the XML representation of a STL TTI block.
type TTIXML struct {
	XMLName xml.Name       `xml:"TTI"`
	SGN     SGNXML         `xml:"SGN"` // Subtitle Group Number
	SN      SNXML          `xml:"SN"`  // Subtitle Number
	EBN     EBNXML         `xml:"EBN"` // Extension Block Number
	CS      CSXML          `xml:"CS"`  // Cumulative Status
	TCI     TCIXML         `xml:"TCI"` // Time Code: In
	TCO     TCOXML         `xml:"TCO"` // Time Code: Out
	VP      VPXML          `xml:"VP"`  // Vertical Position
	JC      JCXML          `xml:"JC"`  // Justification Code
	CF      CFXML          `xml:"CF"`  // Comment Flag
	TF      TFXMLContainer `xml:"TF"`  // Text Field
}

type TFXMLContainer struct {
	InnerXML string `xml:",innerxml"`
}

// FromSTL converts a stl.TTIBlock to a stlxml.TTIXML.
func (ttiXML *TTIXML) FromSTL(TTIstl stl.TTIBlock, cct stl.CharacterCodeTable) {
	ttiXML.SGN = SGNXML(TTIstl.SGN)
	ttiXML.SN = SNXML(TTIstl.SN)
	ttiXML.EBN = EBNXML(TTIstl.EBN)
	ttiXML.CS = CSXML(TTIstl.CS)
	ttiXML.TCI = TCIXML(TTIstl.TCI)
	ttiXML.TCO = TCOXML(TTIstl.TCO)
	ttiXML.VP = VPXML(TTIstl.VP)
	ttiXML.JC = JCXML(TTIstl.JC)
	ttiXML.CF = CFXML(TTIstl.CF)

	// TODO: remove this hack and implement a proper XML encoder
	s, _ := encodeTextField(TTIstl.TF, cct)
	ttiXML.TF.InnerXML = s
}

// ToSTL converts a stlxml.TTIXML to a stl.TTIBlock
func (ttiXML TTIXML) ToSTL(cct stl.CharacterCodeTable) stl.TTIBlock {
	// TODO: remove this hack and implement a proper XML decoder
	s, _ := decodeTextField(ttiXML.TF.InnerXML, cct)

	return stl.TTIBlock{
		SGN: int(ttiXML.SGN),
		SN:  int(ttiXML.SN),
		EBN: int(ttiXML.EBN),
		CS:  stl.CumulativeStatus(ttiXML.CS),
		TCI: stl.Timecode(ttiXML.TCI),
		TCO: stl.Timecode(ttiXML.TCO),
		VP:  int(ttiXML.VP),
		JC:  stl.JustificationCode(ttiXML.JC),
		CF:  stl.CommentFlag(ttiXML.CF),
		TF:  s,
	}
}

// Validate validates the TTI block.
func (tti *TTIXML) Validate() []error {
	panic("not implemented")
}

// SGNXML is the XML representation of STL Subtitle Group Number (SGN).
type SGNXML int

// SGNXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*SGNXML)(nil)
var _ xml.Marshaler = (*SGNXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the SGNXML value pointed to by sgn.
func (sgn *SGNXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(sgn), d, start)
}

// MarshalXML returns the XML encoding of sgn.
func (sgn SGNXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(sgn), e, start, 1)
}

// SNXML is the XML representation of STL Subtitle Number (SN).
type SNXML int

// SNXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*SNXML)(nil)
var _ xml.Marshaler = (*SNXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the SNXML value pointed to by sn.
func (sn *SNXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(sn), d, start)
}

// MarshalXML returns the XML encoding of sn.
func (sn SNXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(sn), e, start, 1)
}

// EBNXML is the XML representation of STL Extension Block Number (EBN).
type EBNXML int

// EBNXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*EBNXML)(nil)
var _ xml.Marshaler = (*EBNXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the EBNXML value pointed to by ebn.
func (ebn *EBNXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLHexInt((*int)(ebn), d, start, false)
}

// MarshalXML returns the XML encoding of ebn.
func (ebn EBNXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLHexInt(int(ebn), e, start, 2, false)
}

// CSXML is the XML representation of STL Cumulative Status (CS).
type CSXML stl.CumulativeStatus

// CSXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*CSXML)(nil)
var _ xml.Marshaler = (*CSXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the CSXML value pointed to by cs.
func (cs *CSXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLHexByte((*byte)(cs), d, start, false)
}

// MarshalXML returns the XML encoding of cs.
func (cs CSXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLHexByte(byte(cs), e, start, 2, false)
}

// TCIXML is the XML representation of STL Time Code In (TCI).
type TCIXML stl.Timecode

// TCIXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TCIXML)(nil)
var _ xml.Marshaler = (*TCIXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TCIXML value pointed to by tci.
func (tci *TCIXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLTimecode((*stl.Timecode)(tci), d, start)
}

// MarshalXML returns the XML encoding of tci.
func (tci TCIXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLTimecode(stl.Timecode(tci), e, start)
}

// TCOXML is the XML representation of STL Time Code Out (TCO).
type TCOXML stl.Timecode

// TCOXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TCOXML)(nil)
var _ xml.Marshaler = (*TCOXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TCOXML value pointed to by tco.
func (tco *TCOXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLTimecode((*stl.Timecode)(tco), d, start)
}

// MarshalXML returns the XML encoding of tco.
func (tco TCOXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLTimecode(stl.Timecode(tco), e, start)
}

// VPXML is the XML representation of STL Vertical Position (VP).
type VPXML int

// VPXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*VPXML)(nil)
var _ xml.Marshaler = (*VPXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the VPXML value pointed to by vp.
func (vp *VPXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(vp), d, start)
}

// MarshalXML returns the XML encoding of vp.
func (vp VPXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(vp), e, start, 1)
}

// JCXML is the XML representation of STL Justification Code (JC).
type JCXML stl.JustificationCode

// JCXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*JCXML)(nil)
var _ xml.Marshaler = (*JCXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the JCXML value pointed to by jc.
func (jc *JCXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLHexByte((*byte)(jc), d, start, false)
}

// MarshalXML returns the XML encoding of jc.
func (jc JCXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLHexByte(byte(jc), e, start, 2, false)
}

// CCXML is the XML representation of STL Color Code (CC).
type CFXML stl.CommentFlag

// CFXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*CFXML)(nil)
var _ xml.Marshaler = (*CFXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the CFXML value pointed to by cf.
func (cf *CFXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLHexByte((*byte)(cf), d, start, false)
}

// MarshalXML returns the XML encoding of cf.
func (cf CFXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLHexByte(byte(cf), e, start, 2, false)
}

/*
// CCXML is the XML representation of STL Color Code (CC).
type TFXML string

// TFXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TFXML)(nil)
var _ xml.Marshaler = (*TFXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TFXML value pointed to by tf.
func (tf *TFXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(tf), d, start)
}

// MarshalXML returns the XML encoding of tf.
func (tf TFXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//return marshalXMLString(string(tf), e, start, 0)
	s, err := decodeTextField(string(tf))
	if err != nil {
		return err
	}
	return marshalXMLString(s, e, start, 0)
}
*/
/*
// UnmarshalXML decodes the XML-encoded data and stores the result in the TFXML value pointed to by tf.
func (tf *TFXMLContainer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString(&tf.Text, d, start)
}

// MarshalXML returns the XML encoding of tf.
func (tf TFXMLContainer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//return marshalXMLString(string(tf), e, start, 0)
	//s, err := decodeTextField(tf.Text)
	//if err != nil {
	//	return err
	//}
	return marshalXMLString("<test>", e, start, 0)
}
*/
