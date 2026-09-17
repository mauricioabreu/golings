// arrays61
// Make the tests pass!

// I AM NOT DONE
//
// sum must add up every element of an array.
// Practices a loop over an array.
package main_test

import "testing"

func sum(a [4]int) int {
	total := 0
	for i := 1; i < len(a); i++ {
		total += a[i]
	}
	return total
}

func TestSum(t *testing.T) {
	if got := sum([4]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("sum = %d, want 10", got)
	}
}
