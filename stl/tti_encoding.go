package stl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Decode reads and decodes TTI block from reader.
func (tti *TTIBlock) Decode(r io.Reader) error {
	b := make([]byte, TTIBlockSize)
	if _, err := io.ReadFull(r, b); err != nil {
		return err
	}

	tti.Reset()

	decodeTTIInt(b[0:1], &tti.SGN)            // SGN - byte 0 (1 byte)
	decodeTTIInt(b[1:3], &tti.SN)             // SN - bytes 1..2 (2 bytes)
	decodeTTIInt(b[3:4], &tti.EBN)            // EBN - byte 3 (1 byte)
	decodeTTIByte(b[4:5], (*byte)(&tti.CS))   // CS - byte 4 (1 byte)
	decodeTTITimecode(b[5:9], &tti.TCI)       // TCI - bytes 5..8 (4 bytes)
	decodeTTITimecode(b[9:13], &tti.TCO)      // TCO - bytes 9..12 (4 bytes)
	decodeTTIInt(b[13:14], &tti.VP)           // VP - byte 13 (1 byte)
	decodeTTIByte(b[14:15], (*byte)(&tti.JC)) // JC - byte 14 (1 byte)
	decodeTTIByte(b[15:16], (*byte)(&tti.CF)) // CF - byte 15 (1 byte)
	tti.terminatedBySpace = b[127] == 0x8F    // Store if the last byte is 0x8F (space) for further validation
	decodeTTIString(b[16:128], &tti.TF)       // Text Field (TF) - bytes 16..127 (112 bytes)

	return nil
}

// Encode encodes and writes TTI block to writer.
func (tti *TTIBlock) Encode(w io.Writer) error {
	b := make([]byte, TTIBlockSize)

	encodeTTIInt(b[0:1], tti.SGN)           // SGN - byte 0 (1 byte)
	encodeTTIInt(b[1:3], tti.SN)            // SN - bytes 1..2 (2 bytes)
	encodeTTIInt(b[3:4], tti.EBN)           // EBN - byte 3 (1 byte)
	encodeTTIByte(b[4:5], (byte)(tti.CS))   // CS - byte 4 (1 byte)
	encodeTTITimecode(b[5:9], tti.TCI)      // TCI - bytes 5..8 (4 bytes)
	encodeTTITimecode(b[9:13], tti.TCO)     // TCO - bytes 9..12 (4 bytes)
	encodeTTIInt(b[13:14], tti.VP)          // VP - byte 13 (1 byte)
	encodeTTIByte(b[14:15], (byte)(tti.JC)) // JC - byte 14 (1 byte)
	encodeTTIByte(b[15:16], (byte)(tti.CF)) // CF - byte 15 (1 byte)
	encodeTTIString(b[16:128], tti.TF)      // TF - bytes 16..127 (112 bytes)

	_, err := w.Write(b)
	return err
}

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
