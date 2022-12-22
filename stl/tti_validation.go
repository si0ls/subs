package stl

import "errors"

var (
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
