// algorithms96
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a single pass tracking the minimum. Prices are given per day;
// buy once and sell later. Return the maximum profit (0 if there is none).
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func maxProfit(prices []int) int {
	return 0
}

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{nil, 0},
		{[]int{5}, 0},
		{[]int{2, 4, 1, 1000000000}, 999999999},
	}
	for _, c := range cases {
		in := append([]int(nil), c.prices...)
		if got := maxProfit(in); got != c.want {
			t.Errorf("maxProfit(%v) = %d, want %d", c.prices, got, c.want)
		}
	}
}
