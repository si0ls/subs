package stl

import "fmt"

// GSIError is an error that occurred on a GSI field.
// It extends FieldError that carries the concerned GSI field.
type GSIError struct {
	error
	FieldError
	field GSIField
}

func gsiErr(err error, field GSIField) error {
	if err == nil {
		return nil
	}
	return &GSIError{error: err, field: field}
}

// Error returns the error message.
func (e *GSIError) Error() string {
	return fmt.Sprintf("GSI %s: %s", e.field, e.error.Error())
}

// Unwrap returns the underlying error.
func (e *GSIError) Unwrap() error {
	return e.error
}

// Field returns the concerned GSI field.
func (e *GSIError) Field() GSIField {
	return e.field
}
