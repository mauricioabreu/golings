// functions60
// Make the tests pass!

// I AM NOT DONE
//
// reduce(nums, init, f) folds the slice starting from init: f(f(init, a0), a1)...
// Practices an accumulator function.
package main_test

import "testing"

func reduce(nums []int, init int, f func(acc, x int) int) int {
	acc := 0
	for _, n := range nums {
		acc = f(n, acc)
	}
	return acc
}

func TestReduce(t *testing.T) {
	sum := reduce([]int{1, 2, 3}, 10, func(a, x int) int { return a + x })
	if sum != 16 {
		t.Errorf("sum = %d, want 16", sum)
	}
	digits := reduce([]int{1, 2, 3}, 0, func(a, x int) int { return a*10 + x })
	if digits != 123 {
		t.Errorf("digits = %d, want 123", digits)
	}
}
