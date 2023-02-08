package stlxml

import (
	"encoding/xml"
	"time"

	"github.com/si0ls/subs/stl"
)

// GSIXML is the XML representation of STL GSI block.
type GSIXML struct {
	XMLName xml.Name `xml:"GSI"`
	CPN     CPNXML   `xml:"CPN"` // Code Page Number
	DFC     DFCXML   `xml:"DFC"` // Disk Format Code
	DSC     DSCXML   `xml:"DSC"` // Display Standard Code
	CCT     CCTXML   `xml:"CCT"` // Character Code Table number
	LC      LCXML    `xml:"LC"`  // Language Code
	OPT     OPTXML   `xml:"OPT"` // Original Program Title
	OET     OETXML   `xml:"OET"` // Original Episode Title
	TPT     TPTXML   `xml:"TPT"` // Translated Program Title
	TET     TETXML   `xml:"TET"` // Translated Episode Title
	TN      TNXML    `xml:"TN"`  // Translator's Name
	TCD     TCDXML   `xml:"TCD"` // Translator's Contact Details
	SLR     SLRXML   `xml:"SLR"` // Subtitle List Reference Code
	CD      CDXML    `xml:"CD"`  // Creation Date
	RD      RDXML    `xml:"RD"`  // Revision Date
	RN      RNXML    `xml:"RN"`  // Revision Number
	TNB     TNBXML   `xml:"TNB"` // Total Number of Text and Timing Information (TTI) blocks
	TNS     TNSXML   `xml:"TNS"` // Total Number of Subtitles
	TNG     TNGXML   `xml:"TNG"` // Total Number of Subtitle Groups
	MNC     MNCXML   `xml:"MNC"` // Maximum Number of Displayable Characters in any text row
	MNR     MNRXML   `xml:"MNR"` // Maximum Number of Displayable Rows
	TCS     TCSXML   `xml:"TCS"` // Time Code: Status
	TCP     TCPXML   `xml:"TCP"` // Time Code: Start-of-Program
	TCF     TCFXML   `xml:"TCF"` // Time Code: First In-Cue
	TND     TNDXML   `xml:"TND"` // Total Number of Disks
	DSN     DSNXML   `xml:"DSN"` // Disk Sequence Number
	CO      COXML    `xml:"CO"`  // Country of Origin
	PUB     PUBXML   `xml:"PUB"` // Publisher
	EN      ENXML    `xml:"EN"`  // Editor's Name
	ECD     ECDXML   `xml:"ECD"` // Editor's Contact
	UDA     UDAXML   `xml:"UDA"` // User-Defined Area
}

// FromSTL converts a stl.GSIBlock to a stlxml.GSIXML.
func (gsiXML *GSIXML) FromSTL(GSIstl stl.GSIBlock) {
	gsiXML.CPN = CPNXML(GSIstl.CPN)
	gsiXML.DFC = DFCXML(GSIstl.DFC)
	gsiXML.DSC = DSCXML(GSIstl.DSC)
	gsiXML.CCT = CCTXML(GSIstl.CCT)
	gsiXML.LC = LCXML(GSIstl.LC)
	gsiXML.OPT = OPTXML(GSIstl.OPT)
	gsiXML.OET = OETXML(GSIstl.OET)
	gsiXML.TPT = TPTXML(GSIstl.TPT)
	gsiXML.TET = TETXML(GSIstl.TET)
	gsiXML.TN = TNXML(GSIstl.TN)
	gsiXML.TCD = TCDXML(GSIstl.TCD)
	gsiXML.SLR = SLRXML(GSIstl.SLR)
	gsiXML.CD = CDXML(GSIstl.CD)
	gsiXML.RD = RDXML(GSIstl.RD)
	gsiXML.RN = RNXML(GSIstl.RN)
	gsiXML.TNB = TNBXML(GSIstl.TNB)
	gsiXML.TNS = TNSXML(GSIstl.TNS)
	gsiXML.TNG = TNGXML(GSIstl.TNG)
	gsiXML.MNC = MNCXML(GSIstl.MNC)
	gsiXML.MNR = MNRXML(GSIstl.MNR)
	gsiXML.TCS = TCSXML(GSIstl.TCS)
	gsiXML.TCP = TCPXML(GSIstl.TCP)
	gsiXML.TCF = TCFXML(GSIstl.TCF)
	gsiXML.TND = TNDXML(GSIstl.TND)
	gsiXML.DSN = DSNXML(GSIstl.DSN)
	gsiXML.CO = COXML(GSIstl.CO)
	gsiXML.PUB = PUBXML(GSIstl.PUB)
	gsiXML.EN = ENXML(GSIstl.EN)
	gsiXML.ECD = ECDXML(GSIstl.ECD)
	gsiXML.UDA = UDAXML(GSIstl.UDA)
}

