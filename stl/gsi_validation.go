package stl

import (
	"errors"
	"fmt"
)

// Validate validates GSI block.
// It returns a slice of warnings and an error if any.
// Warnings are returned for each field that is invalid, warnings can flagged
// as fatal if they are considered to be fatal to further file processing.
// An error is returned if a field is invalid and prevents validation of
// further fields.
func (gsi *GSIBlock) Validate() ([]error, error) {
	var errs []error

	// CPN - in list -> fatal
	errs = appendNonNilErrs(errs, gsiErr(validateList(gsi.CPN, cpnValidValues, ErrUnsupportedCPN, true), GSIFieldCPN))

	// DFC - in list -> fatal
	errs = appendNonNilErrs(errs, gsiErr(validateList(gsi.DFC, dfcValidValues, ErrUnsupportedDFC, true), GSIFieldDFC))

	// Framerate -> raise error
	if gsi.Framerate() != 25 && gsi.Framerate() != 30 {
		return errs, validateErr(fmt.Errorf("%w: must be 25 or 30, prevents further validation", ErrUnsupportedFramerate), gsi.Framerate(), true)
	}

	// DSC - in list
	errs = appendNonNilErrs(errs, gsiErr(validateList(gsi.DSC, dscValidValues, ErrUnsupportedDSC, false), GSIFieldDSC))

	// CCT - in list -> fatal
	errs = appendNonNilErrs(errs, gsiErr(validateList(gsi.CCT, cctValidValues, ErrUnsupportedCCT, true), GSIFieldCCT))

	// LC - in list
	// Trick, do not validate list to avoid enormous error message
	if gsi.LC > LanguageCodeWallon &&
		(gsi.LC < LanguageCodeZulu || gsi.LC > LanguageCodeAmharic) {
		errs = appendNonNilErrs(errs, gsiErr(validateErr(fmt.Errorf("%w: must be one of the supported MAC/packet family language code", ErrUnsupportedLC), gsi.LC, false), GSIFieldLC))
	}

	// OPT - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.OPT, ErrEmptyOPT, false), GSIFieldOPT))

	// OET - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.OET, ErrEmptyOET, false), GSIFieldOET))

	// TPT - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.TPT, ErrEmptyTPT, false), GSIFieldTPT))

	// TET - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.TET, ErrEmptyTET, false), GSIFieldTET))

	// TN - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.TN, ErrEmptyTN, false), GSIFieldTN))

	// TCD - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.TCD, ErrEmptyTCD, false), GSIFieldTCD))

	// SLR - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.SLR, ErrEmptySLR, false), GSIFieldSLR))

	// CD - valid dat
	errs = appendNonNilErrs(errs, gsiErr(validateDate(gsi.CD, ErrEmptyCD, false), GSIFieldCD))

	// RD - valid date
	errs = appendNonNilErrs(errs, gsiErr(validateDate(gsi.RD, ErrEmptyRD, false), GSIFieldRD))

	// Dates (CD, RD) - CD <= RD
	errs = appendNonNilErrs(errs, gsiErr(validateDateOrder(gsi.CD, gsi.RD, ErrCDGreaterThanRD, false), GSIFieldCD))

	// RN - between 0 and 99
	errs = appendNonNilErrs(errs, gsiErr(validateRange(gsi.RN, 0, 99, ErrUnsupportedRN, false), GSIFieldRN))

	// TNB - between 0 and 99999
	errs = appendNonNilErrs(errs, gsiErr(validateRange(gsi.TNB, 0, 99999, ErrUnsupportedTNB, false), GSIFieldTNB))

	// TNS - between 0 and 99999
	errs = appendNonNilErrs(errs, gsiErr(validateRange(gsi.TNS, 0, 99999, ErrUnsupportedTNS, false), GSIFieldTNS))

	// TNG - between 0 and 255
	errs = appendNonNilErrs(errs, gsiErr(validateRange(gsi.TNG, 0, 255, ErrUnsupportedTNG, false), GSIFieldTNG))

	// MNC - between 0 and 99
	errs = appendNonNilErrs(errs, gsiErr(validateRange(gsi.MNC, 0, 99, ErrUnsupportedMNC, false), GSIFieldMNC))

	// MNR - between 0 and 99, between 1 and 23 or null (0) if teletext (based on DSC)
	if gsi.DSC == DisplayStandardCodeLevel1Teletext || gsi.DSC == DisplayStandardCodeLevel2Teletext {
		errs = appendNonNilErrs(errs, gsiErr(validateRange(gsi.MNR, 0, 23, ErrUnsupportedMNRTeletext, false), GSIFieldMNR))
	} else if gsi.DSC == DisplayStandardCodeOpenSubtitling {
		errs = appendNonNilErrs(errs, gsiErr(validateRange(gsi.MNR, 0, 99, ErrUnsupportedMNROpenSubtitling, false), GSIFieldMNR))
	}

	// TCS - in list
	errs = appendNonNilErrs(errs, gsiErr(validateList(gsi.TCS, tcsValidValues, ErrUnsupportedTCS, false), GSIFieldTCS))

	// TCP - valid timecode
	errs = appendNonNilErrs(errs, gsiErr(validateTimecode(gsi.TCP, gsi.Framerate(), ErrEmptyTCP, false), GSIFieldTCP))

	// TCF - valid timecode
	errs = appendNonNilErrs(errs, gsiErr(validateTimecode(gsi.TCF, gsi.Framerate(), ErrEmptyTCF, false), GSIFieldTCF))

	// Timecodes (CTP, TCF) - CTP <= TCF
	errs = appendNonNilErrs(errs, gsiErr(validateTimecodeOrder(gsi.TCP, gsi.TCF, gsi.Framerate(), ErrTCPTCFOrder, false), GSIFieldTCP))

	// TND - between 0 and 9
	errs = appendNonNilErrs(errs, gsiErr(validateRange(gsi.TND, 0, 9, ErrUnsupportedTND, false), GSIFieldTND))

	// DSN - between 1 and TND
	errs = appendNonNilErrs(errs, gsiErr(validateRange(gsi.DSN, 1, gsi.TND, ErrUnsupportedDSN, false), GSIFieldDSN))

	// CO - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.CO, ErrEmptyCO, false), GSIFieldCO))

	// PUB - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.PUB, ErrEmptyPUB, false), GSIFieldPUB))

	// EN - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.EN, ErrEmptyEN, false), GSIFieldEN))

	// ECD - not empty
	errs = appendNonNilErrs(errs, gsiErr(validateNonEmptyString(gsi.ECD, ErrEmptyECD, false), GSIFieldECD))

	return errs, nil
}

