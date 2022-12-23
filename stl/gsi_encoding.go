package stl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// Decode reads and decodes GSI block from reader.
// It returns a slice of warnings and an error if any.
// Warnings are returned for each field that is invalid and are not fatal.
// An error is returned if a fatal error occurs that prevents further decoding.
func (gsi *GSIBlock) Decode(r io.Reader) ([]error, error) {
	b := make([]byte, GSIBlockSize)
	if _, err := io.ReadFull(r, b); err != nil {
		return nil, err
	}

	gsi.Reset()

	var warns []error

	warns = appendNonNilErrs(warns, gsiErr(decodeGSIInt(b[0:3], (*int)(&gsi.CPN)), GSIFieldCPN))                 // CPN - bytes 0..2 (3 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[3:11], (*string)(&gsi.DFC), gsi.CPN), GSIFieldDFC)) // DFC - bytes 3..10 (8 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIByte(b[11:12], (*byte)(&gsi.DSC)), GSIFieldDSC))             // DSC - byte 11 (1 byte)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIByte(b[12:14], (*byte)(&gsi.CCT)), GSIFieldCCT))             // CCT - bytes 12..13 (2 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIHex(b[14:16], (*byte)(&gsi.LC)), GSIFieldLC))                // LC - bytes 14..15 (2 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[16:48], &gsi.OPT, gsi.CPN), GSIFieldOPT))           // OPT - bytes 16..47 (32 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[48:80], &gsi.OET, gsi.CPN), GSIFieldOET))           // OET - bytes 48..79 (32 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[80:112], &gsi.TPT, gsi.CPN), GSIFieldTPT))          // TPT - bytes 80..111 (32 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[112:144], &gsi.TET, gsi.CPN), GSIFieldTET))         // TET - bytes 112..143 (32 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[144:176], &gsi.TN, gsi.CPN), GSIFieldTN))           // TN - bytes 144..175 (32 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[176:208], &gsi.TCD, gsi.CPN), GSIFieldTCD))         // TCD - bytes 176..207 (32 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[208:224], &gsi.SLR, gsi.CPN), GSIFieldSLR))         // SLR - bytes 208..223 (16 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIDate(b[224:230], &gsi.CD), GSIFieldCD))                      // CD - bytes 224..229 (6 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIDate(b[230:236], &gsi.RD), GSIFieldRD))                      // RD - bytes 230..235 (6 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIInt(b[236:238], &gsi.RN), GSIFieldRN))                       // RN - bytes 236..237 (2 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIInt(b[238:243], &gsi.TNB), GSIFieldTNB))                     // TNB - bytes 238..242 (5 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIInt(b[243:248], &gsi.TNS), GSIFieldTNS))                     // TNS - bytes 243..247 (5 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIInt(b[248:251], &gsi.TNG), GSIFieldTNG))                     // TNG - bytes 248..250 (3 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIInt(b[251:253], &gsi.MNC), GSIFieldMNC))                     // MNC - bytes 251..252 (2 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIInt(b[253:255], &gsi.MNR), GSIFieldMNR))                     // MNR - bytes 253..254 (2 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIByte(b[255:256], (*byte)(&gsi.TCS)), GSIFieldTCS))           // TCS - bytes 255 (1 byte)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSITimecode(b[256:264], &gsi.TCP), GSIFieldTCP))                // TCP - bytes 256..263 (8 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSITimecode(b[264:272], &gsi.TCF), GSIFieldTCF))                // TCF - bytes 264..271 (8 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIInt(b[272:273], &gsi.TND), GSIFieldTND))                     // TND - byte 272 (1 byte)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIInt(b[272:274], &gsi.DSN), GSIFieldDSN))                     // DSN - byte 273 (1 byte)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[274:277], &gsi.CO, gsi.CPN), GSIFieldCO))           // CO - bytes 274..276 (3 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[277:309], &gsi.PUB, gsi.CPN), GSIFieldPUB))         // PUB - bytes 277..308 (32 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[309:341], &gsi.EN, gsi.CPN), GSIFieldEN))           // EN - bytes 309..340 (32 bytes)
	warns = appendNonNilErrs(warns, gsiErr(decodeGSIString(b[341:373], &gsi.ECD, gsi.CPN), GSIFieldECD))         // ECD - bytes 341..372 (32 bytes)
	copy(gsi.UDA, b[448:1024])                                                                                   // UDA - bytes 448..1023 (576 bytes)

	return warns, nil
}

// Encode encodes and writes GSI block to writer.
// An error is returned if a fatal error occurs that prevents further encoding.
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

func decodeGSIInt(b []byte, v *int) error {
	s := strings.Trim(string(b), string([]byte(" ")))
	if len(s) == 0 {
		return decodeErr(ErrEmptyGSIIntValue, b)
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return decodeErr(ErrInvalidGSIIntValue, b)
	}
	*v = i
	return nil
}

