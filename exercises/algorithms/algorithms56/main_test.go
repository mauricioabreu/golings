// algorithms56
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: DP over states (position, jump length). Stones are given as
// increasing coordinates. The first jump has length 1; after a jump of
// length k the next one may be k-1, k, or k+1. Can you reach the last stone?
// Expected asymptotics: O(n²) time, O(n²) space.
package main_test

import "testing"

func canCross(stones []int) bool {
	return false
}

func TestCanCross(t *testing.T) {
	cases := []struct {
		stones []int
		want   bool
	}{
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
		{[]int{0, 1}, true},
		{[]int{0, 2}, false},
		{[]int{0}, true},
		{nil, false},
	}
	for _, c := range cases {
		if got := canCross(c.stones); got != c.want {
			t.Errorf("canCross(%v) = %v, want %v", c.stones, got, c.want)
		}
	}
}
