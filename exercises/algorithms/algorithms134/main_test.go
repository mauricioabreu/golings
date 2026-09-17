// algorithms134
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: interval DP. Popping balloon i earns you nums[i-1]*nums[i]*nums[i+1]
// (values outside the bounds count as 1). Find the maximum score over all
// balloons.
// Expected asymptotics: O(n³) time, O(n²) space.
package main_test

import "testing"

func maxCoins(nums []int) int {
	return 0
}

func TestMaxCoins(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
		{[]int{7}, 7},
		{nil, 0},
		{[]int{9, 76, 64, 21}, 116718},
	}
	for _, c := range cases {
		if got := maxCoins(c.nums); got != c.want {
			t.Errorf("maxCoins(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
