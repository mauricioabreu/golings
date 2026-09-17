// algorithms65
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two pointers. Check whether s can be obtained from t by
// removing some characters without changing the order of the rest.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func isSubsequence(s, t string) bool {
	return false
}

func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		s, t string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"", "abc", true},
		{"a", "", false},
		{"", "", true},
		{"aaa", "aa", false},
	}
	for _, c := range cases {
		if got := isSubsequence(c.s, c.t); got != c.want {
			t.Errorf("isSubsequence(%q, %q) = %v, want %v", c.s, c.t, got, c.want)
		}
	}
}
