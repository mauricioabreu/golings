// arrays86
// Make the tests pass!

// I AM NOT DONE
//
// maxOf returns the largest element of an array.
// Practices the starting value of a search for a maximum.
package main_test

import "testing"

func maxOf(a [4]int) int {
	m := 0
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

func TestMaxOf(t *testing.T) {
	if got := maxOf([4]int{-5, -2, -9, -3}); got != -2 {
		t.Errorf("maxOf = %d, want -2", got)
	}
}
