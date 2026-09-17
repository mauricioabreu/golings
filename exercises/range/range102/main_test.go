// range102
// Make the tests pass!

// I AM NOT DONE
//
// sum adds up the elements of a slice with a range.
// Practices the blank identifier for an index that is not needed.
package main_test

import "testing"

func sum(nums []int) int {
	total := 0
	for v := range nums {
		total += v
	}
	return total
}

func TestSum(t *testing.T) {
	if got := sum([]int{10, 20, 30}); got != 60 {
		t.Errorf("sum = %d, want 60", got)
	}
}
