// range94
// Make the tests pass!

// I AM NOT DONE
//
// withCommas puts a comma after every character but the last.
// When the last character is multibyte, a comma turns up after it too.
// i == len(s)-1 does not identify the last rune.
package main_test

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func withCommas(s string) string {
	var b strings.Builder
	for i, r := range s {
		b.WriteRune(r)
		if i != len(s)-1 {
			b.WriteByte(',')
		}
	}
	return b.String()
}

func TestWithCommas(t *testing.T) {
	_ = utf8.RuneLen
	cases := map[string]string{"abc": "a,b,c", "abω": "a,b,ω", "λ": "λ"}
	for in, want := range cases {
		if got := withCommas(in); got != want {
			t.Errorf("withCommas(%q) = %q, want %q", in, got, want)
		}
	}
}
