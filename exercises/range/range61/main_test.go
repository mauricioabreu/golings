// range61
// Make the tests pass!

// I AM NOT DONE
//
// longestWord returns the longest word, the first one on a tie.
// Practices a range over the result of strings.Fields.
package main_test

import (
	"strings"
	"testing"
)

func longestWord(text string) string {
	best := ""
	for _, w := range strings.Fields(text) {
		if len(w) >= len(best) {
			best = w
		}
	}
	return best
}

func TestLongestWord(t *testing.T) {
	if got := longestWord("go is fun and cool"); got != "cool" {
		t.Errorf("longestWord = %q, want cool", got)
	}
	if got := longestWord("abc xyz"); got != "abc" {
		t.Errorf("longestWord = %q, want abc", got)
	}
}
