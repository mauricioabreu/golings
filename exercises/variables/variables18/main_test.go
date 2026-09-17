// variables18
// Make the tests pass!

// I AM NOT DONE
//
// indexOf must return the index of the first occurrence of a character, or -1.
// Right now it always returns the initial value of i.
// A := in a for header declares a new variable.
package main_test

import "testing"

func indexOf(s string, c byte) int {
	i := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			break
		}
	}
	if i == len(s) {
		return -1
	}
	return i
}

func TestIndexOf(t *testing.T) {
	cases := []struct {
		s    string
		c    byte
		want int
	}{{"golang", 'l', 2}, {"golang", 'g', 0}, {"golang", 'x', -1}, {"", 'a', -1}}
	for _, c := range cases {
		if got := indexOf(c.s, c.c); got != c.want {
			t.Errorf("indexOf(%q, %q) = %d, want %d", c.s, c.c, got, c.want)
		}
	}
}
