// functions62
// Make the tests pass!

// I AM NOT DONE
//
// cleanTag must strip everything but letters and digits from both ends of a string.
// Practices strings.TrimFunc with a predicate of your own.
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func notAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func cleanTag(s string) string {
	return strings.TrimFunc(s, notAlnum)
}

func TestCleanTag(t *testing.T) {
	cases := map[string]string{"  #golang! ": "golang", "go1.22": "go1.22", "--": ""}
	for in, want := range cases {
		if got := cleanTag(in); got != want {
			t.Errorf("cleanTag(%q) = %q, want %q", in, got, want)
		}
	}
}
