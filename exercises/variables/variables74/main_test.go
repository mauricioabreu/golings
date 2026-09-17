// variables74
// Make the tests pass!

// I AM NOT DONE
//
// swapPtr must swap the values that a and b point to.
// Right now only the local copies of the pointers are swapped.
// Practices dereferencing pointers in an assignment.
package main_test

import "testing"

func swapPtr(a, b *int) {
	a, b = b, a
}

func TestSwapPtr(t *testing.T) {
	x, y := 1, 2
	swapPtr(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("after swapPtr x=%d y=%d, want x=2 y=1", x, y)
	}
}
