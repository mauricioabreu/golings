// anonymous_functions71
// Make the tests pass!

// I AM NOT DONE
//
// minMax calls a literal returning the minimum and the maximum of a pair.
// Practices a function literal with several results.
package main_test

import "testing"

func minMax(a, b int) (int, int) {
	order := func(x, y int) (int, int) {
		if x > y {
			return x, y
		}
		return x, y
	}
	return order(a, b)
}

func TestMinMax(t *testing.T) {
	if lo, hi := minMax(9, 2); lo != 2 || hi != 9 {
		t.Errorf("minMax = %d, %d", lo, hi)
	}
}
