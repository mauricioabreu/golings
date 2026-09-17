// functions11
// Make the tests pass!

// I AM NOT DONE
//
// longest must return the longest word and its length.
// The code does not compile, even though len is used "as usual".
// A predeclared identifier can be shadowed by a local variable.
package main_test

import "testing"

func longest(words []string) (string, int) {
	len := len(words)
	if len == 0 {
		return "", 0
	}
	best := words[0]
	for _, w := range words[1:] {
		if len(w) > len(best) {
			best = w
		}
	}
	return best, len(best)
}

func TestLongest(t *testing.T) {
	if w, n := longest([]string{"go", "gopher", "gc"}); w != "gopher" || n != 6 {
		t.Errorf("longest = %q, %d; want gopher, 6", w, n)
	}
	if w, n := longest(nil); w != "" || n != 0 {
		t.Errorf("longest(nil) = %q, %d", w, n)
	}
}
