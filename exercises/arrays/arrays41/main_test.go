// arrays41
// Make the tests pass!

// I AM NOT DONE
//
// minOf returns the smallest element of an array.
// Practices a comparison in a loop over an array.
package main_test

import "testing"

func minOf(a [5]int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func TestMinOf(t *testing.T) {
	if got := minOf([5]int{4, 8, -1, 3, 0}); got != -1 {
		t.Errorf("minOf = %d, want -1", got)
	}
}
