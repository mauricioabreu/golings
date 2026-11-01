// interfaces5
// Make the tests pass!

package main_test

import "testing"

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// validate returns an error if the age is negative, and nil if it is fine.
func validate(age int) error {
	if age < 0 {
		return &ValidationError{Message: "age can't be negative"}
	}
	return nil
}

func TestValidateAcceptsAValidAge(t *testing.T) {
	if err := validate(30); err != nil {
		t.Errorf("validate(30) returned an error, want nil: %v", err)
	}
}

func TestValidateRejectsANegativeAge(t *testing.T) {
	err := validate(-1)

	if err == nil {
		t.Fatal("validate(-1) returned nil, want an error")
	}
	if err.Error() != "age can't be negative" {
		t.Errorf("got %q", err.Error())
	}
}
