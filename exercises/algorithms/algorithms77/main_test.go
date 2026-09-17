// algorithms77
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: prefix sums and a hash table. Count the number of contiguous
// subarrays whose sum equals k (the numbers may be negative).
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import "testing"

func subarraySum(nums []int, k int) int {
	return 0
}

func TestSubarraySum(t *testing.T) {
	cases := []struct {
		nums    []int
		k, want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, -1, 0}, 0, 3},
		{nil, 0, 0},
		{[]int{5}, 5, 1},
		{[]int{0, 0, 0}, 0, 6},
	}
	for _, c := range cases {
		if got := subarraySum(c.nums, c.k); got != c.want {
			t.Errorf("subarraySum(%v, %d) = %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
