// if46
// Make the tests pass!

// I AM NOT DONE
//
// median3 returns the middle of three numbers without sorting them.
// Practices combining conditions with && and ||.
package main_test

import "testing"

func median3(a, b, c int) int {
	if a > b && a < c {
		return a
	}
	if b > a && b < c {
		return b
	}
	return c
}

func TestMedian3(t *testing.T) {
	cases := [][4]int{{1, 2, 3, 2}, {3, 2, 1, 2}, {2, 3, 1, 2}, {5, 5, 1, 5}, {7, 1, 4, 4}, {9, 3, 5, 5}}
	for _, c := range cases {
		if got := median3(c[0], c[1], c[2]); got != c[3] {
			t.Errorf("median3(%d,%d,%d) = %d, want %d", c[0], c[1], c[2], got, c[3])
		}
	}
}
