// switch23
// Make the tests pass!

// I AM NOT DONE
//
// greetingLang guesses the language from the first letter: 'γ' is "el", 'h' is "en".
// The code does not compile: the rune 'γ' does not fit in a byte.
// s[0] is a byte, and the case constants have to fit the type of the tag.
package main_test

import (
	"testing"
	"unicode/utf8"
)

func greetingLang(s string) string {
	if s == "" {
		return "unknown"
	}
	switch s[0] {
	case 'γ':
		return "el"
	case 'h':
		return "en"
	}
	return "unknown"
}

func TestGreetingLang(t *testing.T) {
	_ = utf8.RuneLen
	cases := map[string]string{"γειά": "el", "hello": "en", "hola": "en", "bonjour": "unknown", "": "unknown"}
	for in, want := range cases {
		if got := greetingLang(in); got != want {
			t.Errorf("greetingLang(%q) = %s, want %s", in, got, want)
		}
	}
}
