package stl

// ControlCode
type ControlCode byte

const (
	ControlCodeItalicOn     ControlCode = 0x80
	ControlCodeItalicOff    ControlCode = 0x81
	ControlCodeUnderlineOn  ControlCode = 0x82
	ControlCodeUnderlineOff ControlCode = 0x83
	ControlCodeBoxingOn     ControlCode = 0x84
	ControlCodeBoxingOff    ControlCode = 0x85
	ControlCodeLineBreak    ControlCode = 0x8A
	ControlCodeUnusedSpace  ControlCode = 0x8F
)

// String returns the string representation of ControlCode.
func (cc ControlCode) String() string {
	switch cc {
	case ControlCodeItalicOn:
		return "Italic on"
	case ControlCodeItalicOff:
		return "Italic off"
	case ControlCodeUnderlineOn:
		return "Underline on"
	case ControlCodeUnderlineOff:
		return "Underline off"
	case ControlCodeBoxingOn:
		return "Boxing on"
	case ControlCodeBoxingOff:
		return "Boxing off"
	case ControlCodeLineBreak:
		return "Line break"
	case ControlCodeUnusedSpace:
		return "Unused space"
	}
	return "Unknown"
}

// TeletextControlCode (from EBU document Tech. 3240)
type TeletextControlCode ControlCode

const (
	TeletextControlCodeAlphaBlack       TeletextControlCode = 0x00
	TeletextControlCodeAlphaRed         TeletextControlCode = 0x01
	TeletextControlCodeAlphaGreen       TeletextControlCode = 0x02
	TeletextControlCodeAlphaYellow      TeletextControlCode = 0x03
	TeletextControlCodeAlphaBlue        TeletextControlCode = 0x04
	TeletextControlCodeAlphaMagenta     TeletextControlCode = 0x05
	TeletextControlCodeAlphaCyan        TeletextControlCode = 0x06
	TeletextControlCodeAlphaWhite       TeletextControlCode = 0x07
	TeletextControlCodeFlash            TeletextControlCode = 0x08
	TeletextControlCodeSteady           TeletextControlCode = 0x09
	TeletextControlCodeEndBox           TeletextControlCode = 0x0A
	TeletextControlCodeStartBox         TeletextControlCode = 0x0B
	TeletextControlCodeNormalHeight     TeletextControlCode = 0x0C
	TeletextControlCodeDoubleHeight     TeletextControlCode = 0x0D
	TeletextControlCodeDoubleWidth      TeletextControlCode = 0x0E
	TeletextControlCodeDoubleSize       TeletextControlCode = 0x0F
	TeletextControlCodeMosaicBlack      TeletextControlCode = 0x10
	TeletextControlCodeMosaicRed        TeletextControlCode = 0x11
	TeletextControlCodeMosaicGreen      TeletextControlCode = 0x12
	TeletextControlCodeMosaicYellow     TeletextControlCode = 0x13
	TeletextControlCodeMosaicBlue       TeletextControlCode = 0x14
	TeletextControlCodeMosaicMagenta    TeletextControlCode = 0x15
	TeletextControlCodeMosaicCyan       TeletextControlCode = 0x16
	TeletextControlCodeMosaicWhite      TeletextControlCode = 0x17
	TeletextControlCodeConceal          TeletextControlCode = 0x18
	TeletextControlCodeContiguousMosaic TeletextControlCode = 0x19
	TeletextControlCodeSeparatedMosaic  TeletextControlCode = 0x1A
	TeletextControlCodeReserved         TeletextControlCode = 0x1B
	TeletextControlCodeBlackBackground  TeletextControlCode = 0x1C
	TeletextControlCodeNewBackground    TeletextControlCode = 0x1D
	TeletextControlCodeHoldMosaic       TeletextControlCode = 0x1E
	TeletextControlCodeReleaseMosaic    TeletextControlCode = 0x1F
)

// String returns the string representation of TeletextControlCode.
func (c TeletextControlCode) String() string {
	switch c {
	case TeletextControlCodeAlphaBlack:
		return "Alpha black"
	case TeletextControlCodeAlphaRed:
		return "Alpha red"
	case TeletextControlCodeAlphaGreen:
		return "Alpha green"
	case TeletextControlCodeAlphaYellow:
		return "Alpha yellow"
	case TeletextControlCodeAlphaBlue:
		return "Alpha blue"
	case TeletextControlCodeAlphaMagenta:
		return "Alpha magenta"
	case TeletextControlCodeAlphaCyan:
		return "Alpha cyan"
	case TeletextControlCodeAlphaWhite:
		return "Alpha white"
	case TeletextControlCodeFlash:
		return "Flash"
	case TeletextControlCodeSteady:
		return "Steady"
	case TeletextControlCodeEndBox:
		return "End box"
	case TeletextControlCodeStartBox:
		return "Start box"
	case TeletextControlCodeNormalHeight:
		return "Normal height"
	case TeletextControlCodeDoubleHeight:
		return "Double height"
	case TeletextControlCodeDoubleWidth:
		return "Double width"
	case TeletextControlCodeDoubleSize:
		return "Double size"
	case TeletextControlCodeMosaicBlack:
		return "Mosaic black"
	case TeletextControlCodeMosaicRed:
		return "Mosaic red"
	case TeletextControlCodeMosaicGreen:
		return "Mosaic green"
	case TeletextControlCodeMosaicYellow:
		return "Mosaic yellow"
	case TeletextControlCodeMosaicBlue:
		return "Mosaic blue"
	case TeletextControlCodeMosaicMagenta:
		return "Mosaic magenta"
	case TeletextControlCodeMosaicCyan:
		return "Mosaic cyan"
	case TeletextControlCodeMosaicWhite:
		return "Mosaic white"
	case TeletextControlCodeConceal:
		return "Conceal"
	case TeletextControlCodeContiguousMosaic:
		return "Contiguous mosaic"
	case TeletextControlCodeSeparatedMosaic:
		return "Separated mosaic"
	case TeletextControlCodeReserved:
		return "Reserved"
	case TeletextControlCodeBlackBackground:
		return "Black background"
	case TeletextControlCodeNewBackground:
		return "New background"
	case TeletextControlCodeHoldMosaic:
		return "Hold mosaic"
	case TeletextControlCodeReleaseMosaic:
		return "Release mosaic"
	}
	return "Unknown"
}

// TeletextColor is the code used to represent a Teletext color (from EBU document Tech. 3240)
type TeletextColor TeletextControlCode

const (
	TeletextColorBlack   TeletextColor = 0x0
	TeletextColorRed     TeletextColor = 0x1
	TeletextColorGreen   TeletextColor = 0x2
	TeletextColorYellow  TeletextColor = 0x3
	TeletextColorBlue    TeletextColor = 0x4
	TeletextColorMagenta TeletextColor = 0x5
	TeletextColorCyan    TeletextColor = 0x6
	TeletextColorWhite   TeletextColor = 0x7
)

// String returns the string representation of TeletextColor.
func (c TeletextColor) String() string {
	switch c {
	case TeletextColorBlack:
		return "Black"
	case TeletextColorRed:
		return "Red"
	case TeletextColorGreen:
		return "Green"
	case TeletextColorYellow:
		return "Yellow"
	case TeletextColorBlue:
		return "Blue"
	case TeletextColorMagenta:
		return "Magenta"
	case TeletextColorCyan:
		return "Cyan"
	case TeletextColorWhite:
		return "White"
	}
	return "Unknown"
}
