package stl

import "errors"

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
