// if6
// Make the tests pass!

// I AM NOT DONE
//
// validHex checks a string such as "#1a2B3c": a hash and exactly 6 hexadecimal digits.
// Practices a length check together with per character checks.
package main_test

import "testing"

func validHex(s string) bool {
	if s[0] != '#' {
		return false
	}
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			continue
		}
		return false
	}
	return true
}

func TestValidHex(t *testing.T) {
	cases := map[string]bool{"#1a2B3c": true, "#000000": true, "123456": false, "#12345": false, "#12345g": false, "": false}
	for in, want := range cases {
		if got := validHex(in); got != want {
			t.Errorf("validHex(%q) = %v, want %v", in, got, want)
		}
	}
}