// ToSTL converts a stlxml.GSIXML to a stl.GSIBlock.
func (gsiXML *GSIXML) ToSTL() stl.GSIBlock {
	return stl.GSIBlock{
		CPN: stl.CodePageNumber(gsiXML.CPN),
		DFC: stl.DiskFormatCode(gsiXML.DFC),
		DSC: stl.DisplayStandardCode(gsiXML.DSC),
		CCT: stl.CharacterCodeTable(gsiXML.CCT),
		LC:  stl.LanguageCode(gsiXML.LC),
		OPT: string(gsiXML.OPT),
		OET: string(gsiXML.OET),
		TPT: string(gsiXML.TPT),
		TET: string(gsiXML.TET),
		TN:  string(gsiXML.TN),
		TCD: string(gsiXML.TCD),
		SLR: string(gsiXML.SLR),
		CD:  time.Time(gsiXML.CD),
		RD:  time.Time(gsiXML.RD),
		RN:  int(gsiXML.RN),
		TNB: int(gsiXML.TNB),
		TNS: int(gsiXML.TNS),
		TNG: int(gsiXML.TNG),
		MNC: int(gsiXML.MNC),
		MNR: int(gsiXML.MNR),
		TCS: stl.TimeCodeStatus(gsiXML.TCS),
		TCP: stl.Timecode(gsiXML.TCP),
		TCF: stl.Timecode(gsiXML.TCF),
		TND: int(gsiXML.TND),
		DSN: int(gsiXML.DSN),
		CO:  string(gsiXML.CO),
		PUB: string(gsiXML.PUB),
		EN:  string(gsiXML.EN),
		ECD: string(gsiXML.ECD),
		UDA: []byte(gsiXML.UDA),
	}
}

// Validate validates the GSO block.
func (gsi *GSIXML) Validate() []error {
	panic("not implemented")
}

// CPNXML is the XML representation of STL Code Page Number (CPN).
type CPNXML stl.CodePageNumber

// CPNXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*CPNXML)(nil)
var _ xml.Marshaler = (*CPNXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the CPNXML pointed to by cpn.
func (cpn *CPNXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(cpn), d, start)
}

// MarshalXML returns the XML encoding of cpn.
func (cpn CPNXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(cpn), e, start, 3)
}

// DFCXML is the XML representation of STL Disk Format Code (DFC).
type DFCXML stl.DiskFormatCode

// DFCXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*DFCXML)(nil)
var _ xml.Marshaler = (*DFCXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the DFCXML pointed to by dfc.
func (dfc *DFCXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(dfc), d, start)
}

// MarshalXML returns the XML encoding of dfc.
func (dfc DFCXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(dfc), e, start, 8)
}

// DSCXML is the XML representation of STL Display Standard Code (DSC).
type DSCXML stl.DisplayStandardCode

// DSCXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*DSCXML)(nil)
var _ xml.Marshaler = (*DSCXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the DSCXML pointed to by dsc.
func (dsc *DSCXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLByte((*byte)(dsc), d, start)
}

// MarshalXML returns the XML encoding of dsc.
func (dsc DSCXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLByte(byte(dsc), e, start, 1)
}

// CCTXML is the XML representation of STL Character Code Table (CCT).
type CCTXML stl.CharacterCodeTable

// CCTXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*CCTXML)(nil)
var _ xml.Marshaler = (*CCTXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the CCTXML pointed to by cct.
func (cct *CCTXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLByte((*byte)(cct), d, start)
}

// MarshalXML returns the XML encoding of cct.
func (cct CCTXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLByte(byte(cct), e, start, 2)
}

// LCXML is the XML representation of STL Language Cod (LC).
type LCXML stl.LanguageCode

// LCXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*LCXML)(nil)
var _ xml.Marshaler = (*LCXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the LCXML pointed to by lc.
func (lc *LCXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLByte((*byte)(lc), d, start)
}

// MarshalXML returns the XML encoding of lc.
func (lc LCXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLHexByte(byte(lc), e, start, 2, true)
}

// OPTXML is the XML representation of STL Originating Programme Title (OPT).
type OPTXML string

// OPTXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*OPTXML)(nil)
var _ xml.Marshaler = (*OPTXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the OPTXML pointed to by opt.
func (opt *OPTXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(opt), d, start)
}

// MarshalXML returns the XML encoding of opt.
func (opt OPTXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(opt), e, start, 32)
}

// OETXML is the XML representation of STL Originating Programme Episode Title (OET).
type OETXML string

// OETXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*OETXML)(nil)
var _ xml.Marshaler = (*OETXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the OETXML pointed to by oet.
func (oet *OETXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(oet), d, start)
}

// MarshalXML returns the XML encoding of oet.
func (oet OETXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(oet), e, start, 32)
}

// TPTXML is the XML representation of STL Translated Programme Title (TPT).
type TPTXML string

// TPTXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TPTXML)(nil)
var _ xml.Marshaler = (*TPTXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TPTXML pointed to by tpt.
func (tpt *TPTXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(tpt), d, start)
}

// MarshalXML returns the XML encoding of tpt.
func (tpt TPTXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(tpt), e, start, 32)
}

