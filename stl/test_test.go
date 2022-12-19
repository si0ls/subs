package stl

import (
	"errors"
	"testing"
)

func testPanic(t *testing.T, shouldPanic bool) {
	t.Helper()

	if r := recover(); r != nil && !shouldPanic {
		t.Errorf("unexpected panic: %v", r)
	} else if r == nil && shouldPanic {
		t.Errorf("expected panic")
	}
}

func testError(t *testing.T, err error, expectedErr error) bool {
	t.Helper()

	if err != nil {
		err = errors.Unwrap(err)
		if expectedErr == nil {
			t.Errorf("unexpected error: %v", err)
			return false
		} else if err != expectedErr {
			t.Errorf("expected error %v but got %v", expectedErr, err)
			return false
		}
	}
	return true
}
