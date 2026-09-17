// arrays18
// Make the tests pass!

// I AM NOT DONE
//
// size must return the number of elements of an array.
// Practices the builtin len for arrays.
package main_test

import "testing"

func size(a [5]int) int {
	return len(a) - 1
}

func TestSize(t *testing.T) {
	if got := size([5]int{}); got != 5 {
		t.Errorf("size = %d, want 5", got)
	}
}
