package stlxml

import (
	"fmt"

	"github.com/si0ls/subs/stl"
	"golang.org/x/text/unicode/norm"
)

func decodeTextField(s string, cct stl.CharacterCodeTable) (string, error) {
	var strCopy string
	return strCopy, nil
}

func encodeTextField(s string, cct stl.CharacterCodeTable) (string, error) {
	var strCopy string

	var buf []byte
	for _, r := range []byte(s) {
		var handled bool
		for _, h := range handlers {
			if v := h.HandleChar(r); len(v) > 0 {
				trans, err := toUtf8(buf, cct)
				if err != nil {
					return "", err
				}
				strCopy += trans
				buf = []byte{}
				strCopy += string(v)
				handled = true
				break
			}
		}
		if !handled {
			buf = append(buf, r)
		}
	}
	trans, err := toUtf8(buf, cct)
	if err != nil {
		return "", err
	}
	strCopy += trans

	return strCopy, nil
}

func toUtf8(in []byte, cct stl.CharacterCodeTable) (string, error) {
	if dec, ok := stl.CharacterCodeTableDecoders[cct]; ok {
		b, err := dec.Decode(in)
		if err != nil {
			return "", err
		}
		return string(norm.NFC.Bytes(b)), nil
	}
	return "", fmt.Errorf("unknown character code table: %d", cct)
}

type charHandler interface {
	HandleChar(b byte) []byte
}

type controlCodeHandler struct{}

var _ charHandler = (*controlCodeHandler)(nil)

func (h *controlCodeHandler) HandleChar(b byte) []byte {
	if v, exists := stlControlCodeXmlTag[stl.ControlCode(b)]; exists {
		return []byte(fmt.Sprintf("<%s />", v))
	}
	return []byte{}
}

type spaceHandler struct{}

var _ charHandler = (*spaceHandler)(nil)

func (h *spaceHandler) HandleChar(b byte) []byte {
	if b == 0x20 || b == 0xA0 {
		return []byte("<space />")
	}
	return []byte{}
}

type quoteHandler struct{}

var _ charHandler = (*quoteHandler)(nil)

func (h *quoteHandler) HandleChar(b byte) []byte {
	if b == '"' {
		return []byte("&quot;")
	}
	return []byte{}
}

var handlers = []charHandler{
	&controlCodeHandler{},
	&spaceHandler{},
	&quoteHandler{},
}

var stlControlCodeXmlTag = map[stl.ControlCode]string{
	stl.ControlCodeItalicOn:                                  "ItalicOn",
	stl.ControlCodeItalicOff:                                 "ItalicOff",
	stl.ControlCodeUnderlineOn:                               "UnderlineOn",
	stl.ControlCodeUnderlineOff:                              "UnderlineOff",
	stl.ControlCodeBoxingOn:                                  "StartBox",
	stl.ControlCodeBoxingOff:                                 "EndBox",
	stl.ControlCodeLineBreak:                                 "newline",
	stl.ControlCodeUnusedSpace:                               "UnusedSpace",
	stl.ControlCode(stl.TeletextControlCodeAlphaBlack):       "AlphaBlack",
	stl.ControlCode(stl.TeletextControlCodeAlphaRed):         "AlphaRed",
	stl.ControlCode(stl.TeletextControlCodeAlphaGreen):       "AlphaGreen",
	stl.ControlCode(stl.TeletextControlCodeAlphaYellow):      "AlphaYellow",
	stl.ControlCode(stl.TeletextControlCodeAlphaBlue):        "AlphaBlue",
	stl.ControlCode(stl.TeletextControlCodeAlphaMagenta):     "AlphaMagenta",
	stl.ControlCode(stl.TeletextControlCodeAlphaCyan):        "AlphaCyan",
	stl.ControlCode(stl.TeletextControlCodeAlphaWhite):       "AlphaWhite",
	stl.ControlCode(stl.TeletextControlCodeFlash):            "Flash",
	stl.ControlCode(stl.TeletextControlCodeSteady):           "Steady",
	stl.ControlCode(stl.TeletextControlCodeEndBox):           "EndBox",
	stl.ControlCode(stl.TeletextControlCodeStartBox):         "StartBox",
	stl.ControlCode(stl.TeletextControlCodeNormalHeight):     "NormalHeight",
	stl.ControlCode(stl.TeletextControlCodeDoubleHeight):     "DoubleHeight",
	stl.ControlCode(stl.TeletextControlCodeDoubleWidth):      "DoubleWidth",
	stl.ControlCode(stl.TeletextControlCodeDoubleSize):       "DoubleSize",
	stl.ControlCode(stl.TeletextControlCodeMosaicBlack):      "MosaicBlack",
	stl.ControlCode(stl.TeletextControlCodeMosaicRed):        "MosaicRed",
	stl.ControlCode(stl.TeletextControlCodeMosaicGreen):      "MosaicGreen",
	stl.ControlCode(stl.TeletextControlCodeMosaicYellow):     "MosaicYellow",
	stl.ControlCode(stl.TeletextControlCodeMosaicBlue):       "MosaicBlue",
	stl.ControlCode(stl.TeletextControlCodeMosaicMagenta):    "MosaicMagenta",
	stl.ControlCode(stl.TeletextControlCodeMosaicCyan):       "MosaicCyan",
	stl.ControlCode(stl.TeletextControlCodeMosaicWhite):      "MosaicWhite",
	stl.ControlCode(stl.TeletextControlCodeConceal):          "Conceal",
	stl.ControlCode(stl.TeletextControlCodeContiguousMosaic): "ContiguousMosaic",
	stl.ControlCode(stl.TeletextControlCodeSeparatedMosaic):  "SeparatedMosaic",
	stl.ControlCode(stl.TeletextControlCodeBlackBackground):  "BlackBackground",
	stl.ControlCode(stl.TeletextControlCodeNewBackground):    "NewBackground",
	stl.ControlCode(stl.TeletextControlCodeHoldMosaic):       "HoldMosaic",
	stl.ControlCode(stl.TeletextControlCodeReleaseMosaic):    "ReleaseMosaic",
}

