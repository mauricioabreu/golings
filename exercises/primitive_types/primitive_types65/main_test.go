// primitive_types65
// Make the tests pass!

// I AM NOT DONE
//
// mask must replace every character of a string but the first with '•'.
// Right now the result holds junk instead of the dots.
// '•' is a rune outside the byte range and has to be written with WriteRune.
package main_test

import (
	"strings"
	"testing"
)

func mask(s string) string {
	var b strings.Builder
	for i, r := range []rune(s) {
		if i == 0 {
			b.WriteRune(r)
			continue
		}
		b.WriteByte(byte('•'))
	}
	return b.String()
}

func TestMask(t *testing.T) {
	cases := map[string]string{"secret": "s•••••", "κωδικός": "κ••••••", "x": "x"}
	for in, want := range cases {
		if got := mask(in); got != want {
			t.Errorf("mask(%q) = %q, want %q", in, got, want)
		}
	}
}
