package stl

// CumulativeStatus is indicating that a subtitle is part of a cumulative set of subtitles.
type CumulativeStatus byte

const (
	CumulativeStatusInvalid      CumulativeStatus = 0xFF
	CumulativeStatusNone         CumulativeStatus = 0x00
	CumulativeStatusFirst        CumulativeStatus = 0x01
	CumulativeStatusIntermediate CumulativeStatus = 0x02
	CumulativeStatusLast         CumulativeStatus = 0x03
)

var csStringMap = map[CumulativeStatus]string{
	CumulativeStatusInvalid:      "<invalid>",
	CumulativeStatusNone:         "None",
	CumulativeStatusFirst:        "First",
	CumulativeStatusIntermediate: "Intermediate",
	CumulativeStatusLast:         "Last",
}

// String returns the string representation of CumulativeStatus.
func (cs CumulativeStatus) String() string {
	if s, ok := csStringMap[cs]; ok {
		return s
	}
	return "Unknown"
}

// JustificationCode is indicating the horizontal alignment of the displayed subtitle.
// Only "unchanged presentation" (0), "left-justified text" (1), "centered text" (2) and "right-justified text" (3) are supported.
type JustificationCode byte

const (
	JustificationCodeInvalid               JustificationCode = 0xFF
	JustificationCodeUnchangedPresentation JustificationCode = 0x00
	JustificationCodeLeftJustifiedText     JustificationCode = 0x01
	JustificationCodeCenteredText          JustificationCode = 0x02
	JustificationCodeRightJustifiedText    JustificationCode = 0x03
)

var jcStringMap = map[JustificationCode]string{
	JustificationCodeInvalid:               "<invalid>",
	JustificationCodeUnchangedPresentation: "Unchanged presentation",
	JustificationCodeLeftJustifiedText:     "Left-justified text",
	JustificationCodeCenteredText:          "Centered text",
	JustificationCodeRightJustifiedText:    "Right-justified text",
}

// String returns the string representation of JustificationCode.
func (jc JustificationCode) String() string {
	if s, ok := jcStringMap[jc]; ok {
		return s
	}
	return "Unknown"
}

// CommentFlag is used to indicate TTI blocks which contains text as translator's comments instead of subtitle data.
// Only "subtitle data" (0) and "translator's comments" (1) are supported.
type CommentFlag byte

const (
	CommentFlagInvalid            CommentFlag = 0xFF
	CommentFlagSubtitleData       CommentFlag = 0x00
	CommentFlagTranslatorComments CommentFlag = 0x01
)

var cfStringMap = map[CommentFlag]string{
	CommentFlagInvalid:            "<invalid>",
	CommentFlagSubtitleData:       "Subtitle data",
	CommentFlagTranslatorComments: "Translator's comments",
}

// String returns the string representation of CommentFlag.
func (cf CommentFlag) String() string {
	if s, ok := cfStringMap[cf]; ok {
		return s
	}
	return "Unknown"
}
