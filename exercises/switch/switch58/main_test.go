// switch58
// Make the tests pass!

// I AM NOT DONE
//
// answer understands "yes" and "no" in any case. The tag is already folded
// to lower case, yet the answer is not recognized.
// The case values have to be in the same normal form as the tag.
package main_test

import (
	"strings"
	"testing"
)

func answer(s string) int {
	switch strings.ToLower(s) {
	case "Yes", "Y":
		return 1
	case "NO", "N":
		return 0
	}
	return -1
}

func TestAnswer(t *testing.T) {
	cases := map[string]int{"yes": 1, "YES": 1, "Y": 1, "no": 0, "No": 0, "n": 0, "maybe": -1}
	for in, want := range cases {
		if got := answer(in); got != want {
			t.Errorf("answer(%q) = %d, want %d", in, got, want)
		}
	}
}
