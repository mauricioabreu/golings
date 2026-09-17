// range33
// Make the tests pass!

// I AM NOT DONE
//
// minIndex returns the index of the smallest element.
// Practices remembering an index in a range.
package main_test

import "testing"

func minIndex(s []int) int {
	best := 0
	for i, v := range s {
		if v < s[best] {
			best = v
		}
	}
	return best
}

func TestMinIndex(t *testing.T) {
	if got := minIndex([]int{4, 2, 8, 1, 9}); got != 3 {
		t.Errorf("minIndex = %d, want 3", got)
	}
}
