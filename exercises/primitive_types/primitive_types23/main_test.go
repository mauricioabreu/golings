// primitive_types23
// Make the tests pass!

// I AM NOT DONE
//
// boolToInt must return 1 for true and 0 for false.
// The code does not compile: a bool cannot be converted to an int.
// Go has no implicit conversion for a bool.
package main_test

import "testing"

func boolToInt(b bool) int {
	return int(b)
}

func TestBoolToInt(t *testing.T) {
	if boolToInt(true) != 1 || boolToInt(false) != 0 {
		t.Errorf("boolToInt is wrong")
	}
}
