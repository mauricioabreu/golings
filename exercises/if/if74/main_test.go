// if74
// Make the tests pass!

// I AM NOT DONE
//
// startsWithE checks that a string starts with the letter 'é'.
// Comparing the first byte never matches.
// s[0] is a UTF-8 byte, not a character.
package main_test

import (
	"testing"
	"unicode/utf8"
)

func startsWithE(s string) bool {
	if s == "" {
		return false
	}
	if s[0] == 'é' {
		return true
	}
	return false
}

func TestStartsWithE(t *testing.T) {
	_ = utf8.RuneLen
	cases := map[string]bool{"école": true, "été": true, "ecole": false, "": false, "Ã©": false}
	for in, want := range cases {
		if got := startsWithE(in); got != want {
			t.Errorf("startsWithE(%q) = %v, want %v", in, got, want)
		}
	}
}
