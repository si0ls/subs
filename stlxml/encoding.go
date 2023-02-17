package stlxml

import (
	"fmt"
	"strings"
	"time"

	"github.com/si0ls/subs/stl"
)

// decodeString decodes a string.
func decodeString(s string) string {
	return s
}

// encodeString encodes a string padded with spaces to the specified length.
func encodeString(s string, length int) string {
	return fmt.Sprintf("%-*s", length, s)
}

// decodeInt decodes a string as an int.
func decodeInt(s string) (int, error) {
	if s == "" {
		return -1, nil
	}
	var i int
	if _, err := fmt.Sscanf(s, "%d", &i); err != nil {
		return -1, err
	}
	return i, nil
}

// encodeInt encodes an int as a string padded with zeros to the specified length.
func encodeInt(i int, length int) string {
	if i < 0 {
		return strings.Repeat(" ", length)
	}
	return fmt.Sprintf("%0*d", length, i)
}

// decodeByte decodes a string as a byte.
func decodeByte(s string) (byte, error) {
	if s == "" {
		return 0xFF, nil
	}
	var i int
	if _, err := fmt.Sscanf(s, "%d", &i); err != nil {
		return 0xFF, err
	}
	return byte(i), nil
}

// encodeByte encodes a byte as a string padded with zeros to the specified length.
func encodeByte(b byte, length int) string {
	if b == 0xFF {
		return strings.Repeat(" ", length)
	}
	return fmt.Sprintf("%0*d", length, b)
}

// decodeHexInt decodes a hex string as an int.
func decodeHexInt(s string, upper bool) (int, error) {
	if s == "" {
		return -1, nil
	}
	var i int
	format := "%x"
	if upper {
		format = "%X"
	}
	if _, err := fmt.Sscanf(s, format, &i); err != nil {
		return -1, err
	}
	return i, nil
}

// encodeHexInt encodes an int as a hex string padded with zeros to the specified length.
func encodeHexInt(i int, length int, upper bool) string {
	if i < 0 {
		return strings.Repeat(" ", length)
	}
	format := "%0*x"
	if upper {
		format = "%0*X"
	}
	return fmt.Sprintf(format, length, i)
}

// decodeHexByte decodes a hex string as an byte.
func decodeHexByte(s string, upper bool) (byte, error) {
	if s == "" {
		return 0xFF, nil
	}
	var b byte
	format := "%x"
	if upper {
		format = "%X"
	}
	if _, err := fmt.Sscanf(s, format, &b); err != nil {
		return 0xFF, err
	}
	return b, nil
}

// encodeHexByte encodes an byte as a hex string padded with zeros to the specified length.
func encodeHexByte(b byte, length int, upper bool) string {
	if b == 0xFF {
		return strings.Repeat(" ", length)
	}
	format := "%0*x"
	if upper {
		format = "%0*X"
	}
	return fmt.Sprintf(format, length, b)
}

// decodeTimecode decodes a string as a stl.Timecode.
func decodeTimecode(s string) (stl.Timecode, error) {
	if s == "" {
		return stl.Timecode{}, nil
	}
	var tc stl.Timecode
	if _, err := fmt.Sscanf(s, "%02d%02d%02d%02d", &tc.Hours, &tc.Minutes, &tc.Seconds, &tc.Frames); err != nil {
		return stl.Timecode{}, err
	}
	return tc, nil
}

// encodeTimecode encodes a stl.Timecode as a string.
func encodeTimecode(tc stl.Timecode) string {
	if tc.Hours < 0 {
		return "        "
	}
	return fmt.Sprintf("%02d%02d%02d%02d", tc.Hours, tc.Minutes, tc.Seconds, tc.Frames)
}

// decodeTime decodes a string as a time.Time.
func decodeTime(s string) (time.Time, error) {
	if strings.TrimSpace(s) == "" {
		return time.Time{}, nil
	}
	t, err := time.Parse("060102", s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// encodeTime encodes a time.Time as a string.
func encodeTime(t time.Time) string {
	if t.IsZero() {
		return "      "
	}
	return t.Format("060102")
}
