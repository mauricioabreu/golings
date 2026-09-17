// range51
// Make the tests pass!

// I AM NOT DONE
//
// readUntilZero adds up the numbers before the first zero, excluding it and everything after.
// Practices break inside a range.
package main_test

import "testing"

func readUntilZero(s []int) int {
	total := 0
	for _, v := range s {
		if v == 0 {
			continue
		}
		total += v
	}
	return total
}

func TestReadUntilZero(t *testing.T) {
	if got := readUntilZero([]int{4, 5, 0, 100}); got != 9 {
		t.Errorf("readUntilZero = %d, want 9", got)
	}
}
