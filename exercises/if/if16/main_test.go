// if16
// Make the tests pass!

// I AM NOT DONE
//
// isOdd must work for negative numbers too.
// For -3 it currently returns false.
// Practices the sign of the remainder operator in Go.
package main_test

import "testing"

func isOdd(n int) bool {
	if n%2 == 1 {
		return true
	}
	return false
}

func TestIsOdd(t *testing.T) {
	cases := map[int]bool{3: true, -3: true, 4: false, -4: false, 0: false}
	for in, want := range cases {
		if got := isOdd(in); got != want {
			t.Errorf("isOdd(%d) = %v, want %v", in, got, want)
		}
	}
}
