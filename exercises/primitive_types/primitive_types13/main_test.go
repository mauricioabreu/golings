// primitive_types13
// Make the tests pass!

// I AM NOT DONE
//
// shiftBy shifts a number left by n bits, or right when n is negative.
// Right now a negative n panics.
// A shift by a negative count is a run time panic.
package main_test

import "testing"

func shiftBy(x int, n int) int {
	return x << n
}

func TestShiftBy(t *testing.T) {
	cases := [][3]int{{1, 3, 8}, {16, -2, 4}, {5, 0, 5}}
	for _, c := range cases {
		if got := shiftBy(c[0], c[1]); got != c[2] {
			t.Errorf("shiftBy(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
