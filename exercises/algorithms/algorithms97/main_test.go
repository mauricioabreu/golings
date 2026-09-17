// algorithms97
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: one-dimensional DP. Find the minimum number of coins that sum
// to amount (each denomination is unlimited), or -1.
// Expected asymptotics: O(n·amount) time, O(amount) space.
package main_test

import "testing"

func coinChange(coins []int, amount int) int {
	return 0
}

func TestCoinChange(t *testing.T) {
	cases := []struct {
		coins        []int
		amount, want int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{nil, 0, 0},
		{nil, 7, -1},
		{[]int{186, 419, 83, 408}, 6249, 20},
	}
	for _, c := range cases {
		if got := coinChange(c.coins, c.amount); got != c.want {
			t.Errorf("coinChange(%v, %d) = %d, want %d", c.coins, c.amount, got, c.want)
		}
	}
}
