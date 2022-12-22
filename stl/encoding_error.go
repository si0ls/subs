package stl

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidGSIIntValue      = errors.New("invalid GSI int value")
	ErrEmptyGSIIntValue        = errors.New("empty GSI int value")
	ErrInvalidGSIByteValue     = errors.New("invalid GSI byte value")
	ErrEmptyGSIByteValue       = errors.New("empty GSI byte value")
	ErrInvalidGSIHexValue      = errors.New("invalid GSI hex value")
	ErrEmptyGSIHexValue        = errors.New("empty GSI hex value")
	ErrInvalidGSIStringValue   = errors.New("invalid GSI string value")
	ErrEmptyGSIStringValue     = errors.New("empty GSI string value")
	ErrUnsupportedGSICodePage  = errors.New("unsupported code page")
	ErrInvalidGSIDateValue     = errors.New("invalid GSI date value")
	ErrEmptyGSIDateValue       = errors.New("empty GSI date value")
	ErrInvalidGSITimecodeValue = errors.New("invalid GSI timecode value")
	ErrEmptyGSITimecodeValue   = errors.New("empty GSI timecode value")
)

// EncodingError is an error that occurred during encoding or decoding.
// It carries the input buffer that caused the error.
// EncodingError implements the error and Unwrap interfaces.
type EncodingError struct {
	error
	value interface{}
}

func encodingErr(err error, value interface{}) error {
	return &EncodingError{error: err, value: value}
}

// Error returns the error message.
func (e *EncodingError) Error() string {
	return fmt.Sprintf("%s (input: %v)", e.error.Error(), e.value)
}

// Unwrap returns the underlying error.
func (e *EncodingError) Unwrap() error {
	return e.error
}

// Input returns the input buffer.
func (e *EncodingError) Input() interface{} {
	return e.value
}

// DecodeError is an error that occurred during decoding.
// It extends EncodingError that carries the value that caused the error.
type DecodeError struct {
	EncodingError
}

func decodeErr(err error, value interface{}) error {
	return &DecodeError{EncodingError: EncodingError{error: err, value: value}}
}

// Error returns the error message.
func (e *DecodeError) Error() string {
	return fmt.Sprintf("decode: %s", e.error.Error())
}

// Unwrap returns the underlying error.
func (e *DecodeError) Unwrap() error {
	return e.error
}

// EncodeError is an error that occurred during encoding.
// It extends EncodingError that carries the input buffer that caused the error.
type EncodeError struct {
	EncodingError
}

func encodeErr(err error, value interface{}) error {
	return &EncodeError{EncodingError: EncodingError{error: err, value: value}}
}

// Error returns the error message.
func (e *EncodeError) Error() string {
	return fmt.Sprintf("encode: %s", e.error.Error())
}

// Unwrap returns the underlying error.
func (e *EncodeError) Unwrap() error {
	return e.error
}
