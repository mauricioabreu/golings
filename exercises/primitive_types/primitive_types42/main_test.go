// primitive_types42
// Make the tests pass!

// I AM NOT DONE
//
// charAt must return the n-th character of a string, counting from zero, or an empty string.
// A range over a string gives byte offsets rather than character numbers.
// The index in for i, r := range s is a byte position.
package main_test

import "testing"

func charAt(s string, n int) string {
	for i, r := range s {
		if i == n {
			return string(r)
		}
	}
	return ""
}

func TestCharAt(t *testing.T) {
	cases := []struct {
		s    string
		n    int
		want string
	}{{"hello", 1, "e"}, {"φως", 1, "ω"}, {"φως", 2, "ς"}, {"日本語", 2, "語"}, {"go", 5, ""}}
	for _, c := range cases {
		if got := charAt(c.s, c.n); got != c.want {
			t.Errorf("charAt(%q, %d) = %q, want %q", c.s, c.n, got, c.want)
		}
	}
}