func encodeGSIInt(b []byte, v int) {
	if v < 0 {
		copy(b, bytes.Repeat([]byte(" "), len(b)))
		return
	}
	for i := len(b) - 1; i >= 0; i-- {
		b[i] = byte(v%10) + '0'
		v /= 10
	}
}

func decodeGSIByte(b []byte, v *byte) error {
	if len(b) > 2 {
		panic(fmt.Errorf("invalid GSI byte length %d", len(b)))
	}
	var tmp int
	if err := decodeGSIInt(b, &tmp); err != nil {
		if errors.Unwrap(err) == ErrEmptyGSIIntValue {
			return decodeErr(ErrEmptyGSIByteValue, b)
		}
		return decodeErr(ErrInvalidGSIByteValue, b)
	}
	*v = byte(tmp)
	return nil
}

func encodeGSIByte(b []byte, v byte) {
	if len(b) > 2 {
		panic(fmt.Errorf("invalid GSI byte length %d", len(b)))
	}
	encodeGSIInt(b, int(v))
}

func decodeGSIHex(b []byte, v *byte) error {
	if len(b) != 2 {
		panic(fmt.Errorf("invalid GSI hex length %d", len(b)))
	}

	s := strings.Trim(string(b), string([]byte(" ")))
	if len(s) == 0 {
		return decodeErr(ErrEmptyGSIHexValue, b)
	}
	i, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return decodeErr(ErrInvalidGSIHexValue, b)
	}
	*v = byte(i)
	return nil
}

var hex = "0123456789ABCDEF"

func encodeGSIHex(b []byte, v byte) {
	if len(b) != 2 {
		panic(fmt.Errorf("invalid GSI hex length %d", len(b)))
	}

	for i := len(b) - 1; i >= 0; i-- {
		b[i] = hex[v%16]
		v /= 16
	}
}

func decodeGSIString(b []byte, v *string, cpn CodePageNumber) error {
	if dec, ok := codePageNumberDecoders[cpn]; ok {
		b, err := dec.Decode(bytes.TrimRight(b, string([]byte(" "))))
		if err != nil {
			return decodeErr(ErrInvalidGSIStringValue, b)
		}
		*v = string(b)
	} else {
		return decodeErr(ErrUnsupportedGSICodePage, b)
	}
	return nil
}

func encodeGSIString(b []byte, v string, cpn CodePageNumber) error {
	if enc, ok := codePageNumberEncoders[cpn]; ok {
		e, err := enc.Encode([]byte(v))
		if err != nil {
			return encodeErr(ErrInvalidGSIStringValue, []byte(v))
		}
		copy(b, cutPad(e, len(b), ' '))
	} else {
		return encodeErr(ErrUnsupportedGSICodePage, []byte(v))
	}
	return nil
}

func decodeGSIDate(b []byte, v *time.Time) error {
	if len(b) != 6 {
		panic(fmt.Errorf("invalid GSI date length %d", len(b)))
	}

	var year int = 0
	var month int = 1
	var day int = 1

	if err := decodeGSIInt(b[0:2], &year); err != nil {
		return err
	}
	if err := decodeGSIInt(b[2:4], &month); err != nil {
		return err
	}
	if err := decodeGSIInt(b[4:6], &day); err != nil {
		return err
	}

	if year != 0 || month != 1 || day != 1 {
		*v = time.Date(year+2000, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	return nil
}

func encodeGSIDate(b []byte, v time.Time) {
	if len(b) != 6 {
		panic(fmt.Errorf("invalid GSI date length %d", len(b)))
	}

	encodeGSIInt(b[0:2], v.Year()-2000)
	encodeGSIInt(b[2:4], int(v.Month()))
	encodeGSIInt(b[4:6], v.Day())
}

func decodeGSITimecode(b []byte, v *Timecode) error {
	if len(b) != 8 {
		panic(fmt.Errorf("invalid GSI timecode length %d", len(b)))
	}

	var err error
	if derr := decodeGSIInt(b[0:2], &v.Hours); derr != nil {
		err = derr
	}
	if derr := decodeGSIInt(b[2:4], &v.Minutes); derr != nil {
		err = derr
	}
	if derr := decodeGSIInt(b[4:6], &v.Seconds); derr != nil {
		err = derr
	}
	if derr := decodeGSIInt(b[6:8], &v.Frames); derr != nil {
		err = derr
	}

	if err != nil {
		return decodeErr(ErrInvalidGSITimecodeValue, b)
	}

	return nil
}

func encodeGSITimecode(b []byte, v Timecode) {
	if len(b) != 8 {
		panic(fmt.Errorf("invalid GSI timecode length %d", len(b)))
	}

	encodeGSIInt(b[0:2], v.Hours)
	encodeGSIInt(b[2:4], v.Minutes)
	encodeGSIInt(b[4:6], v.Seconds)
	encodeGSIInt(b[6:8], v.Frames)
}
