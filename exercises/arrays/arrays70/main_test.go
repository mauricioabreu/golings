// arrays70
// Make the tests pass!

// I AM NOT DONE
//
// sameValues compares two arrays of pointers by the values they point at.
// == on a [N]*T compares addresses.
package main_test

import "testing"

func sameValues(a, b [2]*int) bool {
	return a == b
}

func TestSameValues(t *testing.T) {
	x1, y1, x2, y2 := 1, 2, 1, 2
	if !sameValues([2]*int{&x1, &y1}, [2]*int{&x2, &y2}) {
		t.Errorf("equal values should match")
	}
	z := 3
	if sameValues([2]*int{&x1, &y1}, [2]*int{&x1, &z}) {
		t.Errorf("different values should not match")
	}
}
