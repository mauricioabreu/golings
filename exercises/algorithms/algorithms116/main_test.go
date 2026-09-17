// algorithms116
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: greedy over the "reach boundary". nums[i] is the maximum jump
// length from position i. Return the minimum number of jumps to reach the
// end (reaching it is guaranteed).
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func jump(nums []int) int {
	return 0
}

func TestJump(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{0}, 0},
		{nil, 0},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{10, 1, 1}, 1},
	}
	for _, c := range cases {
		if got := jump(c.nums); got != c.want {
			t.Errorf("jump(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
