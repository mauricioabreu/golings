// algorithms17
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: binary search on the answer with a greedy check. Split a slice
// of non-negative numbers into k non-empty contiguous parts so that the
// largest part sum is minimal. Return that sum.
// Expected asymptotics: O(n·log S) time, O(1) space.
package main_test

import "testing"

func splitArray(nums []int, k int) int {
	return 0
}

func TestSplitArray(t *testing.T) {
	cases := []struct {
		nums    []int
		k, want int
	}{
		{[]int{7, 2, 5, 10, 8}, 2, 18},
		{[]int{1, 2, 3, 4, 5}, 2, 9},
		{[]int{1, 4, 4}, 3, 4},
		{[]int{5}, 1, 5},
		{[]int{0, 0, 0}, 2, 0},
		{[]int{1000000000, 1000000000}, 1, 2000000000},
	}
	for _, c := range cases {
		if got := splitArray(c.nums, c.k); got != c.want {
			t.Errorf("splitArray(%v, %d) = %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
