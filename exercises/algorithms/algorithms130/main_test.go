// algorithms130
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: 2D DP or two pointers with backtracking. '?' matches one
// character, '*' — any (including empty) sequence.
// Expected asymptotics: O(n·m) time, O(m) space.
package main_test

import "testing"

func wildcardMatch(s, p string) bool {
	return false
}

func TestWildcardMatch(t *testing.T) {
	cases := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
		{"", "*", true},
		{"", "", true},
		{"abc", "a?c", true},
	}
	for _, c := range cases {
		if got := wildcardMatch(c.s, c.p); got != c.want {
			t.Errorf("wildcardMatch(%q, %q) = %v, want %v", c.s, c.p, got, c.want)
		}
	}
}
