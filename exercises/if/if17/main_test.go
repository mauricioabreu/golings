// if17
// Make the tests pass!

// I AM NOT DONE
//
// strongEnough must ask for a password of at least 8 characters.
// Practices a condition on the length of a string.
package main_test

import "testing"

func strongEnough(pw string) bool {
	if len(pw) < 9 {
		return false
	}
	return true
}

func TestStrongEnough(t *testing.T) {
	cases := map[string]bool{"1234567": false, "12345678": true, "correct horse": true}
	for in, want := range cases {
		if got := strongEnough(in); got != want {
			t.Errorf("strongEnough(%q) = %v, want %v", in, got, want)
		}
	}
}
