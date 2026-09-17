// variables51
// Make the tests pass!

// I AM NOT DONE
//
// rotate must shift three values to the left: (a, b, c) -> (b, c, a).
// Practices assigning three variables at once.
package main_test

import "testing"

func rotate(a, b, c int) (int, int, int) {
	a = b
	b = c
	c = a
	return a, b, c
}

func TestRotate(t *testing.T) {
	a, b, c := rotate(1, 2, 3)
	if a != 2 || b != 3 || c != 1 {
		t.Errorf("rotate(1,2,3) = (%d,%d,%d), want (2,3,1)", a, b, c)
	}
}
