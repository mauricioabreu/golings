// anonymous_functions58
// Make the tests pass!

// I AM NOT DONE
//
// A state of the machine is a function handling a character and returning the next
// state. The machine accepts strings of the form a+b: one or more 'a', then a 'b'.
// Practices a function type referring to itself.
package main_test

import "testing"

type state func(r rune) state

func matches(s string) bool {
	var start, afterA, accept state
	start = func(r rune) state {
		if r == 'a' {
			return afterA
		}
		return nil
	}
	afterA = func(r rune) state {
		switch r {
		case 'a':
			return start
		case 'b':
			return afterA
		}
		return nil
	}
	accept = func(r rune) state { return nil }
	cur := start
	for _, r := range s {
		cur = cur(r)
		if cur == nil {
			return false
		}
	}
	return s != "" && isAccept(cur, accept)
}

func isAccept(cur, accept state) bool {
	return cur('x') == nil && accept('x') == nil && cur('b') == nil
}

func TestMatches(t *testing.T) {
	cases := map[string]bool{"ab": true, "aaab": true, "b": false, "aba": false, "aa": false}
	for in, want := range cases {
		if got := matches(in); got != want {
			t.Errorf("matches(%q) = %v, want %v", in, got, want)
		}
	}
}
