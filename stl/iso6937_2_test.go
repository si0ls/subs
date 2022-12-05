package stl

import (
	"bytes"
	"testing"
)

type iso6937Decode struct {
	v             []byte
	expected      []byte
	expectedError bool
}

var iso6937DecodeTests = []iso6937Decode{
	{[]byte("abc"), []byte("abc"), false},
	{[]byte{0xC1, 'e'}, []byte("e\u0300"), false},
	{[]byte{0xC1}, nil, true},
	{[]byte{0x00}, []byte{0x00}, false},
	{[]byte{0x7F}, []byte("\ufffd"), false},
}

func TestIso6937Decode(t *testing.T) {
	decoder := iso6937{}

	for _, test := range iso6937DecodeTests {
		tmp, err := decoder.Decode(test.v)
		if err != nil && !test.expectedError {
			t.Errorf("Decode(%q) unexpected error: %s", test.v, err)
		}

		if !bytes.Equal(tmp, test.expected) {
			t.Errorf("Decode(%q) = %0 2x, want %0 2x", test.v, tmp, test.expected)
		}
	}
}

type iso6937Encode struct {
	v             []byte
	expected      []byte
	expectedError bool
}

var iso6937EncodeTests = []iso6937Encode{
	{[]byte("abc"), []byte("abc"), false},
	{[]byte("e\u0300"), []byte{0xC1, 'e'}, false},
	{[]byte("\u0300"), []byte{}, true},
	{[]byte{0x00}, []byte{0x00}, false},
	{[]byte("\ufffd"), []byte{0x3F}, false},
	{[]byte("\u7898"), []byte{0x3F}, false},
}

func TestIso6937Encode(t *testing.T) {
	encoder := iso6937{}

	for _, test := range iso6937EncodeTests {
		tmp, err := encoder.Encode(test.v)
		if err != nil && !test.expectedError {
			t.Errorf("Encode(%q) unexpected error: %s", test.v, err)
		}

		if !bytes.Equal(tmp, test.expected) {
			t.Errorf("Encode(%q) = %q, want %q", test.v, tmp, test.expected)
		}
	}
}
