// errors5
// Make the tests pass!

package main_test

import (
	"errors"
	"testing"
)

var (
	ErrNameRequired = errors.New("name is required")
	ErrNegativeAge  = errors.New("age can't be negative")
)

// validateUser checks every field and reports all the problems at once, not
// only the first one. It returns nil when there are none.
func validateUser(name string, age int) error {
	var problems []error
	if name == "" {
		problems = append(problems, ErrNameRequired)
	}
	if age < 0 {
		problems = append(problems, ErrNegativeAge)
	}
	return errors.Join(problems...)
}

func TestValidUser(t *testing.T) {
	if err := validateUser("Ada", 36); err != nil {
		t.Errorf("got %v, want nil", err)
	}
}

func TestOneProblem(t *testing.T) {
	err := validateUser("", 36)

	if !errors.Is(err, ErrNameRequired) || errors.Is(err, ErrNegativeAge) {
		t.Errorf("got %v, want only %v", err, ErrNameRequired)
	}
}

func TestAllTheProblemsAtOnce(t *testing.T) {
	err := validateUser("", -1)

	if !errors.Is(err, ErrNameRequired) || !errors.Is(err, ErrNegativeAge) {
		t.Errorf("got %v, want both %q and %q", err, ErrNameRequired, ErrNegativeAge)
	}
}
