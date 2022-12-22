package stl

import (
	"fmt"
	"io"
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

// Decode reads and decodes TTI block from reader.
func (tti *TTIBlock) Decode(r io.Reader) error {
	b := make([]byte, TTIBlockSize)
	if _, err := io.ReadFull(r, b); err != nil {
		return err
	}

	tti.Reset()

	decodeTTIInt(b[0:1], &tti.SGN)            // SGN - byte 0 (1 byte)
	decodeTTIInt(b[1:3], &tti.SN)             // SN - bytes 1..2 (2 bytes)
	decodeTTIInt(b[3:4], &tti.EBN)            // EBN - byte 3 (1 byte)
	decodeTTIByte(b[4:5], (*byte)(&tti.CS))   // CS - byte 4 (1 byte)
	decodeTTITimecode(b[5:9], &tti.TCI)       // TCI - bytes 5..8 (4 bytes)
	decodeTTITimecode(b[9:13], &tti.TCO)      // TCO - bytes 9..12 (4 bytes)
	decodeTTIInt(b[13:14], &tti.VP)           // VP - byte 13 (1 byte)
	decodeTTIByte(b[14:15], (*byte)(&tti.JC)) // JC - byte 14 (1 byte)
	decodeTTIByte(b[15:16], (*byte)(&tti.CF)) // CF - byte 15 (1 byte)
	tti.terminatedBySpace = b[127] == 0x8F    // Store if the last byte is 0x8F (space) for further validation
	decodeTTIString(b[16:128], &tti.TF)       // Text Field (TF) - bytes 16..127 (112 bytes)

	return nil
}

// Encode encodes and writes TTI block to writer.
func (tti *TTIBlock) Encode(w io.Writer) error {
	b := make([]byte, TTIBlockSize)

	encodeTTIInt(b[0:1], tti.SGN)           // SGN - byte 0 (1 byte)
	encodeTTIInt(b[1:3], tti.SN)            // SN - bytes 1..2 (2 bytes)
	encodeTTIInt(b[3:4], tti.EBN)           // EBN - byte 3 (1 byte)
	encodeTTIByte(b[4:5], (byte)(tti.CS))   // CS - byte 4 (1 byte)
	encodeTTITimecode(b[5:9], tti.TCI)      // TCI - bytes 5..8 (4 bytes)
	encodeTTITimecode(b[9:13], tti.TCO)     // TCO - bytes 9..12 (4 bytes)
	encodeTTIInt(b[13:14], tti.VP)          // VP - byte 13 (1 byte)
	encodeTTIByte(b[14:15], (byte)(tti.JC)) // JC - byte 14 (1 byte)
	encodeTTIByte(b[15:16], (byte)(tti.CF)) // CF - byte 15 (1 byte)
	encodeTTIString(b[16:128], tti.TF)      // TF - bytes 16..127 (112 bytes)

	_, err := w.Write(b)
	return err
}

// Validate validates TTI block.
// It returns a slice of warnings and an error if any.
// Warnings are returned for each field that is invalid, warnings can flagged
// as fatal if they are considered to be fatal to further file processing.
func (tti *TTIBlock) Validate(framerate uint, dsc DisplayStandardCode, mnr int) []error {
	var errs []error

	// SGN - between 0 and 0xFF -> Fatal
	errs = appendNonNilErrs(errs, ttiErr(validateRange(tti.SGN, 0, 0xFF, ErrUnsupportedSGN, true), TTIFieldSGN))

	// SN - between 0 and 0xFFFF -> Fatal
	errs = appendNonNilErrs(errs, ttiErr(validateRange(tti.SN, 0, 0xFFFF, ErrUnsupportedSN, true), TTIFieldSN))

	// EBN - if 0xFF the TF field must be terminated by a space
	if tti.EBN == 0xFF && !tti.terminatedBySpace {
		errs = appendNonNilErrs(errs, ttiErr(validateErr(ErrLastEBNNotTerminatedBySpace, tti.EBN, false), TTIFieldEBN))
	}

	// EBN - not in reserved range (0xF0..0xFD)
	errs = appendNonNilErrs(errs, ttiErr(validateNotInRange(tti.EBN, 0xF0, 0xFD, ErrReservedEBNRange, false), TTIFieldEBN))

	// CS - in list
	errs = appendNonNilErrs(errs, ttiErr(validateList(tti.CS, csValidValues, ErrUnsupportedCS, false), TTIFieldCS))

	// TCI - valid -> fatal
	errs = appendNonNilErrs(errs, ttiErr(validateTimecode(tti.TCI, framerate, ErrInvalidTCI, true), TTIFieldTCI))

	// TCO - valid -> fatal
	errs = appendNonNilErrs(errs, ttiErr(validateTimecode(tti.TCO, framerate, ErrInvalidTCO, true), TTIFieldTCO))

	// Timecodes (TCI, TCO) - TCI < TCO -> fatal
	errs = appendNonNilErrs(errs, ttiErr(validateTimecodeOrderStrict(tti.TCI, tti.TCO, framerate, ErrInvalidTCITCOOrder, true), TTIFieldTCO))

	// VP - between 1 and 23 if teletext, between 0 and MNR if open subtitles, otherwise fatal
	if dsc == DisplayStandardCodeLevel1Teletext || dsc == DisplayStandardCodeLevel2Teletext {
		errs = appendNonNilErrs(errs, ttiErr(validateRange(tti.VP, 1, 23, ErrUnsupportedVPTeletext, false), TTIFieldVP))
	} else if dsc == DisplayStandardCodeOpenSubtitling {
		errs = appendNonNilErrs(errs, ttiErr(validateRange(tti.VP, 0, mnr, ErrUnsupportedVPOpenSubtitling, false), TTIFieldVP))
	} else {
		errs = appendNonNilErrs(errs, ttiErr(validateErr(ErrUnsupportedDSC, dsc, true), TTIFieldVP))
	}

	// JC - in list
	errs = appendNonNilErrs(errs, ttiErr(validateList(tti.JC, jcValidValues, ErrUnsupportedJC, false), TTIFieldJC))

	// CF - in list
	errs = appendNonNilErrs(errs, ttiErr(validateList(tti.CF, cfValidValues, ErrUnsupportedCF, false), TTIFieldCF))

	// TF - no teletext chars if open subtitles, no open subtitles chars if teletext
	//todo: validation

	// TF - out of boxes
	//todo: validation

	// TF - respects MNC
	//todo: validation

	return errs
}

// TTIError is an error that occurred on a TTI field.
// It extends FieldError that carries the concerned TTI field.
// It carries the concerned TTI block number.
// If TTI block number is -1, it means that the TTI block number is unknown.
type TTIError struct {
	error
	FieldError
	field       TTIField
	blockNumber int
}

func ttiErr(err error, field TTIField) error {
	if err == nil {
		return nil
	}
	return &TTIError{error: err, field: field}
}

func ttiErrWithBlockNumber(err error, field TTIField, blockNumber int) error {
	if err == nil {
		return nil
	}
	return &TTIError{error: err, field: field, blockNumber: blockNumber}
}

// Error returns the error message.
func (e *TTIError) Error() string {
	return fmt.Sprintf("GSI %s: %s", e.field, e.error.Error())
}

// Unwrap returns the underlying error.
func (e *TTIError) Unwrap() error {
	return e.error
}

// Field returns the concerned TTI field.
func (e *TTIError) Field() TTIField {
	return e.field
}

// BlockNumber returns the concerned TTI block number.
// If TTI block number is -1, it means that the TTI block number is unknown.
func (e *TTIError) BlockNumber() int {
	return e.blockNumber
}

func (e *TTIError) setBlockNumber(blockNumber int) {
	e.blockNumber = blockNumber
}

func setTTIErrsBlockNumber(errs []error, blockNumber int) {
	for _, err := range errs {
		if ttiErr, ok := err.(*TTIError); ok {
			ttiErr.setBlockNumber(blockNumber)
		}
	}
}
