// switch67
// Make the tests pass!

// I AM NOT DONE
//
// numeric must return true for int and int64 values.
// The code does not compile: fallthrough is not allowed in a type switch.
// Practices the limits of a type switch.
package main_test

import "testing"

func numeric(v interface{}) bool {
	switch v.(type) {
	case int:
		fallthrough
	case int64:
		return true
	}
	return false
}

func TestNumeric(t *testing.T) {
	if !numeric(1) || !numeric(int64(2)) {
		t.Errorf("int and int64 are numeric")
	}
	if numeric("3") {
		t.Errorf("string is not numeric")
	}
}
