// algorithms135
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: 2D DP over prefixes. Find the length of the longest common
// subsequence of two strings.
// Expected asymptotics: O(n·m) time, O(min(n, m)) space.
package main_test

import "testing"

func longestCommonSubsequence(a, b string) int {
	return 0
}

func TestLongestCommonSubsequence(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"", "abc", 0},
		{"", "", 0},
		{"bsbininm", "jmjkbkjkv", 1},
	}
	for _, c := range cases {
		if got := longestCommonSubsequence(c.a, c.b); got != c.want {
			t.Errorf("lcs(%q, %q) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
