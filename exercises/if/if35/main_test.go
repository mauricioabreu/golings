// if35
// Make the tests pass!

// I AM NOT DONE
//
// isComment must return true for lines starting with '#'.
// It panics on an empty string.
// Practices the short circuit of the logical operators.
package main_test

import "testing"

func isComment(line string) bool {
	if line[0] == '#' && len(line) > 0 {
		return true
	}
	return false
}

func TestIsComment(t *testing.T) {
	cases := map[string]bool{"# note": true, "code": false, "": false, "#": true}
	for in, want := range cases {
		if got := isComment(in); got != want {
			t.Errorf("isComment(%q) = %v, want %v", in, got, want)
		}
	}
}
