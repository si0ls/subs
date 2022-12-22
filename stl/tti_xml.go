package stl

import (
	"encoding/xml"
	"fmt"
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

func encodeStlXmlTTIHexa(i int, pad int) string {
	if i <= 0 {
		return ""
	}
	return fmt.Sprintf(fmt.Sprintf("%%0%dx", pad), i)
}

var controlCodeStlXmlTag = map[ControlCode]string{

	ControlCodeItalicOn:     "",
	ControlCodeItalicOff:    "",
	ControlCodeUnderlineOn:  "",
	ControlCodeUnderlineOff: "",

	ControlCodeBoxingOn:    "StartBox",
	ControlCodeBoxingOff:   "EndBox",
	ControlCodeLineBreak:   "newline",
	ControlCodeUnusedSpace: "",
}

var teletextControlCodeStlXmlTag = map[TeletextControlCode]string{

	TeletextControlCodeAlphaBlack:       "AlphaBlack",
	TeletextControlCodeAlphaRed:         "AlphaRed",
	TeletextControlCodeAlphaGreen:       "AlphaGreen",
	TeletextControlCodeAlphaYellow:      "AlphaYellow",
	TeletextControlCodeAlphaBlue:        "AlphaBlue",
	TeletextControlCodeAlphaMagenta:     "AlphaMagenta",
	TeletextControlCodeAlphaCyan:        "AlphaCyan",
	TeletextControlCodeAlphaWhite:       "AlphaWhite",
	TeletextControlCodeFlash:            "Flash",
	TeletextControlCodeSteady:           "Steady",
	TeletextControlCodeEndBox:           "EndBox",
	TeletextControlCodeStartBox:         "StartBox",
	TeletextControlCodeNormalHeight:     "NormalHeight",
	TeletextControlCodeDoubleHeight:     "DoubleHeight",
	TeletextControlCodeDoubleWidth:      "DoubleWidth",
	TeletextControlCodeDoubleSize:       "DoubleSize",
	TeletextControlCodeMosaicBlack:      "MosaicBlack",
	TeletextControlCodeMosaicRed:        "MosaicRed",
	TeletextControlCodeMosaicGreen:      "MosaicGreen",
	TeletextControlCodeMosaicYellow:     "MosaicYellow",
	TeletextControlCodeMosaicBlue:       "MosaicBlue",
	TeletextControlCodeMosaicMagenta:    "MosaicMagenta",
	TeletextControlCodeMosaicCyan:       "MosaicCyan",
	TeletextControlCodeMosaicWhite:      "MosaicWhite",
	TeletextControlCodeConceal:          "Conceal",
	TeletextControlCodeContiguousMosaic: "ContiguousMosaic",
	TeletextControlCodeSeparatedMosaic:  "SeparatedMosaic",
	TeletextControlCodeBlackBackground:  "BlackBackground",
	TeletextControlCodeNewBackground:    "NewBackground",
	TeletextControlCodeHoldMosaic:       "HoldMosaic",
	TeletextControlCodeReleaseMosaic:    "ReleaseMosaic",
}

func encodeStlXmlTextField(e *xml.Encoder, s string) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "TF"}})

	for _, c := range []byte(s) {
		if tag, ok := controlCodeStlXmlTag[ControlCode(byte(c))]; ok {
			e.EncodeToken(xml.StartElement{Name: xml.Name{Local: tag}})
			e.EncodeToken(xml.EndElement{Name: xml.Name{Local: tag}})

		} else if tag, ok := teletextControlCodeStlXmlTag[TeletextControlCode(byte(c))]; ok {
			e.EncodeToken(xml.StartElement{Name: xml.Name{Local: tag}})
			e.EncodeToken(xml.EndElement{Name: xml.Name{Local: tag}})
		} else if c == 0x20 {
			e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "space"}})
			e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "space"}})
		} else {
			e.EncodeToken(xml.CharData([]byte{c}))
		}
	}

	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "TF"}})

	return nil
}

func (tti *TTIBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(tti.SGN, 1), "SGN"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(tti.SN, 1), "SN"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlTTIHexa(tti.EBN, 2), "EBN"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlHexa(int(tti.CS), 2), "CS"); err != nil {
		return err
	}

	if err := encodeStlXmlElement(e, tti.TCI, "TCI"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, tti.TCO, "TCO"); err != nil {
		return err
	}

	if err := encodeStlXmlElement(e, encodeStlXmlInt(tti.VP, 1), "VP"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlHexa(int(tti.JC), 2), "JC"); err != nil {
		return err
	}

	if err := encodeStlXmlElement(e, encodeStlXmlHexa(int(tti.CF), 2), "CF"); err != nil {
		return err
	}

	if err := encodeStlXmlTextField(e, tti.TF); err != nil {
		return err
	}

	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}

	return nil
}
