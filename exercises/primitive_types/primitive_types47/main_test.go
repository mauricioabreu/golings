// primitive_types47
// Make the tests pass!

// I AM NOT DONE
//
// floorDiv must divide rounding towards minus infinity: -7/2 = -4.
// Integer division in Go rounds towards zero.
package main_test

import "testing"

func floorDiv(a, b int) int {
	q := a / b
	return q
}

func TestFloorDiv(t *testing.T) {
	cases := [][3]int{{7, 2, 3}, {-7, 2, -4}, {7, -2, -4}, {-7, -2, 3}, {-8, 2, -4}, {0, 5, 0}}
	for _, c := range cases {
		if got := floorDiv(c[0], c[1]); got != c[2] {
			t.Errorf("floorDiv(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
