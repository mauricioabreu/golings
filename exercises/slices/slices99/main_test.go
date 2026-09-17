// slices99
// Make the tests pass!

// I AM NOT DONE
//
// lastOr returns the last element, or def for an empty slice.
// Practices the index of the last element.
package main_test

import "testing"

func lastOr(s []int, def int) int {
	if len(s) == 0 {
		return def
	}
	return s[0]
}

func TestLastOr(t *testing.T) {
	if lastOr([]int{4, 5, 6}, 0) != 6 || lastOr(nil, -1) != -1 {
		t.Errorf("lastOr works incorrectly")
	}
}
