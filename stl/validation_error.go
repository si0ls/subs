package stl

import (
	"errors"
	"fmt"
)

// ValidateError is an error that occurred during validation.
// It carries the value that caused the error.
// ValidateError implements the error and Unwrap interfaces.
type ValidateError struct {
	error
	value any
	fatal bool
}

func validateErr(err error, value any, fatal bool) error {
	if err == nil {
		return nil
	}
	return &ValidateError{error: err, value: value, fatal: fatal}
}

// Error returns the error message.
func (e *ValidateError) Error() string {
	if e.fatal {
		return fmt.Sprintf("fatal: %s (value: %v)", e.error.Error(), e.value)
	}
	return fmt.Sprintf("%s (value: %v)", e.error.Error(), e.value)
}

// Unwrap returns the underlying error.
func (e *ValidateError) Unwrap() error {
	return e.error
}

// Value returns the value that caused the error.
func (e *ValidateError) Value() any {
	return e.value
}

// IsFatal returns true if the error is fatal.
func (e *ValidateError) IsFatal() bool {
	return e.fatal
}

var ErrNoTTIBlocks = errors.New("no TTI blocks")
