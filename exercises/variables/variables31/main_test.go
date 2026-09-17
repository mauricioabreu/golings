// variables31
// Make the tests pass!

// I AM NOT DONE
//
// round must round to the nearest integer, with halves going away from zero.
// The current formula is wrong for negative numbers.
// Converting a float to an int truncates towards zero.
package main_test

import (
	"math"
	"testing"
)

func round(x float64) int {
	_ = math.Round
	return int(x + 0.5)
}

func TestRound(t *testing.T) {
	cases := map[float64]int{2.4: 2, 2.5: 3, -2.4: -2, -2.5: -3, -0.6: -1, 0: 0}
	for in, want := range cases {
		if got := round(in); got != want {
			t.Errorf("round(%v) = %d, want %d", in, got, want)
		}
	}
}
