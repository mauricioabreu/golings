// algorithms106
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two pointers. Check that a string is a palindrome, considering
// only ASCII letters and digits and ignoring case.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func isAlnum(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}

func isPalindrome(s string) bool {
	return false
}

func TestIsPalindrome(t *testing.T) {
	cases := map[string]bool{
		"A man, a plan, a canal: Panama": true,
		"race a car":                     false,
		" ":                              true,
		"":                               true,
		"0P":                             false,
		"ab_a":                           true,
	}
	for in, want := range cases {
		if got := isPalindrome(in); got != want {
			t.Errorf("isPalindrome(%q) = %v, want %v", in, got, want)
		}
	}
}
