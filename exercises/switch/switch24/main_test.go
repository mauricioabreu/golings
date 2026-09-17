// switch24
// Make the tests pass!

// I AM NOT DONE
//
// calc returns the result, or the errors "division by zero" and "unknown operator".
// Practices a switch with checks inside the branches and a default.
package main_test

import (
	"errors"
	"testing"
)

func calc(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "/":
		return a / b, nil
	}
	return 0, nil
}

func TestCalc(t *testing.T) {
	_ = errors.New
	if v, err := calc(8, 2, "/"); err != nil || v != 4 {
		t.Errorf("calc(8/2) = %d, %v", v, err)
	}
	if _, err := calc(1, 0, "/"); err == nil || err.Error() != "division by zero" {
		t.Errorf("calc(1/0) err = %v", err)
	}
	if _, err := calc(1, 2, "^"); err == nil || err.Error() != "unknown operator" {
		t.Errorf("calc(1^2) err = %v", err)
	}
}