/*
var xmlTagStlControlCode = map[string]stl.ControlCode{
	"ItalicOn":         stl.ControlCodeItalicOn,
	"ItalicOff":        stl.ControlCodeItalicOff,
	"UnderlineOn":      stl.ControlCodeUnderlineOn,
	"UnderlineOff":     stl.ControlCodeUnderlineOff,
	"StartBox":         stl.ControlCodeBoxingOn,
	"EndBox":           stl.ControlCodeBoxingOff,
	"newline":          stl.ControlCodeLineBreak,
	"UnusedSpace":      stl.ControlCodeUnusedSpace,
	"AlphaBlack":       stl.ControlCode(stl.TeletextControlCodeAlphaBlack),
	"AlphaRed":         stl.ControlCode(stl.TeletextControlCodeAlphaRed),
	"AlphaGreen":       stl.ControlCode(stl.TeletextControlCodeAlphaGreen),
	"AlphaYellow":      stl.ControlCode(stl.TeletextControlCodeAlphaYellow),
	"AlphaBlue":        stl.ControlCode(stl.TeletextControlCodeAlphaBlue),
	"AlphaMagenta":     stl.ControlCode(stl.TeletextControlCodeAlphaMagenta),
	"AlphaCyan":        stl.ControlCode(stl.TeletextControlCodeAlphaCyan),
	"AlphaWhite":       stl.ControlCode(stl.TeletextControlCodeAlphaWhite),
	"Flash":            stl.ControlCode(stl.TeletextControlCodeFlash),
	"Steady":           stl.ControlCode(stl.TeletextControlCodeSteady),
	"EndBox":           stl.ControlCode(stl.TeletextControlCodeEndBox),
	"StartBox":         stl.ControlCode(stl.TeletextControlCodeStartBox),
	"NormalHeight":     stl.ControlCode(stl.TeletextControlCodeNormalHeight),
	"DoubleHeight":     stl.ControlCode(stl.TeletextControlCodeDoubleHeight),
	"DoubleWidth":      stl.ControlCode(stl.TeletextControlCodeDoubleWidth),
	"DoubleSize":       stl.ControlCode(stl.TeletextControlCodeDoubleSize),
	"MosaicBlack":      stl.ControlCode(stl.TeletextControlCodeMosaicBlack),
	"MosaicRed":        stl.ControlCode(stl.TeletextControlCodeMosaicRed),
	"MosaicGreen":      stl.ControlCode(stl.TeletextControlCodeMosaicGreen),
	"MosaicYellow":     stl.ControlCode(stl.TeletextControlCodeMosaicYellow),
	"MosaicBlue":       stl.ControlCode(stl.TeletextControlCodeMosaicBlue),
	"MosaicMagenta":    stl.ControlCode(stl.TeletextControlCodeMosaicMagenta),
	"MosaicCyan":       stl.ControlCode(stl.TeletextControlCodeMosaicCyan),
	"MosaicWhite":      stl.ControlCode(stl.TeletextControlCodeMosaicWhite),
	"Conceal":          stl.ControlCode(stl.TeletextControlCodeConceal),
	"ContiguousMosaic": stl.ControlCode(stl.TeletextControlCodeContiguousMosaic),
	"SeparatedMosaic":  stl.ControlCode(stl.TeletextControlCodeSeparatedMosaic),
	"BlackBackground":  stl.ControlCode(stl.TeletextControlCodeBlackBackground),
	"NewBackground":    stl.ControlCode(stl.TeletextControlCodeNewBackground),
	"HoldMosaic":       stl.ControlCode(stl.TeletextControlCodeHoldMosaic),
	"ReleaseMosaic":    stl.ControlCode(stl.TeletextControlCodeReleaseMosaic),
}
*/
