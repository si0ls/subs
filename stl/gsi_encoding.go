package stl

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

	var err error
	if derr := decodeGSIInt(b[0:2], &year); derr != nil {
		err = derr
	}
	if derr := decodeGSIInt(b[2:4], &month); derr != nil {
		err = derr
	}
	if derr := decodeGSIInt(b[4:6], &day); derr != nil {
		err = derr
	}

	*v = time.Date(year+2000, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	if err != nil {
		return decodeErr(ErrInvalidGSIDateValue, b)
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
