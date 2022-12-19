package stl

import (
	"bytes"
	"testing"
	"time"
)

type decodeGSIIntTest struct {
	input  []byte
	output int
	err    error
	panic  bool
}

var decodeGSIIntTests = []decodeGSIIntTest{
	{[]byte("   123"), 123, nil, false},
	{[]byte("123   "), 123, nil, false},
	{[]byte("  123 "), 123, nil, false},
	{[]byte("   "), 0, ErrEmptyGSIIntValue, false},
	{[]byte(""), 0, ErrEmptyGSIIntValue, false},
	{[]byte("ABC"), 0, ErrInvalidGSIIntValue, false},
	{[]byte("0"), 0, nil, false},
	{[]byte("0000"), 0, nil, false},
	{[]byte("2147483647"), 2147483647, nil, false},
}

func TestDecodeGSIInt(t *testing.T) {
	for _, test := range decodeGSIIntTests {
		func() {
			defer testPanic(t, test.panic)

			var v int
			err := decodeGSIInt(test.input, &v)
			testError(t, err, test.err)
			if v != test.output {
				t.Errorf("expected %d but got %d", test.output, v)
			}
		}()
	}
}

type encodeGSIIntTest struct {
	input  int
	length int
	output []byte
	panic  bool
}

var encodeGSIIntTests = []encodeGSIIntTest{
	{123, 3, []byte("123"), false},
	{123, 6, []byte("000123"), false},
	{123, 2, []byte("23"), false},
	{123, 0, []byte(""), false},
	{0, 1, []byte("0"), false},
	{-1, 3, []byte("   "), false},
	{2147483647, 10, []byte("2147483647"), false},
}

func TestEncodeGSIInt(t *testing.T) {
	for _, test := range encodeGSIIntTests {
		func() {
			defer testPanic(t, test.panic)

			var v []byte = make([]byte, test.length)
			encodeGSIInt(v, test.input)
			if !bytes.Equal(v, test.output) {
				t.Errorf("expected %s but got %s", test.output, v)
			}
		}()
	}
}

type decodeGSIByteTest struct {
	input  []byte
	output byte
	err    error
	panic  bool
}

var decodeGSIByteTests = []decodeGSIByteTest{
	{[]byte(" 1"), 1, nil, false},
	{[]byte("1 "), 1, nil, false},
	{[]byte("12"), 12, nil, false},
	{[]byte("  "), 0, ErrEmptyGSIByteValue, false},
	{[]byte(""), 0, ErrEmptyGSIByteValue, false},
	{[]byte("AB"), 0, ErrInvalidGSIByteValue, false},
	{[]byte("123"), 0, nil, true},
	{[]byte("0"), 0, nil, false},
	{[]byte("00"), 0, nil, false},
	{[]byte("99"), 99, nil, false},
}

func TestDecodeGSIByte(t *testing.T) {
	for _, test := range decodeGSIByteTests {
		func() {
			defer testPanic(t, test.panic)

			var v byte
			err := decodeGSIByte(test.input, &v)
			testError(t, err, test.err)
			if v != test.output {
				t.Errorf("expected %d but got %d", test.output, v)
			}
		}()
	}
}

type encodeGSIByteTest struct {
	input  byte
	length int
	output []byte
	panic  bool
}

var encodeGSIByteTests = []encodeGSIByteTest{
	{1, 1, []byte("1"), false},
	{1, 2, []byte("01"), false},
	{1, 0, []byte(""), false},
	{0, 1, []byte("0"), false},
	{0, 2, []byte("00"), false},
	{99, 2, []byte("99"), false},
	{123, 2, []byte("23"), false},
	{123, 3, []byte(""), true},
}

func TestEncodeGSIByte(t *testing.T) {
	for _, test := range encodeGSIByteTests {
		func() {
			defer testPanic(t, test.panic)

			var v []byte = make([]byte, test.length)
			encodeGSIByte(v, test.input)
			if !bytes.Equal(v, test.output) {
				t.Errorf("expected %s but got %s", test.output, v)
			}
		}()
	}
}

