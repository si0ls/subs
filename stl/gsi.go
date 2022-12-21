package stl

import (
	"fmt"
	"io"
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

// Decode reads and decodes GSI block from reader.
func (gsi *GSIBlock) Decode(r io.Reader) ([]error, error) {
	b := make([]byte, GSIBlockSize)
	if _, err := io.ReadFull(r, b); err != nil {
		return nil, err
	}

	gsi.Reset()

	var errs []error

	appendNonNilErrs(errs, gsiErr(decodeGSIInt(b[0:3], (*int)(&gsi.CPN)), GSIFieldCPN))                 // CPN - bytes 0..2 (3 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[3:11], (*string)(&gsi.DFC), gsi.CPN), GSIFieldDFC)) // DFC - bytes 3..10 (8 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIByte(b[11:12], (*byte)(&gsi.DSC)), GSIFieldDSC))             // DSC - byte 11 (1 byte)
	appendNonNilErrs(errs, gsiErr(decodeGSIByte(b[12:14], (*byte)(&gsi.CCT)), GSIFieldCCT))             // CCT - bytes 12..13 (2 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIHex(b[14:16], (*byte)(&gsi.LC)), GSIFieldLC))                // LC - bytes 14..15 (2 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[16:48], &gsi.OPT, gsi.CPN), GSIFieldOPT))           // OPT - bytes 16..47 (32 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[48:80], &gsi.OET, gsi.CPN), GSIFieldOET))           // OET - bytes 48..79 (32 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[80:112], &gsi.TPT, gsi.CPN), GSIFieldTPT))          // TPT - bytes 80..111 (32 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[112:144], &gsi.TET, gsi.CPN), GSIFieldTET))         // TET - bytes 112..143 (32 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[144:176], &gsi.TN, gsi.CPN), GSIFieldTN))           // TN - bytes 144..175 (32 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[176:208], &gsi.TCD, gsi.CPN), GSIFieldTCD))         // TCD - bytes 176..207 (32 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[208:224], &gsi.SLR, gsi.CPN), GSIFieldSLR))         // SLR - bytes 208..223 (16 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIDate(b[224:230], &gsi.CD), GSIFieldCD))                      // CD - bytes 224..229 (6 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIDate(b[230:236], &gsi.RD), GSIFieldRD))                      // RD - bytes 230..235 (6 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIInt(b[236:238], &gsi.RN), GSIFieldRN))                       // RN - bytes 236..237 (2 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIInt(b[238:243], &gsi.TNB), GSIFieldTNB))                     // TNB - bytes 238..242 (5 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIInt(b[243:248], &gsi.TNS), GSIFieldTNS))                     // TNB - bytes 243..247 (5 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIInt(b[248:251], &gsi.TNG), GSIFieldTNG))                     // TNG - bytes 248..250 (3 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIInt(b[251:253], &gsi.MNC), GSIFieldMNC))                     // MNC - bytes 251..252 (2 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIInt(b[253:255], &gsi.MNR), GSIFieldMNR))                     // MNR - bytes 253..254 (2 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIByte(b[255:256], (*byte)(&gsi.TCS)), GSIFieldTCS))           // TCS - bytes 255 (1 byte)
	appendNonNilErrs(errs, gsiErr(decodeGSITimecode(b[256:264], &gsi.TCP), GSIFieldTCP))                // TCP - bytes 256..263 (8 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSITimecode(b[264:272], &gsi.TCF), GSIFieldTCF))                // TCF - bytes 264..271 (8 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIInt(b[272:273], &gsi.TND), GSIFieldTND))                     // TND - byte 272 (1 byte)
	appendNonNilErrs(errs, gsiErr(decodeGSIInt(b[272:274], &gsi.DSN), GSIFieldDSN))                     // DSN - byte 273 (1 byte)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[274:277], &gsi.CO, gsi.CPN), GSIFieldCO))           // CO - bytes 274..276 (3 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[277:309], &gsi.PUB, gsi.CPN), GSIFieldPUB))         // PUB - bytes 277..308 (32 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[309:341], &gsi.EN, gsi.CPN), GSIFieldEN))           // EN - bytes 309..340 (32 bytes)
	appendNonNilErrs(errs, gsiErr(decodeGSIString(b[341:373], &gsi.ECD, gsi.CPN), GSIFieldECD))         // ECD - bytes 341..372 (32 bytes)
	copy(gsi.UDA, b[448:1024])                                                                          // UDA - bytes 448..1023 (576 bytes)

	return errs, nil
}

// Encode encodes and writes GSI block to writer.
func (gsi *GSIBlock) Encode(w io.Writer) error {
	b := make([]byte, GSIBlockSize)

	// CPN - bytes 0..2 (3 bytes)
	encodeGSIInt(b[0:2], (int)(gsi.CPN))

	// DFC - bytes 3..10 (8 bytes)
	if err := encodeGSIString(b[3:11], (string)(gsi.DFC), gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldDFC)
	}

	// DSC - byte 11 (1 byte)
	encodeGSIByte(b[11:12], (byte)(gsi.DSC))

	// CCT - bytes 12..13 (2 bytes)
	encodeGSIByte(b[12:14], (byte)(gsi.CCT))

	// LC - bytes 14..15 (2 bytes)
	encodeGSIHex(b[14:16], (byte)(gsi.LC))

	// OPT - bytes 16..47 (32 bytes)
	if err := encodeGSIString(b[16:48], gsi.OPT, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldOPT)
	}

	// OET - bytes 48..79 (32 bytes)
	if err := encodeGSIString(b[48:80], gsi.OET, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldOET)
	}

	// TPT - bytes 80..111 (32 bytes)
	if err := encodeGSIString(b[80:112], gsi.TPT, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldTPT)
	}

	// TET - bytes 112..143 (32 bytes)
	if err := encodeGSIString(b[112:144], gsi.TET, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldTET)
	}

	// TN - bytes 144..175 (32 bytes)
	if err := encodeGSIString(b[144:176], gsi.TN, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldTN)
	}

	// TCD - bytes 176..207 (32 bytes)
	if err := encodeGSIString(b[176:208], gsi.TCD, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldTCD)
	}

	// SLR - bytes 208..223 (16 bytes)
	if err := encodeGSIString(b[208:224], gsi.SLR, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldSLR)
	}

	// CD - bytes 224..229 (6 bytes)
	encodeGSIDate(b[224:230], gsi.CD)

	// RD - bytes 230..235 (6 bytes)
	encodeGSIDate(b[130:236], gsi.RD)

	// RN - bytes 236..237 (2 bytes)
	encodeGSIInt(b[236:238], gsi.RN)

	// TNB - bytes 238..242 (5 bytes)
	encodeGSIInt(b[238:243], gsi.TNB)

	// TNS - bytes 243..247 (5 bytes)
	encodeGSIInt(b[243:248], gsi.TNS)

	// TNG - bytes 248..250 (3 bytes)
	encodeGSIInt(b[248:251], gsi.TNG)

	// MNC - bytes 251..252 (2 bytes)
	encodeGSIInt(b[251:253], gsi.MNC)

	// MNR - bytes 253..254 (2 bytes)
	encodeGSIInt(b[253:255], gsi.MNR)

	// TCS - bytes 255 (1 byte)
	encodeGSIByte(b[255:256], (byte)(gsi.TCS))

	// TCP - bytes 256..263 (8 bytes)
	encodeGSITimecode(b[256:264], gsi.TCP)

	// TCF - bytes 264..271 (8 bytes)
	encodeGSITimecode(b[264:272], gsi.TCF)

	// TND - byte 272 (1 byte)
	encodeGSIInt(b[272:273], gsi.TND)

	// DSN - byte 273 (1 byte)
	encodeGSIInt(b[272:274], gsi.DSN)

	// CO - bytes 274..276 (3 bytes)
	if err := encodeGSIString(b[274:277], gsi.CO, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldCO)
	}

	// PUB - bytes 277..308 (32 bytes)
	if err := encodeGSIString(b[277:309], gsi.PUB, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldPUB)
	}

	// EN - bytes 309..340 (32 bytes)
	if err := encodeGSIString(b[309:341], gsi.EN, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldEN)
	}

	// ECD - bytes 341..372 (32 bytes)
	if err := encodeGSIString(b[341:373], gsi.ECD, gsi.CPN); err != nil {
		return gsiErr(err, GSIFieldECD)
	}

	// UDA - bytes 448..1023 (576 bytes)
	copy(b[448:1024], gsi.UDA)

	_, err := w.Write(b)
	return err
}

// GSIError is an error that occurred on a GSI field.
// It extends FieldError that carries the concerned GSI field.
type GSIError struct {
	error
	FieldError
	field GSIField
}

func gsiErr(err error, field GSIField) *GSIError {
	return &GSIError{error: err, field: field}
}

// Error returns the error message.
func (e *GSIError) Error() string {
	return fmt.Sprintf("GSI %s: %s", e.field, e.error.Error())
}

// Unwrap returns the underlying error.
func (e *GSIError) Unwrap() error {
	return e.error
}

// Field returns the concerned GSI field.
func (e *GSIError) Field() GSIField {
	return e.field
}
