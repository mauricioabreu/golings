// algorithms138
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a stack. A string like "3[a2[c]]" expands to "accaccacc".
// Numbers are positive, brackets are balanced, and the inside contains
// Latin letters.
// Expected asymptotics: O(n·k) time, O(n) space.
package main_test

import (
	"strings"
	"testing"
)

func decodeString(s string) string {
	return ""
}

func TestDecodeString(t *testing.T) {
	_ = strings.Repeat
	cases := map[string]string{
		"3[a]2[bc]":     "aaabcbc",
		"3[a2[c]]":      "accaccacc",
		"2[abc]3[cd]ef": "abcabccdcdcdef",
		"abc":           "abc",
		"":              "",
		"10[x]":         "xxxxxxxxxx",
	}
	for in, want := range cases {
		if got := decodeString(in); got != want {
			t.Errorf("decodeString(%q) = %q, want %q", in, got, want)
		}
	}
}
