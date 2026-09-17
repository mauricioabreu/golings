// algorithms23
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: binary search on the answer. There are piles of bananas and
// h hours. In one hour you can eat up to k bananas from a single pile.
// Find the minimum speed k.
// Expected asymptotics: O(n·log m) time, O(1) space.
package main_test

import "testing"

func minEatingSpeed(piles []int, h int) int {
	return 0
}

func TestMinEatingSpeed(t *testing.T) {
	cases := []struct {
		piles   []int
		h, want int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
		{[]int{1}, 1, 1},
		{[]int{1000000000}, 2, 500000000},
	}
	for _, c := range cases {
		if got := minEatingSpeed(c.piles, c.h); got != c.want {
			t.Errorf("minEatingSpeed(%v, %d) = %d, want %d", c.piles, c.h, got, c.want)
		}
	}
}
