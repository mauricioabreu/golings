// algorithms95
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: expansion from the center. Find the longest palindromic substring
// (in bytes); if several have the same length — the leftmost one.
// Expected asymptotics: O(n²) time, O(1) space.
package main_test

import "testing"

func longestPalindrome(s string) string {
	return ""
}

func TestLongestPalindrome(t *testing.T) {
	cases := map[string]string{
		"babad": "bab",
		"cbbd":  "bb",
		"a":     "a",
		"":      "",
		"ac":    "a",
		"aaaa":  "aaaa",
	}
	for in, want := range cases {
		if got := longestPalindrome(in); got != want {
			t.Errorf("longestPalindrome(%q) = %q, want %q", in, got, want)
		}
	}
}