type decodeGSIHexTest struct {
	input  []byte
	output byte
	err    error
	panic  bool
}

var decodeGSIHexTests = []decodeGSIHexTest{
	{[]byte("12"), 18, nil, false},
	{[]byte(" 1"), 1, nil, false},
	{[]byte("1 "), 1, nil, false},
	{[]byte("00"), 0, nil, false},
	{[]byte("FF"), 255, nil, false},
	{[]byte("GG"), 0, ErrInvalidGSIHexValue, false},
	{[]byte("  "), 0, ErrEmptyGSIHexValue, false},
	{[]byte(""), 0, nil, true},
	{[]byte("ABC"), 0, nil, true},
}

func TestDecodeGSIHex(t *testing.T) {
	for _, test := range decodeGSIHexTests {
		func() {
			defer testPanic(t, test.panic)

			var v byte
			err := decodeGSIHex(test.input, &v)
			testError(t, err, test.err)
			if v != test.output {
				t.Errorf("expected %v but got %v", test.output, v)
			}
		}()
	}
}

type encodeGSIHexTest struct {
	input  byte
	length int
	output []byte
	panic  bool
}

var encodeGSIHexTests = []encodeGSIHexTest{
	{1, 2, []byte("01"), false},
	{0, 2, []byte("00"), false},
	{12, 2, []byte("0C"), false},
	{18, 2, []byte("12"), false},
	{255, 2, []byte("FF"), false},
	{123, 2, []byte("7B"), false},
	{12, 0, []byte(""), true},
	{0, 0, []byte(""), true},
	{123, 3, []byte(""), true},
}

func TestEncodeGSIHex(t *testing.T) {
	for _, test := range encodeGSIHexTests {
		func() {
			defer testPanic(t, test.panic)

			var v []byte = make([]byte, test.length)
			encodeGSIHex(v, test.input)
			if !bytes.Equal(v, test.output) {
				t.Errorf("expected %s but got %s", test.output, v)
			}
		}()
	}
}

type decodeGSIStringTest struct {
	input  []byte
	output string
	cpn    CodePageNumber
	err    error
}

var decodeGSIStringTests = []decodeGSIStringTest{
	{[]byte("ABC"), "ABC", CodePageNumberMultiLingual, nil},
	{[]byte("ABC"), "", CodePageNumberInvalid, ErrUnsupportedGSICodePage},
}

func TestDecodeGSIString(t *testing.T) {
	for _, test := range decodeGSIStringTests {
		var v string
		err := decodeGSIString(test.input, &v, test.cpn)
		testError(t, err, test.err)
		if v != test.output {
			t.Errorf("expected %s but got %s", test.output, v)
		}
	}
}

type encodeGSIStringTest struct {
	input  string
	output []byte
	cpn    CodePageNumber
	err    error
}

var encodeGSIStringTests = []encodeGSIStringTest{
	{"ABC", []byte("ABC"), CodePageNumberMultiLingual, nil},
	{"ABC", []byte{0, 0, 0}, CodePageNumberInvalid, ErrUnsupportedGSICodePage},
}

func TestEncodeGSIString(t *testing.T) {
	for _, test := range encodeGSIStringTests {
		var v []byte = make([]byte, len(test.input))
		err := encodeGSIString(v, test.input, test.cpn)
		testError(t, err, test.err)
		if !bytes.Equal(v, test.output) {
			t.Errorf("expected %x but got %x", test.output, v)
		}
	}
}

type decodeGSIDateTest struct {
	input  []byte
	output time.Time
	err    error
	panic  bool
}

