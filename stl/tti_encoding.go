package stl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func decodeTTIInt(b []byte, v *int) {
	c := make([]byte, 8)
	copy(c[:8-len(b)], b)
	*v = int(binary.LittleEndian.Uint64(c))
}

func encodeTTIInt(b []byte, v int) {
	if len(b) > 8 {
		panic(fmt.Errorf("invalid TTI int length %d", len(b)))
	}

	c := make([]byte, 8)
	if v < 0 {
		copy(b, []byte(bytes.Repeat([]byte("\x00"), len(b))))
	} else {
		binary.LittleEndian.PutUint64(c, uint64(v))
	}
	copy(b, c[8-len(b):])
}

func decodeTTIString(b []byte, v *string) {
	*v = strings.TrimRight(string(b), string([]byte{0x8F}))
}

func encodeTTIString(b []byte, v string) {
	copy(b, cutPad([]byte(v), len(b), 0x8F))
}

func decodeTTIByte(b []byte, v *byte) {
	if len(b) != 1 {
		panic(fmt.Errorf("invalid TTI byte length %d", len(b)))
	}

	*v = b[0]
}

func encodeTTIByte(b []byte, v byte) {
	if len(b) != 1 {
		panic(fmt.Errorf("invalid TTI byte length %d", len(b)))
	}

	switch v {
	case 0xFF:
		b[0] = ' '
	default:
		b[0] = v
	}
}

func decodeTTITimecode(b []byte, v *Timecode) {
	if len(b) != 4 {
		panic(fmt.Errorf("invalid TTI timecode length %d", len(b)))
	}

	decodeTTIInt(b[0:1], &v.Hours)
	decodeTTIInt(b[1:2], &v.Minutes)
	decodeTTIInt(b[2:3], &v.Seconds)
	decodeTTIInt(b[3:4], &v.Frames)
}

func encodeTTITimecode(b []byte, v Timecode) {
	if len(b) != 4 {
		panic(fmt.Errorf("invalid TTI timecode length %d", len(b)))
	}

	encodeTTIInt(b[0:1], v.Hours)
	encodeTTIInt(b[1:2], v.Minutes)
	encodeTTIInt(b[2:3], v.Seconds)
	encodeTTIInt(b[3:4], v.Frames)
}
