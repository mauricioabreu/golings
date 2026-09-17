// switch103
// Make the tests pass!

// I AM NOT DONE
//
// balanced checks that the brackets ()[]{} are properly nested.
// Practices a switch on a character together with a stack.
package main_test

import "testing"

func balanced(s string) bool {
	var st []rune
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			st = append(st, r)
		case ')', ']', '}':
			st = st[:len(st)-1]
		}
	}
	return true
}

func TestBalanced(t *testing.T) {
	cases := map[string]bool{"": true, "([]{})": true, "a(b)c": true, "(]": false, "((": false, "())": false}
	for in, want := range cases {
		if got := balanced(in); got != want {
			t.Errorf("balanced(%q) = %v, want %v", in, got, want)
		}
	}
}
