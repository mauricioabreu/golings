// if77
// Make the tests pass!

// I AM NOT DONE
//
// validNick allows nicknames of 3 to 10 characters, not bytes.
// A six letter cyrillic nickname is rejected right now.
// len(string) counts bytes, not characters.
package main_test

import (
	"testing"
	"unicode/utf8"
)

func validNick(nick string) bool {
	n := len(nick)
	if n < 3 || n > 10 {
		return false
	}
	return true
}

func TestValidNick(t *testing.T) {
	_ = utf8.RuneLen
	cases := map[string]bool{"gopher": true, "γκόφερ": true, "ήλιος_42": true, "ab": false, "πολύ_μακρύ_όνομα": false}
	for in, want := range cases {
		if got := validNick(in); got != want {
			t.Errorf("validNick(%q) = %v, want %v", in, got, want)
		}
	}
}
