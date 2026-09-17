// variables104
// Make the tests pass!

// I AM NOT DONE
//
// clamp must confine v to the range [lo, hi], even when the bounds
// are passed the other way round.
// Practices swapping values and successive checks.
package main_test

import "testing"

func clamp(v, lo, hi int) int {
	if v < lo {
		return hi
	}
	if v > hi {
		return hi
	}
	return v
}

func TestClamp(t *testing.T) {
	cases := []struct{ v, lo, hi, want int }{
		{5, 0, 10, 5}, {-1, 0, 10, 0}, {11, 0, 10, 10}, {11, 10, 0, 10}, {-3, 10, 0, 0},
	}
	for _, c := range cases {
		if got := clamp(c.v, c.lo, c.hi); got != c.want {
			t.Errorf("clamp(%d, %d, %d) = %d, want %d", c.v, c.lo, c.hi, got, c.want)
		}
	}
}
