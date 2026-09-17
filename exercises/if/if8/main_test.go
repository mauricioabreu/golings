// if8
// Make the tests pass!

// I AM NOT DONE
//
// max3 must return the largest of three numbers.
// Practices successive ifs updating a result.
package main_test

import "testing"

func max3(a, b, c int) int {
	m := a
	if b > m {
		m = b
	}
	if c > a {
		m = c
	}
	return m
}

func TestMax3(t *testing.T) {
	cases := [][4]int{{1, 2, 3, 3}, {3, 2, 1, 3}, {1, 5, 4, 5}, {-1, -2, -3, -1}, {2, 9, 3, 9}}
	for _, c := range cases {
		if got := max3(c[0], c[1], c[2]); got != c[3] {
			t.Errorf("max3(%d, %d, %d) = %d, want %d", c[0], c[1], c[2], got, c[3])
		}
	}
}
