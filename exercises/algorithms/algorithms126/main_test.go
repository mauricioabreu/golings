// algorithms126
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: greedy with accumulation. On a circular route, station i gives
// gas[i] fuel, and driving to the next station costs cost[i]. Return the
// starting index from which the full circle can be completed, or -1.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func canCompleteCircuit(gas, cost []int) int {
	return 0
}

func TestCanCompleteCircuit(t *testing.T) {
	cases := []struct {
		gas, cost []int
		want      int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		{[]int{5}, []int{4}, 0},
		{[]int{1}, []int{2}, -1},
		{nil, nil, 0},
	}
	for _, c := range cases {
		if got := canCompleteCircuit(c.gas, c.cost); got != c.want {
			t.Errorf("canCompleteCircuit(%v, %v) = %d, want %d", c.gas, c.cost, got, c.want)
		}
	}
}
