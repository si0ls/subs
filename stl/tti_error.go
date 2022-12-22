package stl

import "fmt"

// TTIError is an error that occurred on a TTI field.
// It extends FieldError that carries the concerned TTI field.
// It carries the concerned TTI block number.
// If TTI block number is -1, it means that the TTI block number is unknown.
type TTIError struct {
	error
	FieldError
	field       TTIField
	blockNumber int
}

func ttiErr(err error, field TTIField) error {
	if err == nil {
		return nil
	}
	return &TTIError{error: err, field: field}
}

func ttiErrWithBlockNumber(err error, field TTIField, blockNumber int) error {
	if err == nil {
		return nil
	}
	return &TTIError{error: err, field: field, blockNumber: blockNumber}
}

// Error returns the error message.
func (e *TTIError) Error() string {
	return fmt.Sprintf("GSI %s: %s", e.field, e.error.Error())
}

// Unwrap returns the underlying error.
func (e *TTIError) Unwrap() error {
	return e.error
}

// Field returns the concerned TTI field.
func (e *TTIError) Field() TTIField {
	return e.field
}

// BlockNumber returns the concerned TTI block number.
// If TTI block number is -1, it means that the TTI block number is unknown.
func (e *TTIError) BlockNumber() int {
	return e.blockNumber
}

func (e *TTIError) setBlockNumber(blockNumber int) {
	e.blockNumber = blockNumber
}

func setTTIErrsBlockNumber(errs []error, blockNumber int) {
	for _, err := range errs {
		if ttiErr, ok := err.(*TTIError); ok {
			ttiErr.setBlockNumber(blockNumber)
		}
	}
}
