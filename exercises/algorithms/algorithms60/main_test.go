// algorithms60
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: 1D DP. Each step costs cost[i]; you may start at step zero or
// one and step up by 1 or 2. Return the minimum cost of reaching the top
// (beyond the last step).
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func minCostClimbingStairs(cost []int) int {
	return 0
}

func TestMinCostClimbingStairs(t *testing.T) {
	cases := []struct {
		cost []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{[]int{5, 5}, 5},
		{nil, 0},
		{[]int{7}, 0},
	}
	for _, c := range cases {
		if got := minCostClimbingStairs(c.cost); got != c.want {
			t.Errorf("minCostClimbingStairs(%v) = %d, want %d", c.cost, got, c.want)
		}
	}
}
