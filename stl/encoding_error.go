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
	input []byte
}

func encodingErr(err error, input []byte) *EncodingError {
	return &EncodingError{err, input}
}

// Error returns the error message.
func (e *EncodingError) Error() string {
	return fmt.Sprintf("%s (input: %v)", e.error, e.input)
}

// Unwrap returns the underlying error.
func (e *EncodingError) Unwrap() error {
	return e.error
}

// Input returns the input buffer.
func (e *EncodingError) Input() []byte {
	return e.input
}

// GSIEncodingError is an error that occurred during GSI encoding.
// It extends EncodingError that carries the input buffer that caused the error.
// It carries the concerned GSI field.
type GSIEncodingError struct {
	error
	field Field
}

// Error returns the error message.
func (e *GSIEncodingError) Error() string {
	return fmt.Sprintf("GSI %s: %s", e.field, e.error.Error())
}

// Unwrap returns the underlying error.
func (e *GSIEncodingError) Unwrap() error {
	return e.error
}

// Field returns the concerned GSI field.
func (e *GSIEncodingError) Field() Field {
	return e.field
}

// wrapGSIEncodingErr wraps an error in a GSIEncodingError.
func wrapGSIEncodingErr(err error, field Field) *GSIEncodingError {
	if err == nil {
		return nil
	}
	return &GSIEncodingError{err, field}
}

// TTIEncodingError is an error that occurred during TTI encoding.
// It extends EncodingError that carries the input buffer that caused the error.
// It carries the concerned TTI field and the TTI block number.
// If TTI block is -1, it means that the TTI block number is unknown.
type TTIEncodingError struct {
	error
	field Field
	block int
}

// Error returns the error message.
// If TTI block is -1, it means that the TTI block number is unknown.
func (e *TTIEncodingError) Error() string {
	return fmt.Sprintf("TTI block %d %s: %s", e.block, e.field, e.error.Error())
}

// Unwrap returns the underlying error.
func (e *TTIEncodingError) Unwrap() error {
	return e.error
}

// Field returns the concerned TTI field.
func (e *TTIEncodingError) Field() Field {
	return e.field
}

// Block returns the concerned TTI block number.
// If TTI block is -1, it means that the TTI block number is unknown.
func (e *TTIEncodingError) Block() int {
	return e.block
}

// wrapTTIEncodingErr wraps an error in a TTIEncodingError.
func wrapTTIEncodingErr(err error, field Field, block int) *TTIEncodingError {
	if err == nil {
		return nil
	}
	return &TTIEncodingError{err, field, block}
}

// setTTIEncodingErrBlock sets the TTI block number in a TTIEncodingError.
func setTTIEncodingErrBlock(err *TTIEncodingError, block int) *TTIEncodingError {
	if err == nil {
		return nil
	}
	err.block = block
	return err
}
