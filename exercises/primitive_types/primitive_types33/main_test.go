// primitive_types33
// Make the tests pass!

// I AM NOT DONE
//
// truncate must keep at most n characters of a string and add "…" when it was
// longer. Right now cyrillic text turns into broken characters.
// Slicing a string cuts bytes and can split a UTF-8 character.
package main_test

import (
	"testing"
	"unicode/utf8"
)

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}

func TestTruncate(t *testing.T) {
	cases := []struct {
		in   string
		n    int
		want string
	}{{"γειά σου", 4, "γειά…"}, {"go", 4, "go"}, {"hello!", 5, "hello…"}, {"φως", 3, "φως"}}
	for _, c := range cases {
		got := truncate(c.in, c.n)
		if got != c.want || !utf8.ValidString(got) {
			t.Errorf("truncate(%q, %d) = %q, want %q", c.in, c.n, got, c.want)
		}
	}
}
