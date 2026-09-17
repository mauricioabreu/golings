// algorithms15
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: sorting and the longest increasing subsequence.
// An envelope [w, h] fits into another if it is strictly smaller on
// both sides. Return the maximum number of nested envelopes.
// Expected asymptotics: O(n·log n) time, O(n) space.
package main_test

import (
	"sort"
	"testing"
)

func maxEnvelopes(envelopes [][2]int) int {
	_ = sort.SearchInts
	return 0
}

func TestMaxEnvelopes(t *testing.T) {
	cases := []struct {
		env  [][2]int
		want int
	}{
		{[][2]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
		{[][2]int{{1, 1}, {1, 1}, {1, 1}}, 1},
		{nil, 0},
		{[][2]int{{4, 5}}, 1},
		{[][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 4},
	}
	for _, c := range cases {
		if got := maxEnvelopes(c.env); got != c.want {
			t.Errorf("maxEnvelopes(%v) = %d, want %d", c.env, got, c.want)
		}
	}
}
