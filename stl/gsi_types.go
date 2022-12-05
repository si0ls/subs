package stl

// CodePageNumber is the number of the code page used in the GSI block.
// Note: Other code pages may be used within a given national environment.
type CodePageNumber int

const (
	CodePageNumberInvalid        CodePageNumber = -1
	CodePageNumberUnitedStates   CodePageNumber = 437
	CodePageNumberMultiLingual   CodePageNumber = 850
	CodePageNumberPortugal       CodePageNumber = 860
	CodePageNumberCanadianFrench CodePageNumber = 863
	CodePageNumberNordic         CodePageNumber = 865
)

var cpnStringMap = map[CodePageNumber]string{
	CodePageNumberInvalid:        "<invalid>",
	CodePageNumberUnitedStates:   "United States",
	CodePageNumberMultiLingual:   "Multilingual",
	CodePageNumberPortugal:       "Portugal",
	CodePageNumberCanadianFrench: "Canadian/French",
	CodePageNumberNordic:         "Nordic",
}

// String returns the string representation of CodePageNumber.
func (cpn CodePageNumber) String() string {
	if s, ok := cpnStringMap[cpn]; ok {
		return s
	}
	return "Unknown"
}

// DiskFormatCode is the code used to represent the frame-rate of the data.
// Only "STL25.01" (25 fps) and "STL30.01" (30 fps) are supported values.
type DiskFormatCode string

const (
	DiskFormatCodeInvalid DiskFormatCode = "        "
	DiskFormatCode25_01   DiskFormatCode = "STL25.01"
	DiskFormatCode30_01   DiskFormatCode = "STL30.01"
)

// String returns the string representation of DiskFormatCode.
func (dfc DiskFormatCode) String() string {
	if dfc == DiskFormatCodeInvalid {
		return "<invalid>"
	} else if dfc != DiskFormatCode25_01 && dfc != DiskFormatCode30_01 {
		return "Unknown"
	}
	return string(dfc)
}

// DisplayStandardCode is the code used to represent the display mode.
// Only "Undefined" (blank), "Open Subtitling" (0), "Level-1 Teletext" (1) and "Level-2 Teletext" (2) are supported values.
type DisplayStandardCode byte

const (
	DisplayStandardCodeBlank          DisplayStandardCode = 0xFF
	DisplayStandardCodeOpenSubtitling DisplayStandardCode = 0x00
	DisplayStandardCodeLevel1Teletext DisplayStandardCode = 0x01
	DisplayStandardCodeLevel2Teletext DisplayStandardCode = 0x02
)

var dscStringMap = map[DisplayStandardCode]string{
	DisplayStandardCodeBlank:          "Blank",
	DisplayStandardCodeOpenSubtitling: "Open Subtitling",
	DisplayStandardCodeLevel1Teletext: "Level-1 Teletext",
	DisplayStandardCodeLevel2Teletext: "Level-2 Teletext",
}

// String returns the string representation of DisplayStandardCode.
func (dsc DisplayStandardCode) String() string {
	if s, ok := dscStringMap[dsc]; ok {
		return s
	}
	return "Unknown"
}

// CharacterCodeTable is the code used to represent the character code table used to define the text in the Text Field (TF) of the TTI blocks.
// Only "Latin" (0), "Latin/Cyrillic" (1), "Latin/Arabic" (2), "Latin/Greek" (3) and "Latin/Hebrew" (4) are supported values.
type CharacterCodeTable byte

const (
	CharacterCodeTableInvalid       CharacterCodeTable = 0xFF
	CharacterCodeTableLatin         CharacterCodeTable = 0x00
	CharacterCodeTableLatinCyrillic CharacterCodeTable = 0x01
	CharacterCodeTableLatinArabic   CharacterCodeTable = 0x02
	CharacterCodeTableLatinGreek    CharacterCodeTable = 0x03
	CharacterCodeTableLatinHebrew   CharacterCodeTable = 0x04
)

var cctStringMap = map[CharacterCodeTable]string{
	CharacterCodeTableInvalid:       "<invalid>",
	CharacterCodeTableLatin:         "Latin",
	CharacterCodeTableLatinCyrillic: "Latin/Cyrillic",
	CharacterCodeTableLatinArabic:   "Latin/Arabic",
	CharacterCodeTableLatinGreek:    "Latin/Greek",
	CharacterCodeTableLatinHebrew:   "Latin/Hebrew",
}

// String returns the string representation of CharacterCodeTable.
func (cct CharacterCodeTable) String() string {
	if s, ok := cctStringMap[cct]; ok {
		return s
	}
	return "Unknown"
}

