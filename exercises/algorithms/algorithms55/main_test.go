// algorithms55
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two pointers. Wall heights are given in a slice. Find the maximum
// area of water between two walls (width × the smaller height).
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func maxArea(h []int) int {
	return 0
}

func TestMaxArea(t *testing.T) {
	cases := []struct {
		h    []int
		want int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{5}, 0},
		{nil, 0},
		{[]int{1000000, 1, 1000000}, 2000000},
	}
	for _, c := range cases {
		if got := maxArea(c.h); got != c.want {
			t.Errorf("maxArea(%v) = %d, want %d", c.h, got, c.want)
		}
	}
}