var (
	ErrUnsupportedCPN               = errors.New("unsupported CPN")
	ErrUnsupportedDFC               = errors.New("unsupported DFC")
	ErrUnsupportedFramerate         = errors.New("unsupported framerate")
	ErrUnsupportedDSC               = errors.New("unsupported DSC")
	ErrUnsupportedCCT               = errors.New("unsupported CCT")
	ErrUnsupportedLC                = errors.New("unsupported LC")
	ErrEmptyOPT                     = errors.New("empty OPT")
	ErrEmptyOET                     = errors.New("empty OET")
	ErrEmptyTPT                     = errors.New("empty TPT")
	ErrEmptyTET                     = errors.New("empty TET")
	ErrEmptyTN                      = errors.New("empty TN")
	ErrEmptyTCD                     = errors.New("empty TCD")
	ErrEmptySLR                     = errors.New("empty SLR")
	ErrEmptyCD                      = errors.New("empty CD")
	ErrEmptyCR                      = errors.New("empty CR")
	ErrEmptyRD                      = errors.New("empty RD")
	ErrCDGreaterThanRD              = errors.New("CD greater than RD")
	ErrUnsupportedRN                = errors.New("unsupported RN")
	ErrUnsupportedTNB               = errors.New("unsupported TNB")
	ErrUnsupportedTNS               = errors.New("unsupported TNS")
	ErrUnsupportedTNG               = errors.New("unsupported TNG")
	ErrUnsupportedMNC               = errors.New("unsupported MNC")
	ErrUnsupportedMNRTeletext       = errors.New("unsupported MNR (teletext)")
	ErrUnsupportedMNROpenSubtitling = errors.New("unsupported MNR (open subtitling)")
	ErrUnsupportedTCS               = errors.New("unsupported TCS")
	ErrEmptyTCP                     = errors.New("empty TCP")
	ErrEmptyTCF                     = errors.New("empty TCF")
	ErrTCPTCFOrder                  = errors.New("TCP greater than TCF")
	ErrInvalidTimecodes             = errors.New("invalid timecodes")
	ErrUnsupportedTND               = errors.New("unsupported TND")
	ErrUnsupportedDSN               = errors.New("unsupported DSN")
	ErrEmptyCO                      = errors.New("empty CO")
	ErrEmptyPUB                     = errors.New("empty PUB")
	ErrEmptyEN                      = errors.New("empty EN")
	ErrEmptyECD                     = errors.New("empty ECD")
)

var cpnValidValues = []CodePageNumber{
	CodePageNumberUnitedStates,
	CodePageNumberMultiLingual,
	CodePageNumberPortugal,
	CodePageNumberCanadianFrench,
	CodePageNumberNordic,
}

var dfcValidValues = []DiskFormatCode{
	DiskFormatCode25_01,
	DiskFormatCode30_01,
}

var dscValidValues = []DisplayStandardCode{
	DisplayStandardCodeBlank,
	DisplayStandardCodeOpenSubtitling,
	DisplayStandardCodeLevel1Teletext,
	DisplayStandardCodeLevel2Teletext,
}

var cctValidValues = []CharacterCodeTable{
	CharacterCodeTableLatin,
	CharacterCodeTableLatinCyrillic,
	CharacterCodeTableLatinArabic,
	CharacterCodeTableLatinGreek,
	CharacterCodeTableLatinHebrew,
}

var tcsValidValues = []TimeCodeStatus{
	TimeCodeStatusNotIntendedForUse,
	TimeCodeStatusIntendedForUse,
}
