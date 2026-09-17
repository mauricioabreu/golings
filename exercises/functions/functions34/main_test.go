// functions34
// Make the tests pass!

// I AM NOT DONE
//
// minMax must return the minimum first and the maximum second.
// Practices returning several values.
package main_test

import "testing"

func minMax(a, b int) (int, int) {
	if a > b {
		a, b = b, a
	}
	return b, a
}

func TestMinMax(t *testing.T) {
	lo, hi := minMax(9, 4)
	if lo != 4 || hi != 9 {
		t.Errorf("minMax(9, 4) = %d, %d; want 4, 9", lo, hi)
	}
}