// TETXML is the XML representation of STL Translated Programme Episode Title (TET).
type TETXML string

// TETXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TETXML)(nil)
var _ xml.Marshaler = (*TETXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TETXML pointed to by tet.
func (tet *TETXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(tet), d, start)
}

// MarshalXML returns the XML encoding of tet.
func (tet TETXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(tet), e, start, 32)
}

// TNXML is the XML representation of STL Translator's Name (TN).
type TNXML string

// TNXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TNXML)(nil)
var _ xml.Marshaler = (*TNXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TNXML pointed to by tn.
func (tn *TNXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(tn), d, start)
}

// MarshalXML returns the XML encoding of tn.
func (tn TNXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(tn), e, start, 32)
}

// TCDXML is the XML representation of STL Translator's Contact Details (TCD).
type TCDXML string

// TCDXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TCDXML)(nil)
var _ xml.Marshaler = (*TCDXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TCDXML pointed to by tcd.
func (tcd *TCDXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(tcd), d, start)
}

// MarshalXML returns the XML encoding of tcd.
func (tcd TCDXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(tcd), e, start, 32)
}

// SLRXML is the XML representation of STL Subtitle List Reference Code (SLR).
type SLRXML string

// SLRXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*SLRXML)(nil)
var _ xml.Marshaler = (*SLRXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the SLRXML pointed to by slr.
func (slr *SLRXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(slr), d, start)
}

// MarshalXML returns the XML encoding of slr.
func (slr SLRXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(slr), e, start, 16)
}

// CDXML is the XML representation of STL Creation Date (CD).
type CDXML time.Time

// CDXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*CDXML)(nil)
var _ xml.Marshaler = (*CDXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the CDXML pointed to by cd.
func (cd *CDXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLTime((*time.Time)(cd), d, start)
}

// MarshalXML returns the XML encoding of cd.
func (cd CDXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLTime(time.Time(cd), e, start)
}

// RDXML is the XML representation of STL Revision Date (RD).
type RDXML time.Time

// RDXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*RDXML)(nil)
var _ xml.Marshaler = (*RDXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the RDXML pointed to by rd.
func (rd *RDXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLTime((*time.Time)(rd), d, start)
}

// MarshalXML returns the XML encoding of rd.
func (rd RDXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLTime(time.Time(rd), e, start)
}

// RNXML is the XML representation of STL Revision Number (RN).
type RNXML int

// RNXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*RNXML)(nil)
var _ xml.Marshaler = (*RNXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the RNXML pointed to by rn.
func (rn *RNXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(rn), d, start)
}

// MarshalXML returns the XML encoding of rn.
func (rn RNXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(rn), e, start, 2)
}

// TNBXML is the XML representation of STL Total Number of Text and Timing Information (TTI) blocks (TNB).
type TNBXML int

// TNBXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TNBXML)(nil)
var _ xml.Marshaler = (*TNBXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TNBXML pointed to by tnb.
func (tnb *TNBXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(tnb), d, start)
}

// MarshalXML returns the XML encoding of tnb.
func (tnb TNBXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(tnb), e, start, 5)
}

// TNSXML is the XML representation of STL Total Number of Subtitles (TNS).
type TNSXML int

// TNSXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TNSXML)(nil)
var _ xml.Marshaler = (*TNSXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TNSXML pointed to by tns.
func (tns *TNSXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(tns), d, start)
}

// MarshalXML returns the XML encoding of tns.
func (tns TNSXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(tns), e, start, 5)
}

// TNGXML is the XML representation of STL Total Number of Subtitle Groups (TNG).
type TNGXML int

// TNGXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TNGXML)(nil)
var _ xml.Marshaler = (*TNGXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TNGXML pointed to by tng.
func (tng *TNGXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(tng), d, start)
}

// MarshalXML returns the XML encoding of tng.
func (tng TNGXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(tng), e, start, 3)
}

// MNCXML is the XML representation of STL Maximum Number of Displayable Characters in any text row (MNC).
type MNCXML int

// MNCXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*MNCXML)(nil)
var _ xml.Marshaler = (*MNCXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the MNCXML pointed to by mnc.
func (mnc *MNCXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(mnc), d, start)
}

// MarshalXML returns the XML encoding of mnc.
func (mnc MNCXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(mnc), e, start, 2)
}

