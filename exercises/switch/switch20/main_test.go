// switch20
// Make the tests pass!

// I AM NOT DONE
//
// isOne must return true when the value is one of any of the integer types
// int, int64 or uint8.
// In a switch over an interface{}, case 1 matches the type (int) as well as the value.
package main_test

import "testing"

func isOne(v interface{}) bool {
	switch v {
	case 1:
		return true
	}
	return false
}

func TestIsOne(t *testing.T) {
	for _, v := range []interface{}{1, int64(1), uint8(1)} {
		if !isOne(v) {
			t.Errorf("isOne(%T(1)) = false", v)
		}
	}
	for _, v := range []interface{}{2, "1", 1.0} {
		if isOne(v) {
			t.Errorf("isOne(%T(%v)) = true", v, v)
		}
	}
}
