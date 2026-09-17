// if86
// Make the tests pass!

// I AM NOT DONE
//
// abs must return the absolute value of an integer.
// Practices changing a value conditionally.
package main_test

import "testing"

func abs(x int) int {
	if x > 0 {
		return -x
	}
	return x
}

func TestAbs(t *testing.T) {
	cases := map[int]int{5: 5, -5: 5, 0: 0}
	for in, want := range cases {
		if got := abs(in); got != want {
			t.Errorf("abs(%d) = %d, want %d", in, got, want)
		}
	}
}
