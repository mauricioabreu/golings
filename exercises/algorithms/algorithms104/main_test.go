// algorithms104
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: character-by-character comparison. Return the longest common
// prefix of all strings; if there is none — an empty string.
// Expected asymptotics: O(S) time (S is the sum of lengths), O(1) space.
package main_test

import (
	"strings"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	return ""
}

func TestLongestCommonPrefix(t *testing.T) {
	_ = strings.HasPrefix
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"same", "same"}, "same"},
		{[]string{"a"}, "a"},
		{nil, ""},
		{[]string{"", "abc"}, ""},
	}
	for _, c := range cases {
		if got := longestCommonPrefix(c.in); got != c.want {
			t.Errorf("longestCommonPrefix(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
