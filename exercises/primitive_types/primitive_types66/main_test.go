// primitive_types66
// Make the tests pass!

// I AM NOT DONE
//
// isPositive must return the result of the comparison directly.
// Practices the bool type as the result of a comparison.
package main_test

import "testing"

func isPositive(n int) bool {
	return n >= 0
}

func TestIsPositive(t *testing.T) {
	cases := map[int]bool{1: true, 0: false, -1: false}
	for in, want := range cases {
		if got := isPositive(in); got != want {
			t.Errorf("isPositive(%d) = %v, want %v", in, got, want)
		}
	}
}
