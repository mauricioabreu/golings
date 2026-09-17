// slices68
// Make the tests pass!

// I AM NOT DONE
//
// maxOf returns the largest element of a non-empty slice.
// Practices taking the starting value from the first element.
package main_test

import "testing"

func maxOf(s []int) int {
	m := 0
	for _, v := range s {
		if v > m {
			m = v
		}
	}
	return m
}

func TestMaxOf(t *testing.T) {
	if got := maxOf([]int{-7, -3, -9}); got != -3 {
		t.Errorf("maxOf = %d, want -3", got)
	}
}
