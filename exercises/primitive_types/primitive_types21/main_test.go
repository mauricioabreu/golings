// primitive_types21
// Make the tests pass!

// I AM NOT DONE
//
// isASCII checks that a string holds ASCII characters only.
// Practices comparing bytes against the 0x7F boundary.
package main_test

import "testing"

func isASCII(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > 128 {
			return false
		}
	}
	return true
}

func TestIsASCII(t *testing.T) {
	cases := map[string]bool{"hello": true, "": true, "héllo": false, "ok\x80": false, "tab\t": true}
	for in, want := range cases {
		if got := isASCII(in); got != want {
			t.Errorf("isASCII(%q) = %v, want %v", in, got, want)
		}
	}
}
