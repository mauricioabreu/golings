// algorithms13
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: counting frequencies. Check that the strings consist of the same
// Unicode characters in the same quantities.
// Expected asymptotics: O(n) time, O(k) space (k is the alphabet size).
package main_test

import "testing"

func isAnagram(s, t string) bool {
	return false
}

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "", false},
		{"ρόδο", "δόρο", true},
		{"aab", "abb", false},
	}
	for _, c := range cases {
		if got := isAnagram(c.s, c.t); got != c.want {
			t.Errorf("isAnagram(%q, %q) = %v, want %v", c.s, c.t, got, c.want)
		}
	}
}
