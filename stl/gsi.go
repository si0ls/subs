package stl

import (
	"io"
	"time"
)

// GSIBlockSize is the size in bytes of the GSI block in a STL file.
const GSIBlockSize = 1024

// GSIBlock is the General Subtitle Information (GSI) block representation.
type GSIBlock struct {
	CPN CodePageNumber      // Code Page Number - bytes 0..2 (3 bytes)
	DFC DiskFormatCode      // Disk Format Code - bytes 3..10 (8 bytes)
	DSC DisplayStandardCode // Display Standard Code - byte 11 (1 byte)
	CCT CharacterCodeTable  // Character Code Table number - bytes 12..13 (2 bytes)
	LC  LanguageCode        // Language Code - bytes 14..15 (2 bytes)
	OPT string              // Original Program Title - bytes 16..47 (32 bytes)
	OET string              // Original Episode Title - bytes 48..79 (32 bytes)
	TPT string              // Translated Program Title - bytes 80..111 (32 bytes)
	TET string              // Translated Episode Title - bytes 112..143 (32 bytes)
	TN  string              // Translator's Name - bytes 144..175 (32 bytes)
	TCD string              // Translator's Contact Details - bytes 176..207 (32 bytes)
	SLR string              // Subtitle List Reference Code - bytes 208..223 (16 bytes)
	CD  time.Time           // Creation Date - bytes 224..229 (6 bytes)
	RD  time.Time           // Revision Date - bytes 230..235 (6 bytes)
	RN  int                 // Revision Number - bytes 236..237 (2 bytes)
	TNB int                 // Total Number of Text and Timing Information (TTI) blocks - bytes 238..242 (5 bytes)
	TNS int                 // Total Number of Subtitles - bytes 243..247 (5 bytes)
	TNG int                 // Total Number of Subtitle Groups - bytes 248..250 (3 bytes)
	MNC int                 // Maximum Number of Displayable Characters in any text row - bytes 251..252 (2 bytes)
	MNR int                 // Maximum Number of Displayable Rows - bytes 253..254 (2 bytes)
	TCS TimeCodeStatus      // Time Code: Status - bytes 255 (1 byte)
	TCP Timecode            // Time Code: Start-of-Program - bytes 256..263 (8 bytes)
	TCF Timecode            // Time Code: First In-Cue - bytes 264..271 (8 bytes)
	TND int                 // Total Number of Disks - byte 272 (1 byte)
	DSN int                 // Disk Sequence Number - byte 273 (1 byte)
	CO  string              // Country of Origin - bytes 274..276 (3 bytes)
	PUB string              // Publisher - bytes 277..308 (32 bytes)
	EN  string              // Editor's Name - bytes 309..340 (32 bytes)
	ECD string              // Editor's Contact Details - bytes 341..372 (32 bytes)
	UDA []byte              // User-Defined Area - bytes 448..1023 (576 bytes)
}

