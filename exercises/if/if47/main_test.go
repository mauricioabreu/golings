// if47
// Make the tests pass!

// I AM NOT DONE
//
// shipping must return 0 for an order of 3000 or more, and 300 otherwise.
// Practices a simple branch on a threshold.
package main_test

import "testing"

func shipping(total int) int {
	if total >= 3000 {
		return 300
	}
	return 300
}

func TestShipping(t *testing.T) {
	cases := map[int]int{2999: 300, 3000: 0, 10000: 0, 0: 300}
	for in, want := range cases {
		if got := shipping(in); got != want {
			t.Errorf("shipping(%d) = %d, want %d", in, got, want)
		}
	}
}
