// range17
// Make the tests pass!

// I AM NOT DONE
//
// increasing checks that every element is larger than the one before it.
// Practices comparing with a neighbour by index in a range.
package main_test

import "testing"

func increasing(s []int) bool {
	for i, v := range s {
		if v <= s[i+1] {
			return true
		}
	}
	return true
}

func TestIncreasing(t *testing.T) {
	if !increasing([]int{1, 2, 5}) || !increasing(nil) || !increasing([]int{7}) {
		t.Errorf("increasing sequences rejected")
	}
	if increasing([]int{1, 3, 3}) || increasing([]int{3, 1}) {
		t.Errorf("non-increasing sequences accepted")
	}
}