// NewGSIBlock returns a new GSI block.
func NewGSIBlock() *GSIBlock {
	gsi := GSIBlock{}
	gsi.reset()
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

// Decode reads and decodes GSI block from reader.
func (gsi *GSIBlock) Decode(r io.Reader) ([]error, error) {
	b := make([]byte, GSIBlockSize)
	if _, err := io.ReadFull(r, b); err != nil {
		return nil, err
	}

	gsi.reset()

	var errs []error

	appendErrs(errs, decodeGSIInt(b[0:3], (*int)(&gsi.CPN)))                 // Code Page Number - bytes 0..2 (3 bytes)
	appendErrs(errs, decodeGSIString(b[3:11], (*string)(&gsi.DFC), gsi.CPN)) // Disk Format Code - bytes 3..10 (8 bytes)
	appendErrs(errs, decodeGSIByte(b[11:12], (*byte)(&gsi.DSC)))             // Display Standard Code - byte 11 (1 byte)
	appendErrs(errs, decodeGSIByte(b[12:14], (*byte)(&gsi.CCT)))             // Character Code Table number - bytes 12..13 (2 bytes)
	appendErrs(errs, decodeGSIHex(b[14:16], (*byte)(&gsi.LC)))               // Language Code - bytes 14..15 (2 bytes)
	appendErrs(errs, decodeGSIString(b[16:48], &gsi.OPT, gsi.CPN))           // Original Program Title - bytes 16..47 (32 bytes)
	appendErrs(errs, decodeGSIString(b[48:80], &gsi.OET, gsi.CPN))           // Original Episode Title - bytes 48..79 (32 bytes)
	appendErrs(errs, decodeGSIString(b[80:112], &gsi.TPT, gsi.CPN))          // Translated Program Title - bytes 80..111 (32 bytes)
	appendErrs(errs, decodeGSIString(b[112:144], &gsi.TET, gsi.CPN))         // Translated Episode Title - bytes 112..143 (32 bytes)
	appendErrs(errs, decodeGSIString(b[144:176], &gsi.TN, gsi.CPN))          // Translator's Name - bytes 144..175 (32 bytes)
	appendErrs(errs, decodeGSIString(b[176:208], &gsi.TCD, gsi.CPN))         // Translator's Contact Details - bytes 176..207 (32 bytes)
	appendErrs(errs, decodeGSIString(b[208:224], &gsi.SLR, gsi.CPN))         // Subtitle List Reference Code - bytes 208..223 (16 bytes)
	appendErrs(errs, decodeGSIDate(b[224:230], &gsi.CD))                     // Creation Date - bytes 224..229 (6 bytes)
	appendErrs(errs, decodeGSIDate(b[230:236], &gsi.RD))                     // Revision Date - bytes 230..235 (6 bytes)
	appendErrs(errs, decodeGSIInt(b[236:238], &gsi.RN))                      // Revision Number - bytes 236..237 (2 bytes)
	appendErrs(errs, decodeGSIInt(b[238:243], &gsi.TNB))                     // Total Number of Text and Timing Information (TTI) blocks - bytes 238..242 (5 bytes)
	appendErrs(errs, decodeGSIInt(b[243:248], &gsi.TNS))                     // Total Number of Subtitles - bytes 243..247 (5 bytes)
	appendErrs(errs, decodeGSIInt(b[248:251], &gsi.TNG))                     // Total Number of Subtitle Groups - bytes 248..250 (3 bytes)
	appendErrs(errs, decodeGSIInt(b[251:253], &gsi.MNC))                     // Maximum Number of Displayable Characters in any text row - bytes 251..252 (2 bytes)
	appendErrs(errs, decodeGSIInt(b[253:255], &gsi.MNR))                     // Maximum Number of Displayable Rows - bytes 253..254 (2 bytes)
	appendErrs(errs, decodeGSIByte(b[255:256], (*byte)(&gsi.TCS)))           // Time Code: Status - bytes 255 (1 byte)
	appendErrs(errs, decodeGSITimecode(b[256:264], &gsi.TCP))                // Time Code: Start-of-Program - bytes 256..263 (8 bytes)
	appendErrs(errs, decodeGSITimecode(b[264:272], &gsi.TCF))                // Time Code: First In-Cue - bytes 264..271 (8 bytes)
	appendErrs(errs, decodeGSIInt(b[272:273], &gsi.TND))                     // Total Number of Disks - byte 272 (1 byte)
	appendErrs(errs, decodeGSIInt(b[272:274], &gsi.DSN))                     // Disk Sequence Number - byte 273 (1 byte)
	appendErrs(errs, decodeGSIString(b[274:277], &gsi.CO, gsi.CPN))          // Country of Origin - bytes 274..276 (3 bytes)
	appendErrs(errs, decodeGSIString(b[277:309], &gsi.PUB, gsi.CPN))         // Publisher - bytes 277..308 (32 bytes)
	appendErrs(errs, decodeGSIString(b[309:341], &gsi.EN, gsi.CPN))          // Editor's Name - bytes 309..340 (32 bytes)
	appendErrs(errs, decodeGSIString(b[341:373], &gsi.ECD, gsi.CPN))         // Editor's Contact Details - bytes 341..372 (32 bytes)
	copy(gsi.UDA, b[448:1024])                                               // User-Defined Area - bytes 448..1023 (576 bytes)

	return errs, nil
}

// Encode encodes and writes GSI block to writer.
func (gsi *GSIBlock) Encode(w io.Writer) error {
	b := make([]byte, GSIBlockSize)

	encodeGSIInt(b[0:2], (int)(gsi.CPN)) // Code Page Number - bytes 0..2 (3 bytes)
	// Disk Format Code - bytes 3..10 (8 bytes)
	if err := encodeGSIString(b[3:11], (string)(gsi.DFC), gsi.CPN); err != nil {
		return err
	}
	encodeGSIByte(b[11:12], (byte)(gsi.DSC)) // Display Standard Code - byte 11 (1 byte)
	encodeGSIByte(b[12:14], (byte)(gsi.CCT)) // Character Code Table number - bytes 12..13 (2 bytes)
	encodeGSIHex(b[14:16], (byte)(gsi.LC))   // Language Code - bytes 14..15 (2 bytes)
	// Original Program Title - bytes 16..47 (32 bytes)
	if err := encodeGSIString(b[16:48], gsi.OPT, gsi.CPN); err != nil {
		return err
	}
	// Original Episode Title - bytes 48..79 (32 bytes)
	if err := encodeGSIString(b[48:80], gsi.OET, gsi.CPN); err != nil {
		return err
	}
	// Translated Program Title - bytes 80..111 (32 bytes)
	if err := encodeGSIString(b[80:112], gsi.TPT, gsi.CPN); err != nil {
		return err
	}
	// Translated Episode Title - bytes 112..143 (32 bytes)
	if err := encodeGSIString(b[112:144], gsi.TET, gsi.CPN); err != nil {
		return err
	}
	// Translator's Name - bytes 144..175 (32 bytes)
	if err := encodeGSIString(b[144:176], gsi.TN, gsi.CPN); err != nil {
		return err
	}
	// Translator's Contact Details - bytes 176..207 (32 bytes)
	if err := encodeGSIString(b[176:208], gsi.TCD, gsi.CPN); err != nil {
		return err
	}
	// Subtitle List Reference Code - bytes 208..223 (16 bytes)
	if err := encodeGSIString(b[208:224], gsi.SLR, gsi.CPN); err != nil {
		return err
	}
	encodeGSIDate(b[224:230], gsi.CD)          // Creation Date - bytes 224..229 (6 bytes)
	encodeGSIDate(b[130:236], gsi.RD)          // Revision Date - bytes 230..235 (6 bytes)
	encodeGSIInt(b[236:238], gsi.RN)           // Revision Number - bytes 236..237 (2 bytes)
	encodeGSIInt(b[238:243], gsi.TNB)          // Total Number of Text and Timing Information (TTI) blocks - bytes 238..242 (5 bytes)
	encodeGSIInt(b[243:248], gsi.TNS)          // Total Number of Subtitles - bytes 243..247 (5 bytes)
	encodeGSIInt(b[248:251], gsi.TNG)          // Total Number of Subtitle Groups - bytes 248..250 (3 bytes)
	encodeGSIInt(b[251:253], gsi.MNC)          // Maximum Number of Displayable Characters in any text row - bytes 251..252 (2 bytes)
	encodeGSIInt(b[253:255], gsi.MNR)          // Maximum Number of Displayable Rows - bytes 253..254 (2 bytes)
	encodeGSIByte(b[255:256], (byte)(gsi.TCS)) // Time Code: Status - bytes 255 (1 byte)
	encodeGSITimecode(b[256:264], gsi.TCP)     // Time Code: Start-of-Program - bytes 256..263 (8 bytes)
	encodeGSITimecode(b[264:272], gsi.TCF)     // Time Code: First In-Cue - bytes 264..271 (8 bytes)
	encodeGSIInt(b[272:273], gsi.TND)          // Total Number of Disks - byte 272 (1 byte)
	encodeGSIInt(b[272:274], gsi.DSN)          // Disk Sequence Number - byte 273 (1 byte)
	// Country of Origin - bytes 274..276 (3 bytes)
	if err := encodeGSIString(b[274:277], gsi.CO, gsi.CPN); err != nil {
		return err
	}
	// Publisher - bytes 277..308 (32 bytes)
	if err := encodeGSIString(b[277:309], gsi.PUB, gsi.CPN); err != nil {
		return err
	}
	// Editor's Name - bytes 309..340 (32 bytes)
	if err := encodeGSIString(b[309:341], gsi.EN, gsi.CPN); err != nil {
		return err
	}
	// Editor's Contact Details - bytes 341..372 (32 bytes)
	if err := encodeGSIString(b[341:373], gsi.ECD, gsi.CPN); err != nil {
		return err
	}
	copy(b[448:1024], gsi.UDA) // User-Defined Area - bytes 448..1023 (576 bytes)

	_, err := w.Write(b)
	return err
}

// Validate validates GSI block.
// It returns a slice of ValidateErr containing warnings and fatal errors.
// If a ValidateErr is flaggued as fatal, then the GSI block is considered invalid.
// A warning will be returned if a field in GSI block have "unconventional" value.
// A fatal error will be returned if a field value make the future GSI processing impossible.
func (gsi *GSIBlock) Validate() []error {
	var errs []error
	//todo: validation
	return errs
}

func (gsi *GSIBlock) reset() {
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
	gsi.TCF = Timecode{}
	gsi.TCP = Timecode{}
	gsi.TND = -1
	gsi.DSN = -1
	gsi.CO = ""
	gsi.PUB = ""
	gsi.EN = ""
	gsi.ECD = ""
	gsi.UDA = []byte{}
}
