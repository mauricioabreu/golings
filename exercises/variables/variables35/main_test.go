// variables35
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the sum of all positive numbers along with the running sums.
// The running sums grow, yet the total comes out as zero.
// Practices shadowing inside the body of a loop.
package main_test

import (
	"reflect"
	"testing"
)

func positiveSum(nums []int) (int, []int) {
	total := 0
	var steps []int
	for _, v := range nums {
		if v > 0 {
			total := total + v
			steps = append(steps, total)
		}
	}
	return total, steps
}

func TestPositiveSum(t *testing.T) {
	total, steps := positiveSum([]int{3, -1, 4, 5})
	if total != 12 {
		t.Errorf("total = %d, want 12", total)
	}
	if want := []int{3, 7, 12}; !reflect.DeepEqual(steps, want) {
		t.Errorf("steps = %v, want %v", steps, want)
	}
}
