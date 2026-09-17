// variables12
// Make the tests pass!

// I AM NOT DONE
//
// validate returns an error and must return nil when nothing is wrong.
// The result is never nil, though.
// An interface holding a nil pointer is not itself nil.
package main_test

import "testing"

type ValidationError struct{ Field string }

func (e *ValidationError) Error() string { return "invalid " + e.Field }

func validate(name string) error {
	var verr *ValidationError
	if name == "" {
		verr = &ValidationError{Field: "name"}
	}
	return verr
}

func TestValidate(t *testing.T) {
	if err := validate("gopher"); err != nil {
		t.Errorf("validate(gopher) = %v, want nil", err)
	}
	if err := validate(""); err == nil {
		t.Errorf("validate(\"\") should return an error")
	}
}
