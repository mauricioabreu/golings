// primitive_types39
// Make the tests pass!

// I AM NOT DONE
//
// isGreek checks that a string holds lowercase greek letters only.
// Words with an accented letter are rejected right now.
// The accented vowels sit outside the contiguous 'α'..'ω' range in Unicode.
package main_test

import "testing"

func isGreek(s string) bool {
	for _, r := range s {
		if r < 'α' || r > 'ω' {
			return false
		}
	}
	return s != ""
}

func TestIsGreek(t *testing.T) {
	cases := map[string]bool{"άλφα": true, "λόγος": true, "φως": true, "hello": false, "Λόγος": false, "": false}
	for in, want := range cases {
		if got := isGreek(in); got != want {
			t.Errorf("isGreek(%q) = %v, want %v", in, got, want)
		}
	}
}
