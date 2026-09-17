// switch75
// Make the tests pass!

// I AM NOT DONE
//
// retryable returns true for ErrTimeout and ErrBusy, even when they are wrapped.
// A switch on the error value does not see through the wrappers.
// A case compares errors with ==.
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

var (
	ErrTimeout = errors.New("timeout")
	ErrBusy    = errors.New("busy")
)

func retryable(err error) bool {
	switch err {
	case ErrTimeout, ErrBusy:
		return true
	}
	return false
}

func TestRetryable(t *testing.T) {
	if !retryable(ErrBusy) || !retryable(fmt.Errorf("db: %w", ErrTimeout)) {
		t.Errorf("timeout and busy errors are retryable")
	}
	if retryable(errors.New("timeout")) || retryable(nil) {
		t.Errorf("other errors are not retryable")
	}
}
