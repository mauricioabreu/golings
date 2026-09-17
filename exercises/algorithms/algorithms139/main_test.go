// algorithms139
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: row histograms and a monotonic stack. In a binary matrix, find
// the area of the largest rectangle consisting only of '1'.
// Expected asymptotics: O(r·c) time, O(c) space.
package main_test

import "testing"

func maximalRectangle(m []string) int {
	return 0
}

func TestMaximalRectangle(t *testing.T) {
	cases := []struct {
		m    []string
		want int
	}{
		{[]string{"10100", "10111", "11111", "10010"}, 6},
		{[]string{"0"}, 0},
		{[]string{"1"}, 1},
		{nil, 0},
		{[]string{"111", "111"}, 6},
		{[]string{"01", "10"}, 1},
	}
	for _, c := range cases {
		if got := maximalRectangle(c.m); got != c.want {
			t.Errorf("maximalRectangle(%v) = %d, want %d", c.m, got, c.want)
		}
	}
}
