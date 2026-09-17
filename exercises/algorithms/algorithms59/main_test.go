// algorithms59
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: topological sort over adjacent word pairs. Given a list of words
// sorted in an unknown alphabet, restore the order of the letters.
// If ambiguous, take the letter with the smaller code; on a contradiction — "".
// Expected asymptotics: O(C) time, O(1) space.
package main_test

import (
	"sort"
	"testing"
)

func alienOrder(words []string) string {
	return ""
}

func TestAlienOrder(t *testing.T) {
	_ = sort.Ints
	cases := []struct {
		words []string
		want  string
	}{
		{[]string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
		{[]string{"z", "x"}, "zx"},
		{[]string{"z", "x", "z"}, ""},
		{[]string{"abc", "ab"}, ""},
		{[]string{"ab", "adc"}, "abcd"},
	}
	for _, c := range cases {
		if got := alienOrder(c.words); got != c.want {
			t.Errorf("alienOrder(%v) = %q, want %q", c.words, got, c.want)
		}
	}
}
