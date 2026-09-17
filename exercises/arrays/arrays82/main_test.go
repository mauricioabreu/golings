// arrays82
// Make the tests pass!

// I AM NOT DONE
//
// minMax returns the minimum and the maximum of an array in a single pass.
// Practices two accumulators over an array.
package main_test

import "testing"

func minMax(a [6]int) (int, int) {
	lo, hi := 0, 0
	for _, v := range a {
		if v < lo {
			lo = v
		} else if v > lo {
			hi = v
		}
	}
	return lo, hi
}

func TestMinMax(t *testing.T) {
	lo, hi := minMax([6]int{5, 8, 3, 9, 1, 4})
	if lo != 1 || hi != 9 {
		t.Errorf("minMax = %d, %d; want 1, 9", lo, hi)
	}
	lo, hi = minMax([6]int{7, 7, 8, 8, 9, 10})
	if lo != 7 || hi != 10 {
		t.Errorf("minMax = %d, %d; want 7, 10", lo, hi)
	}
}
