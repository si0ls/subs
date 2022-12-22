package stl

import "fmt"

// Filed represents a block field.
type Field string

const FieldUnknown Field = "<unknown>" // Unknown field

// FieldError is an error that occurred on a field.
type FieldError struct {
	error
	field Field
}

// Error returns the error message.
func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.field, e.error.Error())
}

// Unwrap returns the underlying error.
func (e *FieldError) Unwrap() error {
	return e.error
}

// Field returns the concerned field.
func (e *FieldError) Field() Field {
	return e.field
}
