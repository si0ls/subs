package stl

import (
	"fmt"
	"io"
)

// TTIBlockSize is the size in bytes of a TTI block in a STL file.
const TTIBlockSize = 128

// ttiBlock is the Text and Timing Information (TTI) block representation.
type ttiBlock struct {
	SGN int               // Subtitle Group Number
	SN  int               // Subtitle Number
	EBN int               // Extension Block Number
	CS  CumulativeStatus  // Cumulative Status
	TCI Timecode          // Time Code In
	TCO Timecode          // Time Code Out
	VP  int               // Vertical Position
	JC  JustificationCode // Justification Code
	CF  CommentFlag       // Comment Flag
	TF  string            // Text Field

	terminatedBySpace bool // used for validation
}

// NewTTIBlock returns a new TTI block.
func NewTTIBlock() *ttiBlock {
	tti := ttiBlock{}
	tti.reset()
	return &tti
}

// Text returns the UTF-8 decoded Text Field (TF).
func (tti *ttiBlock) Text(cct CharacterCodeTable) (string, error) {
	if dec, ok := characterCodeTableDecoders[cct]; ok {
		b, err := dec.Decode([]byte(tti.TF))
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return "", fmt.Errorf("unsupported character code table %d", cct)
}

// SetText sets the Text Field (TF) from the UTF-8 encoded text.
func (tti *ttiBlock) SetText(text string, cct CharacterCodeTable) error {
	if enc, ok := characterCodeTableEncoders[cct]; ok {
		b, err := enc.Encode([]byte(text))
		if err != nil {
			return err
		}
		tti.TF = string(b)
		return nil
	}
	return fmt.Errorf("unsupported character code table %d", cct)
}

// Decode reads and decodes TTI block from reader.
func (tti *ttiBlock) Decode(r io.Reader) error {
	b := make([]byte, TTIBlockSize)
	if _, err := io.ReadFull(r, b); err != nil {
		return err
	}

	tti.reset()

	decodeTTIInt(b[0:1], &tti.SGN)            // Subtitle Group Number (SGN) - byte 0 (1 byte)
	decodeTTIInt(b[1:3], &tti.SN)             // Subtitle Number (SN) - bytes 1..2 (2 bytes)
	decodeTTIInt(b[3:4], &tti.EBN)            // Extension Block Number (EBN) - byte 3 (1 byte)
	decodeTTIByte(b[4:5], (*byte)(&tti.CS))   // Cumulative Status (CS) - byte 4 (1 byte)
	decodeTTITimecode(b[5:9], &tti.TCI)       // Time Code In (TCI) - bytes 5..8 (4 bytes)
	decodeTTITimecode(b[9:13], &tti.TCO)      // Time Code Out (TCO) - bytes 9..12 (4 bytes)
	decodeTTIInt(b[13:14], &tti.VP)           // Vertical Position (VP) - byte 13 (1 byte)
	decodeTTIByte(b[14:15], (*byte)(&tti.JC)) // Justification Code (JC) - byte 14 (1 byte)
	decodeTTIByte(b[15:16], (*byte)(&tti.CF)) // Comment Flag (CF) - byte 15 (1 byte)

	// Text Field (TF) - bytes 16..127 (112 bytes)
	tti.terminatedBySpace = b[127] == 0x8F
	decodeTTIString(b[16:128], &tti.TF)

	return nil
}

// Encode encodes and writes TTI block to writer.
func (tti *ttiBlock) Encode(w io.Writer) error {
	b := make([]byte, TTIBlockSize)

	encodeTTIInt(b[0:1], tti.SGN)           // Subtitle Group Number (SGN) - byte 0 (1 byte)
	encodeTTIInt(b[1:3], tti.SN)            // Subtitle Number (SN) - bytes 1..2 (2 bytes)
	encodeTTIInt(b[3:4], tti.EBN)           // Extension Block Number (EBN) - byte 3 (1 byte)
	encodeTTIByte(b[4:5], (byte)(tti.CS))   // Cumulative Status (CS) - byte 4 (1 byte)
	encodeTTITimecode(b[5:9], tti.TCI)      // Time Code In (TCI) - bytes 5..8 (4 bytes)
	encodeTTITimecode(b[9:13], tti.TCO)     // Time Code Out (TCO) - bytes 9..12 (4 bytes)
	encodeTTIInt(b[13:14], tti.VP)          // Vertical Position (VP) - byte 13 (1 byte)
	encodeTTIByte(b[14:15], (byte)(tti.JC)) // Justification Code (JC) - byte 14 (1 byte)
	encodeTTIByte(b[15:16], (byte)(tti.CF)) // Comment Flag (CF) - byte 15 (1 byte)
	encodeTTIString(b[16:128], tti.TF)      // Text Field (TF) - bytes 16..127 (112 bytes)

	_, err := w.Write(b)
	return err
}

// Validate validates TTI block.
// It returns a slice of ValidateErr containing warnings and fatal errors.
// If a ValidateErr is flaggued as fatal, then the TTI block is considered invalid.
// A warning will be returned if a field in TTI block have "unconventional" value.
// A fatal error will be returned if a field value make the future TTI processing impossible.
func (tti *ttiBlock) Validate() []error {
	var errs []error
	//todo: validation
	return errs
}

func (tti *ttiBlock) reset() {
	tti.SGN = -1
	tti.SN = -1
	tti.EBN = -1
	tti.CS = CumulativeStatusInvalid
	tti.TCI = Timecode{}
	tti.TCO = Timecode{}
	tti.VP = -1
	tti.JC = JustificationCodeInvalid
	tti.CF = CommentFlagInvalid
	tti.TF = ""
}
