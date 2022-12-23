package stl

import (
	"time"
)

// GSIBlockSize is the size in bytes of the GSI block in a STL file.
const GSIBlockSize = 1024

// GSIBlock is the General Subtitle Information (GSI) block representation.
type GSIBlock struct {
	CPN CodePageNumber      // Code Page Number
	DFC DiskFormatCode      // Disk Format Code
	DSC DisplayStandardCode // Display Standard Code
	CCT CharacterCodeTable  // Character Code Table number
	LC  LanguageCode        // Language Code
	OPT string              // Original Program Title
	OET string              // Original Episode Title
	TPT string              // Translated Program Title
	TET string              // Translated Episode Title
	TN  string              // Translator's Name
	TCD string              // Translator's Contact Details
	SLR string              // Subtitle List Reference Code
	CD  time.Time           // Creation Date
	RD  time.Time           // Revision Date
	RN  int                 // Revision Number
	TNB int                 // Total Number of Text and Timing Information (TTI) blocks
	TNS int                 // Total Number of Subtitles
	TNG int                 // Total Number of Subtitle Groups
	MNC int                 // Maximum Number of Displayable Characters in any text row
	MNR int                 // Maximum Number of Displayable Rows
	TCS TimeCodeStatus      // Time Code: Status
	TCP Timecode            // Time Code: Start-of-Program
	TCF Timecode            // Time Code: First In-Cue
	TND int                 // Total Number of Disks
	DSN int                 // Disk Sequence Number
	CO  string              // Country of Origin
	PUB string              // Publisher
	EN  string              // Editor's Name
	ECD string              // Editor's Contact
	UDA []byte              // User-Defined Area
}

// NewGSIBlock returns a new GSI block.
func NewGSIBlock() *GSIBlock {
	gsi := GSIBlock{}
	gsi.Reset()
	return &gsi
}

// Framerate returns the framerate of the GSI block (extracted from the Disk Format Code).
// The supported values are 25 and 30 fps.
// Returns -1 if Disk Format Code is unsupported.
func (gsi *GSIBlock) Framerate() uint {
	switch gsi.DFC {
	case DiskFormatCode25_01:
		return 25
	case DiskFormatCode30_01:
		return 30
	}
	return 0
}

// Reset resets the GSI block to its default values.
func (gsi *GSIBlock) Reset() {
	gsi.CPN = CodePageNumberInvalid
	gsi.DFC = DiskFormatCodeInvalid
	gsi.DSC = DisplayStandardCodeBlank
	gsi.CCT = CharacterCodeTableInvalid
	gsi.LC = LanguageCodeInvalid
	gsi.OPT = ""
	gsi.OET = ""
	gsi.TPT = ""
	gsi.TET = ""
	gsi.TN = ""
	gsi.TCD = ""
	gsi.SLR = ""
	gsi.CD = time.Time{}
	gsi.RD = time.Time{}
	gsi.RN = -1
	gsi.TNB = -1
	gsi.TNS = -1
	gsi.TNG = -1
	gsi.MNC = -1
	gsi.MNR = -1
	gsi.TCS = TimeCodeStatusInvalid
	gsi.TCF = Timecode{
		Hours:   -1,
		Minutes: -1,
		Seconds: -1,
		Frames:  -1,
	}
	gsi.TCP = Timecode{
		Hours:   -1,
		Minutes: -1,
		Seconds: -1,
		Frames:  -1,
	}
	gsi.TND = -1
	gsi.DSN = -1
	gsi.CO = ""
	gsi.PUB = ""
	gsi.EN = ""
	gsi.ECD = ""
	gsi.UDA = []byte{}
}
