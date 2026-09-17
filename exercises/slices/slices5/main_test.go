// slices5
// Make the tests pass!

// I AM NOT DONE
//
// reset empties a buffer while keeping the memory it holds for reuse.
// Practices s[:0].
package main_test

import "testing"

func reset(s []int) []int {
	return nil
}

func TestReset(t *testing.T) {
	s := reset(make([]int, 5, 8))
	if len(s) != 0 || cap(s) != 8 {
		t.Errorf("reset: len=%d cap=%d, want 0 and 8", len(s), cap(s))
	}
}
