package stl

import (
	"encoding/xml"
	"fmt"
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

func encodeStlXmlInt(i int, pad int) string {
	if i < 0 {
		return ""
	}
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", pad), i)
}

func encodeStlXmlHexa(i int, pad int) string {
	if i >= 0xFF {
		return ""
	}
	return fmt.Sprintf(fmt.Sprintf("%%0%dX", pad), i)
}

func encodeStlXmlByte(b byte, pad int) string {
	if b >= 0xFF {
		return ""
	}
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", pad), b)
}

func encodeStlXmlElement(e *xml.Encoder, v any, tag string) error {
	return e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: tag}})
}

func (gsi *GSIBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if err := encodeStlXmlElement(e, encodeStlXmlInt(int(gsi.CPN), 3), "CPN"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.DFC, "DFC"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlByte(byte(gsi.DSC), 1), "DSC"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlByte(byte(gsi.CCT), 2), "CCT"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlHexa(int(gsi.LC), 2), "LC"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.OPT, "OPT"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.OET, "OET"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.TPT, "TPT"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.TET, "TET"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.TN, "TN"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.TCD, "TCD"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.SLR, "SLR"); err != nil {
		return err
	}
	if !gsi.CD.IsZero() {
		if err := encodeStlXmlElement(e, gsi.CD.Format("060201"), "CD"); err != nil {
			return err
		}
	} else {
		if err := encodeStlXmlElement(e, "", "CD"); err != nil {
			return err
		}
	}
	if !gsi.RD.IsZero() {
		if err := encodeStlXmlElement(e, gsi.RD.Format("060201"), "RD"); err != nil {
			return err
		}
	} else {
		if err := encodeStlXmlElement(e, "", "RD"); err != nil {
			return err
		}
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(gsi.RN, 2), "RN"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(gsi.TNB, 5), "TNB"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(gsi.TNS, 5), "TNS"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(gsi.TNG, 3), "TNG"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(gsi.MNC, 2), "MNC"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(gsi.MNR, 2), "MNR"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlByte(byte(gsi.TCS), 1), "TCS"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.TCP, "TCP"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.TCF, "TCF"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(gsi.TND, 1), "TND"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, encodeStlXmlInt(gsi.DSN, 1), "DSN"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.CO, "CO"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.PUB, "PUB"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.EN, "EN"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, gsi.ECD, "ECD"); err != nil {
		return err
	}
	if err := encodeStlXmlElement(e, string(gsi.UDA), "UDA"); err != nil {
		return err
	}

	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}

	return nil
}
