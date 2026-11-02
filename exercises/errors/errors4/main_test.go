// errors4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Reason
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Reason: "can't be negative"}
	}
	return nil
}

func register(age int) error {
	if err := validateAge(age); err != nil {
		return fmt.Errorf("registering: %w", err)
	}
	return nil
}

// failedField returns the field that failed validation, looking through any
// wrapping, or "" if the error is not a ValidationError.
func failedField(err error) string {
	// TODO: Look for a *ValidationError in the chain of err.
	return ""
}

func TestFailedField(t *testing.T) {
	if got := failedField(register(-1)); got != "age" {
		t.Errorf("failedField(register(-1)) = %q, want %q", got, "age")
	}
}

func TestFailedFieldWithoutAValidationError(t *testing.T) {
	if got := failedField(errors.New("disk full")); got != "" {
		t.Errorf("failedField(disk full) = %q, want an empty string", got)
	}
	if got := failedField(register(30)); got != "" {
		t.Errorf("failedField(nil) = %q, want an empty string", got)
	}
}
