// switch7
// Make the tests pass!

// I AM NOT DONE
//
// sign returns -1, 0 or 1.
// Practices a tagless switch with three branches.
package main_test

import "testing"

func sign(x float64) int {
	switch {
	case x > 0:
		return 1
	case x <= 0:
		return -1
	}
	return 0
}

func TestSign(t *testing.T) {
	cases := map[float64]int{2.5: 1, 0: 0, -0.1: -1}
	for in, want := range cases {
		if got := sign(in); got != want {
			t.Errorf("sign(%v) = %d, want %d", in, got, want)
		}
	}
}
