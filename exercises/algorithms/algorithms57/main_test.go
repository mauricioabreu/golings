// algorithms57
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: greedy after sorting. A child with appetite g[i] is content with
// a cookie of size s[j] >= g[i]. Each child gets at most one cookie.
// Return the maximum number of content children.
// Expected asymptotics: O(n·log n) time, O(1) extra space.
package main_test

import (
	"sort"
	"testing"
)

func findContentChildren(g, s []int) int {
	_ = sort.Ints
	return 0
}

func TestFindContentChildren(t *testing.T) {
	cases := []struct {
		g, s []int
		want int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{nil, []int{1}, 0},
		{[]int{5}, nil, 0},
		{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}, 2},
	}
	for _, c := range cases {
		if got := findContentChildren(c.g, c.s); got != c.want {
			t.Errorf("findContentChildren(%v, %v) = %d, want %d", c.g, c.s, got, c.want)
		}
	}
}
