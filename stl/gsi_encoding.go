package stl

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func decodeGSIInt(b []byte, v *int) error {
	s := strings.Trim(string(b), string([]byte(" ")))
	if len(s) == 0 {
		return fmt.Errorf("empty GSI int value %s", b)
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("invalid GSI int value %s (%s)", b, err)
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
		return err
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
		return fmt.Errorf("empty GSI int value %s", b)
	}
	i, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return fmt.Errorf("invalid GSI int value %s (%s)", b, err)
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
			return fmt.Errorf("invalid GSI string value %s (%s)", b, err)
		}
		*v = string(b)
	} else {
		panic(fmt.Errorf("unsupported GSI string code page %d", cpn))
	}
	return nil
}

func encodeGSIString(b []byte, v string, cpn CodePageNumber) error {
	if enc, ok := codePageNumberEncoders[cpn]; ok {
		b, err := enc.Encode([]byte(v))
		if err != nil {
			return fmt.Errorf("invalid GSI string value %s (%s)", b, err)
		}
		copy(b, cutPad(b, len(b), ' '))
	} else {
		panic(fmt.Errorf("unsupported GSI string code page %d", cpn))
	}
	return nil
}

func decodeGSIDate(b []byte, v *time.Time) error {
	if len(b) != 6 {
		panic(fmt.Errorf("invalid GSI date length %d", len(b)))
	}

	tmp, err := time.Parse("060102", string(b))
	if err != nil {
		return fmt.Errorf("invalid GSI date %s (%s)", b, err)
	}
	*v = tmp
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

	if err := decodeGSIInt(b[0:2], &v.Hours); err != nil {
		return err
	}
	if err := decodeGSIInt(b[2:4], &v.Minutes); err != nil {
		return err
	}
	if err := decodeGSIInt(b[4:6], &v.Seconds); err != nil {
		return err
	}
	if err := decodeGSIInt(b[6:8], &v.Frames); err != nil {
		return err
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
