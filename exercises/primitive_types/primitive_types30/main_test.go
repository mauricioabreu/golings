// primitive_types30
// Make the tests pass!

// I AM NOT DONE
//
// titleCase must uppercase the first letter of every word and lowercase the rest.
// Practices working with []rune and the unicode functions.
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		r := []rune(w)
		r[0] = unicode.ToUpper(r[0])
	}
	return strings.Join(words, " ")
}

func TestTitleCase(t *testing.T) {
	cases := map[string]string{"hello WORLD": "Hello World", "γειά κόσμε": "Γειά Κόσμε", "": ""}
	for in, want := range cases {
		if got := titleCase(in); got != want {
			t.Errorf("titleCase(%q) = %q, want %q", in, got, want)
		}
	}
}
