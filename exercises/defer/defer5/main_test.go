// defer5
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"strings"
	"testing"
)

// safeDivide divides a by b. Dividing an integer by zero panics: turn that
// panic into an error, so the caller gets an error and the program goes on.
func safeDivide(a, b int) (result int, err error) {
	// TODO: Recover from the panic in a deferred function, and set err.
	return a / b, nil
}

func TestSafeDivide(t *testing.T) {
	got, err := safeDivide(6, 3)

	if err != nil || got != 2 {
		t.Errorf("safeDivide(6, 3) = (%d, %v), want (2, nil)", got, err)
	}
}

func TestSafeDivideByZero(t *testing.T) {
	_, err := safeDivide(1, 0)

	if err == nil {
		t.Fatal("safeDivide(1, 0) returned no error")
	}
	if !strings.Contains(err.Error(), "divide by zero") {
		t.Errorf("the error should say what went wrong, got %q", err)
	}
}
