// range91
// Make the tests pass!

// I AM NOT DONE
//
// markVowels returns a string with the vowels replaced by '*'.
// The index of the range over the string is used on a []rune and points at the wrong place.
// The index of a range over a string is a byte offset, not a rune number.
package main_test

import (
	"strings"
	"testing"
)

func markVowels(s string) string {
	rs := []rune(s)
	for i, r := range s {
		if strings.ContainsRune("αεηιουaeiou", r) {
			rs[i] = '*'
		}
	}
	return string(rs)
}

func TestMarkVowels(t *testing.T) {
	cases := map[string]string{"hello": "h*ll*", "φως": "φ*ς", "μια γη": "μ** γ*"}
	for in, want := range cases {
		if got := markVowels(in); got != want {
			t.Errorf("markVowels(%q) = %q, want %q", in, got, want)
		}
	}
}
