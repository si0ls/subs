package stl

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrNoTTIBlocks = errors.New("no TTI blocks")
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

func validateRange(value, min, max int, err error, fatal bool) error {
	if value < min || value > max {
		return validateErr(fmt.Errorf("%w: must be in range [%d;%d]", err, min, max), value, fatal)
	}
	return nil
}

func validateNotInRange(value, min, max int, err error, fatal bool) error {
	if value >= min && value <= max {
		return validateErr(fmt.Errorf("%w: must not be in range [%d;%d]", err, min, max), value, fatal)
	}
	return nil
}

func validateList[T comparable](value T, list []T, err error, fatal bool) error {
	for _, v := range list {
		if value == v {
			return nil
		}
	}
	return validateErr(fmt.Errorf("%w: must be one of %v", err, list), value, fatal)
}

func validateTimecode(tc Timecode, framerate uint, err error, fatal bool) error {
	if tcErr := tc.Validate(framerate); tcErr != nil {
		return validateErr(fmt.Errorf("%w: %s", err, tcErr), tc, fatal)
	}
	return nil
}

func validateTimecodeOrder(tc1, tc2 Timecode, framerate uint, err error, fatal bool) error {
	if tc1.ToDuration(framerate) > tc2.ToDuration(framerate) {
		return validateErr(fmt.Errorf("%w: %s > %s", err, tc1, tc2), tc1, fatal)
	}
	return nil
}

func validateTimecodeOrderStrict(tc1, tc2 Timecode, framerate uint, err error, fatal bool) error {
	if tc1.ToDuration(framerate) >= tc2.ToDuration(framerate) {
		return validateErr(fmt.Errorf("%w: %s > %s", err, tc1, tc2), tc1, fatal)
	}
	return nil
}

func validateNonEmptyString(s string, err error, fatal bool) error {
	if s == "" {
		return validateErr(err, s, fatal)
	}
	return nil
}

func validateDate(date time.Time, err error, fatal bool) error {
	if date.IsZero() {
		return validateErr(err, date, fatal)
	}
	return nil
}

func validateDateOrder(date1, date2 time.Time, err error, fatal bool) error {
	if date1.After(date2) {
		return validateErr(fmt.Errorf("%w: %s > %s", err, date1, date2), date1, fatal)
	}
	return nil
}
