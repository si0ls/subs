package stl

import (
	"golang.org/x/text/encoding/charmap"
)

var codePageNumberEncoders = map[CodePageNumber]TextEncoder{
	CodePageNumberUnitedStates:   &Charmap{charmap.CodePage437},
	CodePageNumberMultiLingual:   &Charmap{charmap.CodePage850},
	CodePageNumberPortugal:       &Charmap{charmap.CodePage860},
	CodePageNumberCanadianFrench: &Charmap{charmap.CodePage863},
	CodePageNumberNordic:         &Charmap{charmap.CodePage865},
}

var codePageNumberDecoders = map[CodePageNumber]TextDecoder{
	CodePageNumberUnitedStates:   &Charmap{charmap.CodePage437},
	CodePageNumberMultiLingual:   &Charmap{charmap.CodePage850},
	CodePageNumberPortugal:       &Charmap{charmap.CodePage860},
	CodePageNumberCanadianFrench: &Charmap{charmap.CodePage863},
	CodePageNumberNordic:         &Charmap{charmap.CodePage865},
}

var characterCodeTableEncoders = map[CharacterCodeTable]TextEncoder{
	CharacterCodeTableLatin:         &ISO6937,
	CharacterCodeTableLatinCyrillic: &Charmap{charmap.ISO8859_5},
	CharacterCodeTableLatinArabic:   &Charmap{charmap.ISO8859_6},
	CharacterCodeTableLatinGreek:    &Charmap{charmap.ISO8859_7},
	CharacterCodeTableLatinHebrew:   &Charmap{charmap.ISO8859_8},
}

var characterCodeTableDecoders = map[CharacterCodeTable]TextDecoder{
	CharacterCodeTableLatin:         &ISO6937,
	CharacterCodeTableLatinCyrillic: &Charmap{charmap.ISO8859_5},
	CharacterCodeTableLatinArabic:   &Charmap{charmap.ISO8859_6},
	CharacterCodeTableLatinGreek:    &Charmap{charmap.ISO8859_7},
	CharacterCodeTableLatinHebrew:   &Charmap{charmap.ISO8859_8},
}

// TextDecoder is a decoder for text.
type TextDecoder interface {
	// Decode decodes a X-encoded byte slice to a UTF-8 byte slice.
	Decode(b []byte) ([]byte, error)
}

// TextEncoder is an encoder for text.
type TextEncoder interface {
	// Encode decodes a UTF-8 byte slice to a X-encoded byte slice.
	Encode(b []byte) ([]byte, error)
}

// Charmap is a wrapper around golang.org/x/text/encoding/charmap.Charmap
// It implements the TextDecoder and TextEncoder interfaces.
type Charmap struct {
	codePage *charmap.Charmap
}

// Encode encodes b using the charmap.
func (c *Charmap) Encode(b []byte) ([]byte, error) {
	return c.codePage.NewEncoder().Bytes(b)
}

// Decode decodes b using the charmap.
func (c *Charmap) Decode(b []byte) ([]byte, error) {
	return c.codePage.NewDecoder().Bytes(b)
}
