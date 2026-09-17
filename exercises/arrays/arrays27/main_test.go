// arrays27
// Make the tests pass!

// I AM NOT DONE
//
// cols returns the number of columns of a 2 by 4 matrix.
// Practices len on a nested array.
package main_test

import "testing"

func cols(m [2][4]int) int {
	return len(m)
}

func TestCols(t *testing.T) {
	if got := cols([2][4]int{}); got != 4 {
		t.Errorf("cols = %d, want 4", got)
	}
}
