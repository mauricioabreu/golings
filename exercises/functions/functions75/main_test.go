// functions75
// Make the tests pass!

// I AM NOT DONE
//
// countIf counts the elements that pred returns true for.
// The isEven predicate has a bug.
// Practices predicate functions.
package main_test

import "testing"

func isEven(n int) bool {
	return n%2 == 1
}

func countIf(nums []int, pred func(int) bool) int {
	c := 0
	for _, n := range nums {
		if pred(n) {
			c++
		}
	}
	return c
}

func TestCountEven(t *testing.T) {
	if got := countIf([]int{1, 2, 3, 4, -6, 0}, isEven); got != 4 {
		t.Errorf("countIf(isEven) = %d, want 4", got)
	}
}
