// range65
// Make the tests pass!

// I AM NOT DONE
//
// sumNonNegative adds up the non-negative numbers only.
// Practices continue inside a range.
package main_test

import "testing"

func sumNonNegative(s []int) int {
	total := 0
	for _, v := range s {
		if v >= 0 {
			continue
		}
		total += v
	}
	return total
}

func TestSumNonNegative(t *testing.T) {
	if got := sumNonNegative([]int{3, -2, 0, 5, -1}); got != 8 {
		t.Errorf("sumNonNegative = %d, want 8", got)
	}
}
