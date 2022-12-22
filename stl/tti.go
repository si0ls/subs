package stl

import (
	"fmt"
)

// TTIBlockSize is the size in bytes of a TTI block in a STL file.
const TTIBlockSize = 128

// TTIBlock is the Text and Timing Information (TTI) block representation.
type TTIBlock struct {
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
func NewTTIBlock() *TTIBlock {
	tti := TTIBlock{}
	tti.Reset()
	return &tti
}

// Text returns the UTF-8 decoded Text Field (TF).
func (tti *TTIBlock) Text(cct CharacterCodeTable) (string, error) {
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
func (tti *TTIBlock) SetText(text string, cct CharacterCodeTable) error {
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

// Reset resets the TTI block to its default values.
func (tti *TTIBlock) Reset() {
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
