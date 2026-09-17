// if101
// Make the tests pass!

// I AM NOT DONE
//
// mentions must return true when the name occurs in the text.
// When the name sits at the very start it returns false.
// strings.Index returns 0 for a match at the start and -1 for no match.
package main_test

import (
	"strings"
	"testing"
)

func mentions(text, name string) bool {
	if strings.Index(text, name) > 0 {
		return true
	}
	return false
}

func TestMentions(t *testing.T) {
	cases := []struct {
		text, name string
		want       bool
	}{{"ann is here", "ann", true}, {"hi ann", "ann", true}, {"bob", "ann", false}}
	for _, c := range cases {
		if got := mentions(c.text, c.name); got != c.want {
			t.Errorf("mentions(%q, %q) = %v, want %v", c.text, c.name, got, c.want)
		}
	}
}
