// algorithms67
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: 2D dynamic programming. The pattern contains '.' (any character)
// and '*' (zero or more repetitions of the previous character). The pattern
// must cover the entire string.
// Expected asymptotics: O(n·m) time, O(n·m) space.
package main_test

import "testing"

func isMatch(s, p string) bool {
	return false
}

func TestIsMatch(t *testing.T) {
	cases := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"", ".*", true},
		{"", "", true},
		{"abc", "", false},
	}
	for _, c := range cases {
		if got := isMatch(c.s, c.p); got != c.want {
			t.Errorf("isMatch(%q, %q) = %v, want %v", c.s, c.p, got, c.want)
		}
	}
}
