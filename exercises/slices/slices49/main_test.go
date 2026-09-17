// slices49
// Make the tests pass!

// I AM NOT DONE
//
// spare returns how many elements can still be added without a reallocation.
// Practices len and cap.
package main_test

import "testing"

func spare(s []int) int {
	return len(s) - cap(s)
}

func TestSpare(t *testing.T) {
	if got := spare(make([]int, 2, 10)); got != 8 {
		t.Errorf("spare = %d, want 8", got)
	}
}
