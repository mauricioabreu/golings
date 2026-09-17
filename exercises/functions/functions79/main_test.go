// functions79
// Make the tests pass!

// I AM NOT DONE
//
// isPalindrome must check recursively whether a string reads the same both ways.
// Practices recursion over substrings.
package main_test

import "testing"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	return s[0] == s[len(s)-1] && isPalindrome(s[1:])
}

func TestIsPalindrome(t *testing.T) {
	cases := map[string]bool{"": true, "a": true, "abba": true, "racecar": true, "abca": false, "ab": false}
	for in, want := range cases {
		if got := isPalindrome(in); got != want {
			t.Errorf("isPalindrome(%q) = %v, want %v", in, got, want)
		}
	}
}
