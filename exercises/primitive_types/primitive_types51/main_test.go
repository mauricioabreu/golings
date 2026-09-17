// primitive_types51
// Make the tests pass!

// I AM NOT DONE
//
// atbash replaces a latin letter with its mirror: a<->z, b<->y; the case is kept
// and every other character is left alone.
// Practices rune arithmetic over two ranges.
package main_test

import (
	"strings"
	"testing"
)

func mirror(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return 'a' + (r - 'z')
	}
	return r
}

func atbash(s string) string {
	return strings.Map(mirror, s)
}

func TestAtbash(t *testing.T) {
	cases := map[string]string{"abc": "zyx", "Hello!": "Svool!", "Zz 9": "Aa 9"}
	for in, want := range cases {
		if got := atbash(in); got != want {
			t.Errorf("atbash(%q) = %q, want %q", in, got, want)
		}
	}
}
