package stl

import (
	"errors"
	"fmt"
)

// Validate validates TTI block.
// It returns a slice of warnings and an error if any.
// Warnings are returned for each field that is invalid, warnings can flagged
// as fatal if they are considered to be fatal to further file processing.
// An error is returned if a field is invalid and prevents validation of
// further fields.
func (tti *TTIBlock) Validate(framerate uint, dsc DisplayStandardCode, mnr int) ([]error, error) {
	var warns []error

	// input framerate - 25 or 30 -> exit
	if framerate != 25 && framerate != 30 {
		return warns, validateErr(fmt.Errorf("%w: must be 25 or 30, prevents tti validation", ErrUnsupportedFramerate), framerate, true)
	}

	// input DSC - in list -> exit
	if err := validateList(dsc, dscValidValues, fmt.Errorf("%w: prevents tti validation", ErrUnsupportedDSC), false); err != nil {
		return warns, err
	}

	// input MNR - between 0 and 99 (regardless of DSC) -> exit
	if err := validateRange(mnr, 0, 99, fmt.Errorf("%w: prevents tti validation", ErrUnsupportedMNROpenSubtitling), false); err != nil {
		return warns, err
	}

	// SGN - between 0 and 0xFF -> Fatal
	warns = appendNonNilErrs(warns, ttiErr(validateRange(tti.SGN, 0, 0xFF, ErrUnsupportedSGN, true), TTIFieldSGN))

	// SN - between 0 and 0xFFFF -> Fatal
	warns = appendNonNilErrs(warns, ttiErr(validateRange(tti.SN, 0, 0xFFFF, ErrUnsupportedSN, true), TTIFieldSN))

	// EBN - if 0xFF the TF field must be terminated by a space
	if tti.EBN == 0xFF && !tti.terminatedBySpace {
		warns = appendNonNilErrs(warns, ttiErr(validateErr(ErrLastEBNNotTerminatedBySpace, tti.EBN, false), TTIFieldEBN))
	}

	// EBN - not in reserved range (0xF0..0xFD)
	warns = appendNonNilErrs(warns, ttiErr(validateNotInRange(tti.EBN, 0xF0, 0xFD, ErrReservedEBNRange, false), TTIFieldEBN))

	// CS - in list
	warns = appendNonNilErrs(warns, ttiErr(validateList(tti.CS, csValidValues, ErrUnsupportedCS, false), TTIFieldCS))

	// TCI - valid -> fatal
	warns = appendNonNilErrs(warns, ttiErr(validateTimecode(tti.TCI, framerate, ErrInvalidTCI, true), TTIFieldTCI))

	// TCO - valid -> fatal
	warns = appendNonNilErrs(warns, ttiErr(validateTimecode(tti.TCO, framerate, ErrInvalidTCO, true), TTIFieldTCO))

	// Timecodes (TCI, TCO) - TCI < TCO -> fatal
	warns = appendNonNilErrs(warns, ttiErr(validateTimecodeOrderStrict(tti.TCI, tti.TCO, framerate, ErrInvalidTCITCOOrder, true), TTIFieldTCO))

	// VP - between 1 and 23 if teletext, between 0 and MNR if open subtitles, otherwise fatal
	if dsc == DisplayStandardCodeLevel1Teletext || dsc == DisplayStandardCodeLevel2Teletext {
		warns = appendNonNilErrs(warns, ttiErr(validateRange(tti.VP, 1, 23, ErrUnsupportedVPTeletext, false), TTIFieldVP))
	} else if dsc == DisplayStandardCodeOpenSubtitling {
		warns = appendNonNilErrs(warns, ttiErr(validateRange(tti.VP, 0, mnr, ErrUnsupportedVPOpenSubtitling, false), TTIFieldVP))
	} else {
		warns = appendNonNilErrs(warns, ttiErr(validateErr(ErrUnsupportedDSC, dsc, true), TTIFieldVP))
	}

	// JC - in list
	warns = appendNonNilErrs(warns, ttiErr(validateList(tti.JC, jcValidValues, ErrUnsupportedJC, false), TTIFieldJC))

	// CF - in list
	warns = appendNonNilErrs(warns, ttiErr(validateList(tti.CF, cfValidValues, ErrUnsupportedCF, false), TTIFieldCF))

	// TF - no teletext chars if open subtitles, no open subtitles chars if teletext
	//todo: validation

	// TF - out of boxes
	//todo: validation

	// TF - respects MNC
	//todo: validation

	return warns, nil
}

var (
	ErrInvalidFramerate = errors.New("invalid framerate")

	ErrUnsupportedSGN              = errors.New("unsupported SGN")
	ErrUnsupportedSN               = errors.New("unsupported SN")
	ErrLastEBNNotTerminatedBySpace = errors.New("last EBN not terminated by space")
	ErrReservedEBNRange            = errors.New("reserved EBN range")
	ErrUnsupportedCS               = errors.New("unsupported CS")
	ErrInvalidTCI                  = errors.New("invalid TCI")
	ErrInvalidTCO                  = errors.New("invalid TCO")
	ErrInvalidTCITCOOrder          = errors.New("invalid TCI/TCO order")
	ErrUnsupportedVPTeletext       = errors.New("unsupported VP (teletext)")
	ErrUnsupportedVPOpenSubtitling = errors.New("unsupported VP (for open subtitling)")
	ErrUnsupportedVPDSC            = errors.New("unsupported DSC, cannot use VP")
	ErrUnsupportedJC               = errors.New("unsupported JC")
	ErrUnsupportedCF               = errors.New("unsupported CF")
)

var csValidValues = []CumulativeStatus{
	CumulativeStatusNone,
	CumulativeStatusFirst,
	CumulativeStatusIntermediate,
	CumulativeStatusLast,
}

var jcValidValues = []JustificationCode{
	JustificationCodeUnchangedPresentation,
	JustificationCodeLeftJustifiedText,
	JustificationCodeCenteredText,
	JustificationCodeRightJustifiedText,
}

var cfValidValues = []CommentFlag{
	CommentFlagSubtitleData,
	CommentFlagTranslatorComments,
}