var decodeGSIDateTests = []decodeGSIDateTest{
	{[]byte("170302"), time.Date(2017, 3, 2, 0, 0, 0, 0, time.UTC), nil, false},
	{[]byte("1703AB"), time.Date(2017, 3, 1, 0, 0, 0, 0, time.UTC), ErrInvalidGSIDateValue, false},
	{[]byte("AB0302"), time.Date(2000, 3, 2, 0, 0, 0, 0, time.UTC), ErrInvalidGSIDateValue, false},
	{[]byte("17AB02"), time.Date(2017, 1, 2, 0, 0, 0, 0, time.UTC), ErrInvalidGSIDateValue, false},
	{[]byte("1703  "), time.Date(2017, 3, 1, 0, 0, 0, 0, time.UTC), ErrInvalidGSIDateValue, false},
	{[]byte("1703"), time.Time{}, nil, true},
}

func TestDecodeGSIDate(t *testing.T) {
	for _, test := range decodeGSIDateTests {
		func() {
			defer testPanic(t, test.panic)

			var v time.Time
			err := decodeGSIDate(test.input, &v)
			testError(t, err, test.err)
			if !v.Equal(test.output) {
				t.Errorf("expected %s but got %s", test.output, v)
			}
		}()
	}
}

type encodeGSIDateTest struct {
	input  time.Time
	output []byte
	panic  bool
}

var encodeGSIDateTests = []encodeGSIDateTest{
	{time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC), []byte("170101"), false},
	{time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC), []byte("1701"), true},
	{time.Time{}, []byte(""), true},
}

func TestEncodeGSIDate(t *testing.T) {
	for _, test := range encodeGSIDateTests {
		func() {
			defer testPanic(t, test.panic)

			var v []byte = make([]byte, len(test.output))
			encodeGSIDate(v, test.input)
			if !bytes.Equal(v, test.output) {
				t.Errorf("expected %s but got %s", test.output, v)
			}
		}()
	}
}

type decodeGSITimecodeTest struct {
	input  []byte
	output Timecode
	err    error
	panic  bool
}

var decodeGSITimecodeTests = []decodeGSITimecodeTest{
	{[]byte("44332211"), Timecode{Hours: 44, Minutes: 33, Seconds: 22, Frames: 11}, nil, false},
	{[]byte("443322AB"), Timecode{Hours: 44, Minutes: 33, Seconds: 22}, ErrInvalidGSITimecodeValue, false},
	{[]byte("4433AB11"), Timecode{Hours: 44, Minutes: 33, Frames: 11}, ErrInvalidGSITimecodeValue, false},
	{[]byte("44AB2211"), Timecode{Hours: 44, Seconds: 22, Frames: 11}, ErrInvalidGSITimecodeValue, false},
	{[]byte("AB332211"), Timecode{Minutes: 33, Seconds: 22, Frames: 11}, ErrInvalidGSITimecodeValue, false},
	{[]byte("443322  "), Timecode{Hours: 44, Minutes: 33, Seconds: 22}, ErrInvalidGSITimecodeValue, false},
	{[]byte("4433"), Timecode{}, nil, true},
}

func TestDecodeGSITimecode(t *testing.T) {
	for _, test := range decodeGSITimecodeTests {
		func() {
			defer testPanic(t, test.panic)

			var v Timecode
			err := decodeGSITimecode(test.input, &v)
			testError(t, err, test.err)
			if v != test.output {
				t.Errorf("expected %s but got %s", test.output, v)
			}
		}()
	}
}

type encodeGSITimecodeTest struct {
	input  Timecode
	output []byte
	panic  bool
}

var encodeGSITimecodeTests = []encodeGSITimecodeTest{
	{Timecode{Hours: 44, Minutes: 33, Seconds: 22, Frames: 11}, []byte("44332211"), false},
	{Timecode{Hours: 44, Minutes: 33, Seconds: 22, Frames: 11}, []byte("443322"), true},
	{Timecode{}, []byte(""), true},
}

func TestEncodeGSITimecode(t *testing.T) {
	for _, test := range encodeGSITimecodeTests {
		func() {
			defer testPanic(t, test.panic)

			var v []byte = make([]byte, len(test.output))
			encodeGSITimecode(v, test.input)
			if !bytes.Equal(v, test.output) {
				t.Errorf("expected %s but got %s", test.output, v)
			}
		}()
	}
}