// MNRXML is the XML representation of STL Maximum Number of Displayable Rows (MNR).
type MNRXML int

// MNRXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*MNRXML)(nil)
var _ xml.Marshaler = (*MNRXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the MNRXML pointed to by mnr.
func (mnr *MNRXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(mnr), d, start)
}

// MarshalXML returns the XML encoding of mnr.
func (mnr MNRXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(mnr), e, start, 2)
}

// TCSXML is the XML representation of STL Time Code: Status (TCS).
type TCSXML stl.TimeCodeStatus

// TCSXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TCSXML)(nil)
var _ xml.Marshaler = (*TCSXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TCSXML pointed to by tcs.
func (tcs *TCSXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLByte((*byte)(tcs), d, start)
}

// MarshalXML returns the XML encoding of tcs.
func (tcs TCSXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLByte(byte(tcs), e, start, 2)
}

// TCPXML is the XML representation of STL Time Code: Start-of-Program (TCP).
type TCPXML stl.Timecode

// TCPXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TCPXML)(nil)
var _ xml.Marshaler = (*TCPXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TCPXML pointed to by tcp.
func (tcp *TCPXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLTimecode((*stl.Timecode)(tcp), d, start)
}

// MarshalXML returns the XML encoding of tcp.
func (tcp TCPXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLTimecode(stl.Timecode(tcp), e, start)
}

// TCFXML is the XML representation of STL Time Code: First In-Cue (TCF).
type TCFXML stl.Timecode

// TCFXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TCFXML)(nil)
var _ xml.Marshaler = (*TCFXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TCFXML pointed to by tcf.
func (tcf *TCFXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLTimecode((*stl.Timecode)(tcf), d, start)
}

// MarshalXML returns the XML encoding of tcf.
func (tcf TCFXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLTimecode(stl.Timecode(tcf), e, start)
}

// TNDXML is the XML representation of STL Total Number of Disks (TND).
type TNDXML int

// TNDXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*TNDXML)(nil)
var _ xml.Marshaler = (*TNDXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the TNDXML pointed to by tnd.
func (tnd *TNDXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(tnd), d, start)
}

// MarshalXML returns the XML encoding of tnd.
func (tnd TNDXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(tnd), e, start, 1)
}

// DSNXML is the XML representation of STL Disk Sequence Number (DSN).
type DSNXML int

// DSNXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*DSNXML)(nil)
var _ xml.Marshaler = (*DSNXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the DSNXML pointed to by dsn.
func (dsn *DSNXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLInt((*int)(dsn), d, start)
}

// MarshalXML returns the XML encoding of dsn.
func (dsn DSNXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLInt(int(dsn), e, start, 1)
}

// COXML is the XML representation of STL Country of Origin (CO).
type COXML string

// COXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*COXML)(nil)
var _ xml.Marshaler = (*COXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the COXML pointed to by co.
func (co *COXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(co), d, start)
}

// MarshalXML returns the XML encoding of co.
func (co COXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(co), e, start, 3)
}

// PUBXML is the XML representation of STL Publisher (PUB).
type PUBXML string

// PUBXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*PUBXML)(nil)
var _ xml.Marshaler = (*PUBXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the PUBXML pointed to by pub.
func (pub *PUBXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(pub), d, start)
}

// MarshalXML returns the XML encoding of pub.
func (pub PUBXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(pub), e, start, 32)
}

// ENXML is the XML representation of STL Editor's Name (EN).
type ENXML string

// ENXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*ENXML)(nil)
var _ xml.Marshaler = (*ENXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the ENXML pointed to by en.
func (en *ENXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(en), d, start)
}

// MarshalXML returns the XML encoding of en.
func (en ENXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(en), e, start, 32)
}

// ECDXML is the XML representation of STL Editor's Contact Details (ECD).
type ECDXML string

// ECDXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*ECDXML)(nil)
var _ xml.Marshaler = (*ECDXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the ECDXML pointed to by ecd.
func (ecd *ECDXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLString((*string)(ecd), d, start)
}

// MarshalXML returns the XML encoding of ecd.
func (ecd ECDXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLString(string(ecd), e, start, 32)
}

// UDAXML is the XML representation of STL User-Defined Area (UDA).
type UDAXML []byte

// UDAXML implements xml.Unmarshaler and xml.Marshaler.
var _ xml.Unmarshaler = (*UDAXML)(nil)
var _ xml.Marshaler = (*UDAXML)(nil)

// UnmarshalXML decodes the XML-encoded data and stores the result in the UDAXML pointed to by uda.
func (uda *UDAXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return unmarshalXMLBytes((*[]byte)(uda), d, start)
}

// MarshalXML returns the XML encoding of uda.
func (uda UDAXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLBytes([]byte(uda), e, start)
}
