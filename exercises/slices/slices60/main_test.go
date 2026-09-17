// slices60
// Make the tests pass!

// I AM NOT DONE
//
// reverseWords puts the words of a string in reverse order.
// Practices reversing a slice of strings in place.
package main_test

import (
	"strings"
	"testing"
)

func reverseWords(s string) string {
	w := strings.Fields(s)
	for i := 0; i < len(w); i++ {
		w[i], w[len(w)-1] = w[len(w)-1], w[i]
	}
	return strings.Join(w, " ")
}

func TestReverseWords(t *testing.T) {
	cases := map[string]string{"one two three": "three two one", "go": "go", "": "", "a b c d": "d c b a"}
	for in, want := range cases {
		if got := reverseWords(in); got != want {
			t.Errorf("reverseWords(%q) = %q, want %q", in, got, want)
		}
	}
}
