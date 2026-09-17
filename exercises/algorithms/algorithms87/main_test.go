// algorithms87
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a sliding window with counters. Find the shortest substring of s
// containing all characters of t (with multiplicity). If none — "".
// If several minimal windows exist, return the leftmost one.
// Expected asymptotics: O(n + m) time, O(k) space.
package main_test

import "testing"

func minWindow(s, t string) string {
	return ""
}

func TestMinWindow(t *testing.T) {
	cases := []struct{ s, t, want string }{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"", "a", ""},
		{"abc", "", ""},
		{"aaflslflsldkalskaaa", "aaa", "aaa"},
		{"bba", "ab", "ba"},
	}
	for _, c := range cases {
		if got := minWindow(c.s, c.t); got != c.want {
			t.Errorf("minWindow(%q, %q) = %q, want %q", c.s, c.t, got, c.want)
		}
	}
}
