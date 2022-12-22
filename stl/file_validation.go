package stl

import (
	"errors"
	"fmt"
)

// Validate validates STL file.
// It returns a slice of warnings and an error if any.
// Warnings are returned for each field that is invalid, warnings can flagged
// as fatal if they are considered to be fatal to further file processing.
// An error is returned if a field is invalid and prevents validation of
// further fields.
func (f *File) Validate() ([]error, error) {
	var warns []error

	// non nil GSI block
	if f.GSI == nil {
		panic(fmt.Errorf("GSI block is nil"))
	}

	// validate GSI block
	gsiWarns, gsiErr := f.GSI.Validate()
	if gsiErr != nil {
		return warns, gsiErr
	}
	warns = appendNonNilErrs(warns, gsiWarns...)

	// having at least one TTI block
	if len(f.TTI) == 0 {
		return warns, validateErr(ErrNoTTIBlocks, nil, true)
	}

	// match between gsi.TNB and len(f.TTI)
	if f.GSI.TNB != len(f.TTI) {
		warns = append(warns, validateErr(ErrTTIBlocksCountMismatch, nil, true))
	}

	// check GSI TSF timecode and first TTI TCI timecode
	if f.GSI.TCF != f.TTI[0].TCI {
		warns = append(warns, validateErr(ErrFirstTCMismatch, nil, true))
	}

	var subtitles int
	var groups int

	var lastSN int = 0
	var lastSGN int = f.TTI[0].SGN
	var lastEBN int = 0xFF
	var lastCS CumulativeStatus = CumulativeStatusNone

	for i, tti := range f.TTI {
		// non nil TTI block
		if tti == nil {
			panic(fmt.Errorf("TTI block %d is nil", i))
		}

		// validate TTI block
		ttiWarns, ttiErr := tti.Validate(f.GSI.Framerate(), f.GSI.DSC, f.GSI.MNR)
		if ttiErr != nil {
			return warns, ttiErr
		}
		setTTIErrsBlockNumber(ttiWarns, i)
		warns = appendNonNilErrs(warns, ttiWarns...)

		// --------------------

		// same subtitle (same group)
		if tti.SN == lastSN && tti.SGN == lastSGN {
			// check EBN is consecutive
			if tti.EBN != lastEBN+1 {
				warns = append(warns, validateErr(ErrEBNNotConsecutive, nil, true))
			}
		} else { // new subtitle (same group or not)
			subtitles++
		}

		// new subtitle, same group
		if tti.SN != lastSN && tti.SGN == lastSGN {
			// check SN is consecutive
			if tti.SN != lastSN+1 {
				warns = append(warns, validateErr(ErrSNNotConsecutive, nil, true))
			}

			// closing EBN
			if tti.EBN != 0xFF {
				warns = append(warns, validateErr(ErrNonClosingEBNForLastSubtitle, nil, true))
			}

			// check CS
			switch lastCS {
			case CumulativeStatusNone: // if last CS was None, then CS must be None or First
				if tti.CS != CumulativeStatusNone && tti.CS != CumulativeStatusFirst {
					warns = append(warns, validateErr(ErrCSNotNoneOrFirst, nil, true))
				}
			case CumulativeStatusIntermediate: // if last was Intermediate, then CS must be Intermediate or Last
				if tti.CS != CumulativeStatusIntermediate && tti.CS != CumulativeStatusLast {
					warns = append(warns, validateErr(ErrCSNotIntermediateOrLast, nil, true))
				}
			case CumulativeStatusFirst: // if last was First, then CS must be Intermediate or Last
				if tti.CS != CumulativeStatusIntermediate && tti.CS != CumulativeStatusLast {
					warns = append(warns, validateErr(ErrCSNotIntermediateOrLast, nil, true))
				}
			case CumulativeStatusLast: // if last was Last, then CS must be None or Last
				if tti.CS != CumulativeStatusNone && tti.CS != CumulativeStatusLast {
					warns = append(warns, validateErr(ErrCSNotNoneOrLast, nil, true))
				}
			}
		}

		// new group (new subtitle even if SN is the same -> 0)
		if tti.SGN != lastSGN {
			// check SGN is consecutive
			if tti.SGN != lastSGN+1 {
				warns = append(warns, validateErr(ErrSGNNotConsecutive, nil, true))
			}

			// check subtitle is the first of the group
			if tti.SN != 0 {
				warns = append(warns, validateErr(ErrNoFirstSubtitleInNewGroup, nil, true))
			}

			// closing EBN
			if tti.EBN != 0xFF {
				warns = append(warns, validateErr(ErrNonClosingEBNForLastSubtitle, nil, true))
			}

			// check CS is none or last
			if tti.CS != CumulativeStatusNone && tti.CS != CumulativeStatusLast {
				warns = append(warns, validateErr(ErrCSNotNoneOrLast, nil, true))
			}

			groups++
		}

		// Keep last values
		lastEBN = tti.EBN
		lastSN = tti.SN
		lastSGN = tti.SGN
		lastCS = tti.CS
	}

	// check if subtitle count matches
	if f.GSI.TNS != subtitles {
		warns = append(warns, validateErr(ErrSubtitleCountMismatch, nil, true))
	}

	// check if group count matches
	if f.GSI.TNG != groups {
		warns = append(warns, validateErr(ErrGroupCountMismatch, nil, true))
	}

	return warns, nil
}

var (
	ErrUnknown = errors.New("unknown error")

	ErrNoTTIBlocks            = errors.New("no TTI blocks")
	ErrTTIBlocksCountMismatch = errors.New("TTI blocks count mismatch")
	ErrFirstTCMismatch        = errors.New("first TTI timecode mismatch")

	ErrEBNNotConsecutive = errors.New("EBN not consecutive")
	ErrSNNotConsecutive  = errors.New("SN not consecutive")
	ErrSGNNotConsecutive = errors.New("SGN not consecutive")

	ErrNoFirstSubtitleInNewGroup = errors.New("no first subtitle in new group")

	ErrNonClosingEBNForLastSubtitle = errors.New("non closing EBN for last subtitle")

	ErrCSNotNoneOrFirst        = errors.New("CS not none or first")
	ErrCSNotIntermediateOrLast = errors.New("CS not intermediate or last")
	ErrCSNotNoneOrLast         = errors.New("CS not none or last")

	ErrSubtitleCountMismatch = errors.New("subtitle count mismatch")
	ErrGroupCountMismatch    = errors.New("group count mismatch")
)
