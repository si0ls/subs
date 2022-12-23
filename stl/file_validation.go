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
	gsiWarns, err := f.GSI.Validate()
	if err != nil {
		return warns, err
	}
	warns = appendNonNilErrs(warns, gsiWarns...)

	// having at least one TTI block
	if len(f.TTI) == 0 {
		return warns, validateErr(ErrNoTTIBlocks, nil, true)
	}

	// match between gsi.TNB and len(f.TTI)
	if f.GSI.TNB != len(f.TTI) {
		warns = append(warns, gsiErr(validateErr(ErrTTIBlocksCountMismatch, nil, true), GSIFieldTNB))
	}

	// check GSI TSF timecode and first TTI TCI timecode
	if f.GSI.TCF != f.TTI[0].TCI {
		warns = append(warns, gsiErr(validateErr(ErrTCFFirstTCIMismatch, nil, true), GSIFieldTCF))
	}

	var subtitles int
	var groups int

	var lastSN int = -1
	var lastSGN int = f.TTI[0].SGN
	var lastEBN int = 0xFF
	var lastCS CumulativeStatus = CumulativeStatusNone

	for i, tti := range f.TTI {
		// non nil TTI block
		if tti == nil {
			panic(fmt.Errorf("TTI block %d is nil", i))
		}

		// validate TTI block
		ttiWarns, err := tti.Validate(f.GSI.Framerate(), f.GSI.DSC, f.GSI.MNR)
		if err != nil {
			if ttiErr, ok := err.(*TTIError); ok {
				ttiErr.setBlockNumber(i)
			}
			return warns, err
		}
		setTTIErrsBlockNumber(ttiWarns, i)
		warns = appendNonNilErrs(warns, ttiWarns...)

		// same subtitle (same group)
		if tti.SN == lastSN && tti.SGN == lastSGN {
			// check EBN is consecutive
			if tti.EBN != lastEBN+1 {
				warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrEBNNotConsecutive, tti.EBN, false), TTIFieldEBN, i))
			}
		} else { // new subtitle (same group or not)
			subtitles++
		}

		// new subtitle, same group
		if tti.SN != lastSN && tti.SGN == lastSGN {
			// check SN is consecutive
			if tti.SN != lastSN+1 {
				warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrSNNotConsecutive, tti.SN, false), TTIFieldSN, i))
			}

			// closing EBN
			if tti.EBN != 0xFF {
				warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrNonClosingEBNForLastSubtitle, tti.EBN, false), TTIFieldEBN, i))
			}

			// check CS
			switch lastCS {
			case CumulativeStatusNone: // if last CS was None, then CS must be None or First
				if tti.CS != CumulativeStatusNone && tti.CS != CumulativeStatusFirst {
					warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrCSNotNoneOrFirst, tti.CS, false), TTIFieldCS, i))
				}
			case CumulativeStatusIntermediate: // if last was Intermediate, then CS must be Intermediate or Last
				if tti.CS != CumulativeStatusIntermediate && tti.CS != CumulativeStatusLast {
					warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrCSNotIntermediateOrLast, tti.CS, false), TTIFieldCS, i))
				}
			case CumulativeStatusFirst: // if last was First, then CS must be Intermediate or Last
				if tti.CS != CumulativeStatusIntermediate && tti.CS != CumulativeStatusLast {
					warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrCSNotIntermediateOrLast, tti.CS, false), TTIFieldCS, i))
				}
			case CumulativeStatusLast: // if last was Last, then CS must be None or Last
				if tti.CS != CumulativeStatusNone && tti.CS != CumulativeStatusLast {
					warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrCSNotNoneOrLast, tti.CS, false), TTIFieldCS, i))
				}
			}
		}

		// new group (new subtitle even if SN is the same -> 0)
		if tti.SGN != lastSGN {
			// check SGN is consecutive
			if tti.SGN != lastSGN+1 {
				warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrSGNNotConsecutive, tti.SGN, false), TTIFieldSGN, i))
			}

			// check subtitle is the first of the group
			if tti.SN != 0 {
				warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrNoFirstSubtitleInNewGroup, tti.SN, false), TTIFieldSN, i))
			}

			// closing EBN
			if tti.EBN != 0xFF {
				warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrNonClosingEBNForLastSubtitle, tti.EBN, false), TTIFieldEBN, i))
			}

			// check CS is none or last
			if tti.CS != CumulativeStatusNone && tti.CS != CumulativeStatusLast {
				warns = append(warns, ttiErrWithBlockNumber(validateErr(ErrCSNotNoneOrLast, tti.CS, false), TTIFieldCS, i))
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
		warns = append(warns, gsiErr(validateErr(ErrSubtitleCountMismatch, f.GSI.TNS, false), GSIFieldTNS))
	}

	// check if group count matches
	if f.GSI.TNG != groups {
		warns = append(warns, gsiErr(validateErr(ErrGroupCountMismatch, f.GSI.TNG, false), GSIFieldTNG))
	}

	return warns, nil
}

var (
	ErrUnknown = errors.New("unknown error")

	ErrNoTTIBlocks            = errors.New("no TTI blocks")
	ErrTTIBlocksCountMismatch = errors.New("TTI blocks count mismatch")
	ErrTCFFirstTCIMismatch    = errors.New("first TTI timecode mismatch")

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
