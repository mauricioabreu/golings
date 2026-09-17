// anonymous_functions78
// Make the tests pass!

// I AM NOT DONE
//
// area is computed by a literal called right away with its arguments.
// Practices passing arguments in an immediate call.
package main_test

import "testing"

func area() int {
	return func(w, h int) int {
		return w * h
	}(3, 3)
}

func TestArea(t *testing.T) {
	if area() != 12 {
		t.Errorf("area = %d, want 12", area())
	}
}
