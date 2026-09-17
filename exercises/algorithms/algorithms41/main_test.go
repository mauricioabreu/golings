// algorithms41
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: DP with binary search. Find the length of the longest strictly
// increasing subsequence.
// Expected asymptotics: O(n·log n) time, O(n) space.
package main_test

import (
	"sort"
	"testing"
)

func lengthOfLIS(nums []int) int {
	_ = sort.SearchInts
	return 0
}

func TestLengthOfLIS(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7}, 1},
		{nil, 0},
		{[]int{-5}, 1},
	}
	for _, c := range cases {
		if got := lengthOfLIS(c.nums); got != c.want {
			t.Errorf("lengthOfLIS(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
	big := make([]int, 100000)
	for i := range big {
		big[i] = i % 1000
	}
	if got := lengthOfLIS(big); got != 1000 {
		t.Errorf("big input = %d, want 1000", got)
	}
}
