// anonymous_functions87
// Make the tests pass!

// I AM NOT DONE
//
// best picks an element with a comparator; longest passes a "longer" literal.
// Practices a literal with two parameters.
package main_test

import "testing"

func best(s []string, better func(a, b string) bool) string {
	b := s[0]
	for _, v := range s[1:] {
		if better(v, b) {
			b = v
		}
	}
	return b
}

func longest(s []string) string {
	return best(s, func(a, b string) bool { return len(a) < len(b) })
}

func TestLongest(t *testing.T) {
	if got := longest([]string{"go", "gopher", "gc"}); got != "gopher" {
		t.Errorf("longest = %q", got)
	}
}
