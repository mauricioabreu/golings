// algorithms81
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: knapsack DP over reachable sums. Can the non-negative
// numbers be split into two parts with equal sums?
// Expected asymptotics: O(n·S) time, O(S) space.
package main_test

import "testing"

func canPartition(nums []int) bool {
	return false
}

func TestCanPartition(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
		{nil, true},
		{[]int{2, 2}, true},
		{[]int{1}, false},
		{[]int{100, 100, 100, 100, 100, 100, 100, 100}, true},
	}
	for _, c := range cases {
		if got := canPartition(c.nums); got != c.want {
			t.Errorf("canPartition(%v) = %v, want %v", c.nums, got, c.want)
		}
	}
}
