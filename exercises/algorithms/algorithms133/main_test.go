// algorithms133
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: 1D DP. In how many ways can you climb n stairs taking steps of
// 1 or 2 stairs? For n = 0 the answer is 1.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func climbStairs(n int) int {
	return 0
}

func TestClimbStairs(t *testing.T) {
	cases := map[int]int{0: 1, 1: 1, 2: 2, 3: 3, 5: 8, 45: 1836311903}
	for n, want := range cases {
		if got := climbStairs(n); got != want {
			t.Errorf("climbStairs(%d) = %d, want %d", n, got, want)
		}
	}
}