// Language Code is the code of the language for wih the subtitle list is prepared.
type LanguageCode byte

const (
	LanguageCodeInvalid       LanguageCode = 0xFF
	LanguageCodeUnknown       LanguageCode = 0x00
	LanguageCodeAlbanian      LanguageCode = 0x01
	LanguageCodeBreton        LanguageCode = 0x02
	LanguageCodeCatalan       LanguageCode = 0x03
	LanguageCodeCroatian      LanguageCode = 0x04
	LanguageCodeWelsh         LanguageCode = 0x05
	LanguageCodeCzech         LanguageCode = 0x06
	LanguageCodeDanish        LanguageCode = 0x07
	LanguageCodeGerman        LanguageCode = 0x08
	LanguageCodeEnglish       LanguageCode = 0x09
	LanguageCodeSpanish       LanguageCode = 0x0A
	LanguageCodeEsperanto     LanguageCode = 0x0B
	LanguageCodeEstonian      LanguageCode = 0x0C
	LanguageCodeBasque        LanguageCode = 0x0D
	LanguageCodeFaroese       LanguageCode = 0x0E
	LanguageCodeFrench        LanguageCode = 0x0F
	LanguageCodeFrisian       LanguageCode = 0x10
	LanguageCodeIrish         LanguageCode = 0x11
	LanguageCodeGaelic        LanguageCode = 0x12
	LanguageCodeGalician      LanguageCode = 0x13
	LanguageCodeIcelandic     LanguageCode = 0x14
	LanguageCodeItalian       LanguageCode = 0x15
	LanguageCodeLappish       LanguageCode = 0x16
	LanguageCodeLatin         LanguageCode = 0x17
	LanguageCodeLatvian       LanguageCode = 0x18
	LanguageCodeLuxembourgian LanguageCode = 0x19
	LanguageCodeLithuanian    LanguageCode = 0x1A
	LanguageCodeHungarian     LanguageCode = 0x1B
	LanguageCodeMaltese       LanguageCode = 0x1C
	LanguageCodeDutch         LanguageCode = 0x1D
	LanguageCodeNorwegian     LanguageCode = 0x1E
	LanguageCodeOccitan       LanguageCode = 0x1F
	LanguageCodePolish        LanguageCode = 0x20
	LanguageCodePortugese     LanguageCode = 0x21
	LanguageCodeRomanian      LanguageCode = 0x22
	LanguageCodeRomansh       LanguageCode = 0x23
	LanguageCodeSerbian       LanguageCode = 0x24
	LanguageCodeSlovak        LanguageCode = 0x25
	LanguageCodeSlovenian     LanguageCode = 0x26
	LanguageCodeFinnish       LanguageCode = 0x27
	LanguageCodeSwedish       LanguageCode = 0x28
	LanguageCodeTurkish       LanguageCode = 0x29
	LanguageCodeFlemish       LanguageCode = 0x2A
	LanguageCodeWallon        LanguageCode = 0x2B
	LanguageCodeAmharic       LanguageCode = 0x7F
	LanguageCodeArabic        LanguageCode = 0x7E
	LanguageCodeArmenian      LanguageCode = 0x7D
	LanguageCodeAssamese      LanguageCode = 0x7C
	LanguageCodeAzerbaijani   LanguageCode = 0x7B
	LanguageCodeBambora       LanguageCode = 0x7A
	LanguageCodeBielorussian  LanguageCode = 0x79
	LanguageCodeBengali       LanguageCode = 0x78
	LanguageCodeBulgarian     LanguageCode = 0x77
	LanguageCodeBurmese       LanguageCode = 0x76
	LanguageCodeChinese       LanguageCode = 0x75
	LanguageCodeChurash       LanguageCode = 0x74
	LanguageCodeDari          LanguageCode = 0x73
	LanguageCodeFulani        LanguageCode = 0x72
	LanguageCodeGeorgian      LanguageCode = 0x71
	LanguageCodeGreek         LanguageCode = 0x70
	LanguageCodeGujurati      LanguageCode = 0x6F
	LanguageCodeGurani        LanguageCode = 0x6E
	LanguageCodeHausa         LanguageCode = 0x6D
	LanguageCodeHebrew        LanguageCode = 0x6C
	LanguageCodeHindi         LanguageCode = 0x6B
	LanguageCodeIndonesian    LanguageCode = 0x6A
	LanguageCodeJapanese      LanguageCode = 0x69
	LanguageCodeKannada       LanguageCode = 0x68
	LanguageCodeKazakh        LanguageCode = 0x67
	LanguageCodeKhmer         LanguageCode = 0x66
	LanguageCodeKorean        LanguageCode = 0x65
	LanguageCodeLaotian       LanguageCode = 0x64
	LanguageCodeMacedonian    LanguageCode = 0x63
	LanguageCodeMalagasay     LanguageCode = 0x62
	LanguageCodeMalaysian     LanguageCode = 0x61
	LanguageCodeMoldavian     LanguageCode = 0x60
	LanguageCodeMarathi       LanguageCode = 0x5F
	LanguageCodeNdebele       LanguageCode = 0x5E
	LanguageCodeNepali        LanguageCode = 0x5D
	LanguageCodeOriya         LanguageCode = 0x5C
	LanguageCodePapamiento    LanguageCode = 0x5B
	LanguageCodePersian       LanguageCode = 0x5A
	LanguageCodePunjabi       LanguageCode = 0x59
	LanguageCodePushtu        LanguageCode = 0x58
	LanguageCodeQuechua       LanguageCode = 0x57
	LanguageCodeRussian       LanguageCode = 0x56
	LanguageCodeRuthenian     LanguageCode = 0x55
	LanguageCodeSerboCroat    LanguageCode = 0x54
	LanguageCodeShona         LanguageCode = 0x53
	LanguageCodeSinhalese     LanguageCode = 0x52
	LanguageCodeSomali        LanguageCode = 0x51
	LanguageCodeSrananTongo   LanguageCode = 0x50
	LanguageCodeSwahili       LanguageCode = 0x4F
	LanguageCodeTadzhik       LanguageCode = 0x4E
	LanguageCodeTamil         LanguageCode = 0x4D
	LanguageCodeTatar         LanguageCode = 0x4C
	LanguageCodeTelugu        LanguageCode = 0x4B
	LanguageCodeThai          LanguageCode = 0x4A
	LanguageCodeUkrainian     LanguageCode = 0x49
	LanguageCodeUrdu          LanguageCode = 0x48
	LanguageCodeUzbek         LanguageCode = 0x47
	LanguageCodeVietnamese    LanguageCode = 0x46
	LanguageCodeZulu          LanguageCode = 0x45
)

