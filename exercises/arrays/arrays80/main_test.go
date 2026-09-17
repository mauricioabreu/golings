// arrays80
// Make the tests pass!

// I AM NOT DONE
//
// sameColor compares two colours held as [3]uint8 values.
// Arrays are comparable with ==.
package main_test

import "testing"

func sameColor(a, b [3]uint8) bool {
	return a != b
}

func TestSameColor(t *testing.T) {
	if !sameColor([3]uint8{1, 2, 3}, [3]uint8{1, 2, 3}) {
		t.Errorf("equal colors should match")
	}
	if sameColor([3]uint8{1, 2, 3}, [3]uint8{3, 2, 1}) {
		t.Errorf("different colors should not match")
	}
}
