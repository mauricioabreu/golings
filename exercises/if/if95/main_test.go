// if95
// Make the tests pass!

// I AM NOT DONE
//
// isQuestion must return true for strings ending in '?'.
// It panics on an empty string.
// len(s)-1 is -1 for an empty string.
package main_test

import "testing"

func isQuestion(s string) bool {
	if s[len(s)-1] == '?' {
		return true
	}
	return false
}

func TestIsQuestion(t *testing.T) {
	cases := map[string]bool{"why?": true, "ok": false, "": false, "?": true}
	for in, want := range cases {
		if got := isQuestion(in); got != want {
			t.Errorf("isQuestion(%q) = %v, want %v", in, got, want)
		}
	}
}