var lcStringMap = map[LanguageCode]string{
	LanguageCodeInvalid:       "<invalid>",
	LanguageCodeUnknown:       "Unknown/not applicable",
	LanguageCodeAlbanian:      "Albanian",
	LanguageCodeBreton:        "Breton",
	LanguageCodeCatalan:       "Catalan",
	LanguageCodeCroatian:      "Croatian",
	LanguageCodeWelsh:         "Welsh",
	LanguageCodeCzech:         "Czech",
	LanguageCodeDanish:        "Danish",
	LanguageCodeGerman:        "German",
	LanguageCodeEnglish:       "English",
	LanguageCodeSpanish:       "Spanish",
	LanguageCodeEsperanto:     "Esperanto",
	LanguageCodeEstonian:      "Estonian",
	LanguageCodeBasque:        "Basque",
	LanguageCodeFaroese:       "Faroese",
	LanguageCodeFrench:        "French",
	LanguageCodeFrisian:       "Frisian",
	LanguageCodeIrish:         "Irish",
	LanguageCodeGaelic:        "Gaelic",
	LanguageCodeGalician:      "Galician",
	LanguageCodeIcelandic:     "Icelandic",
	LanguageCodeItalian:       "Italian",
	LanguageCodeLappish:       "Lappish",
	LanguageCodeLatin:         "Latin",
	LanguageCodeLatvian:       "Latvian",
	LanguageCodeLuxembourgian: "Luxembourgian",
	LanguageCodeLithuanian:    "Lithuanian",
	LanguageCodeHungarian:     "Hungarian",
	LanguageCodeMaltese:       "Maltese",
	LanguageCodeDutch:         "Dutch",
	LanguageCodeNorwegian:     "Norwegian",
	LanguageCodeOccitan:       "Occitan",
	LanguageCodePolish:        "Polish",
	LanguageCodePortugese:     "Portugese",
	LanguageCodeRomanian:      "Romanian",
	LanguageCodeRomansh:       "Romansh",
	LanguageCodeSerbian:       "Serbian",
	LanguageCodeSlovak:        "Slovak",
	LanguageCodeSlovenian:     "Slovenian",
	LanguageCodeFinnish:       "Finnish",
	LanguageCodeSwedish:       "Swedish",
	LanguageCodeTurkish:       "Turkish",
	LanguageCodeFlemish:       "Flemish",
	LanguageCodeWallon:        "Wallon",
	LanguageCodeAmharic:       "Amharic",
	LanguageCodeArabic:        "Arabic",
	LanguageCodeArmenian:      "Armenian",
	LanguageCodeAssamese:      "Assamese",
	LanguageCodeAzerbaijani:   "Azerbaijani",
	LanguageCodeBambora:       "Bambora",
	LanguageCodeBielorussian:  "Bielorussian",
	LanguageCodeBengali:       "Bengali",
	LanguageCodeBulgarian:     "Bulgarian",
	LanguageCodeBurmese:       "Burmese",
	LanguageCodeChinese:       "Chinese",
	LanguageCodeChurash:       "Churash",
	LanguageCodeDari:          "Dari",
	LanguageCodeFulani:        "Fulani",
	LanguageCodeGeorgian:      "Georgian",
	LanguageCodeGreek:         "Greek",
	LanguageCodeGujurati:      "Gujurati",
	LanguageCodeGurani:        "Gurani",
	LanguageCodeHausa:         "Hausa",
	LanguageCodeHebrew:        "Hebrew",
	LanguageCodeHindi:         "Hindi",
	LanguageCodeIndonesian:    "Indonesian",
	LanguageCodeJapanese:      "Japanese",
	LanguageCodeKannada:       "Kannada",
	LanguageCodeKazakh:        "Kazakh",
	LanguageCodeKhmer:         "Khmer",
	LanguageCodeKorean:        "Korean",
	LanguageCodeLaotian:       "Laotian",
	LanguageCodeMacedonian:    "Macedonian",
	LanguageCodeMalagasay:     "Malagasay",
	LanguageCodeMalaysian:     "Malaysian",
	LanguageCodeMoldavian:     "Moldavian",
	LanguageCodeMarathi:       "Marathi",
	LanguageCodeNdebele:       "Ndebele",
	LanguageCodeNepali:        "Nepali",
	LanguageCodeOriya:         "Oriya",
	LanguageCodePapamiento:    "Papamiento",
	LanguageCodePersian:       "Persian",
	LanguageCodePunjabi:       "Punjabi",
	LanguageCodePushtu:        "Pushtu",
	LanguageCodeQuechua:       "Quechua",
	LanguageCodeRussian:       "Russian",
	LanguageCodeRuthenian:     "Ruthenian",
	LanguageCodeSerboCroat:    "Serbo-croat",
	LanguageCodeShona:         "Shona",
	LanguageCodeSinhalese:     "Sinhalese",
	LanguageCodeSomali:        "Somali",
	LanguageCodeSrananTongo:   "Sranan Tongo",
	LanguageCodeSwahili:       "Swahili",
	LanguageCodeTadzhik:       "Tadzhik",
	LanguageCodeTamil:         "Tamil",
	LanguageCodeTatar:         "Tatar",
	LanguageCodeTelugu:        "Telugu",
	LanguageCodeThai:          "Thai",
	LanguageCodeUkrainian:     "Ukrainian",
	LanguageCodeUrdu:          "Urdu",
	LanguageCodeUzbek:         "Uzbek",
	LanguageCodeVietnamese:    "Vietnamese",
	LanguageCodeZulu:          "Zulu",
}

// String returns the string representation of LanguageCode.
func (lc LanguageCode) String() string {
	if s, ok := lcStringMap[lc]; ok {
		return s
	}
	return "Unknown"
}

// TimeCodeStatus is indicating the validity of the information given in the GSI and TTO blocks containing time-code data.
// Only "Not intended for use" (0) and "Intended for use" (1) are supported values.
type TimeCodeStatus byte

const (
	TimeCodeStatusInvalid           TimeCodeStatus = 0xFF
	TimeCodeStatusNotIntendedForUse TimeCodeStatus = 0x00
	TimeCodeStatusIntendedForUse    TimeCodeStatus = 0x01
)

var tcsStringMap = map[TimeCodeStatus]string{
	TimeCodeStatusInvalid:           "<invalid>",
	TimeCodeStatusNotIntendedForUse: "Not intended for use",
	TimeCodeStatusIntendedForUse:    "Intended for use",
}

// String returns the string representation of TimeCodeStatus.
func (tcs TimeCodeStatus) String() string {
	if s, ok := tcsStringMap[tcs]; ok {
		return s
	}
	return "Unknown"
}
