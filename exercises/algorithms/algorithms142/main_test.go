// algorithms142
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: 2D DP. Count in how many ways string t can be obtained as a
// subsequence of string s.
// Expected asymptotics: O(n·m) time, O(m) space.
package main_test

import "testing"

func numDistinct(s, t string) int {
	return 0
}

func TestNumDistinct(t *testing.T) {
	cases := []struct {
		s, t string
		want int
	}{
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
		{"abc", "", 1},
		{"", "a", 0},
		{"aaa", "aa", 3},
	}
	for _, c := range cases {
		if got := numDistinct(c.s, c.t); got != c.want {
			t.Errorf("numDistinct(%q, %q) = %d, want %d", c.s, c.t, got, c.want)
		}
	}
}
