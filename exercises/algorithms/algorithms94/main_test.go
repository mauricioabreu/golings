// algorithms94
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a stack of indices. Find the length of the longest substring
// of '(' and ')' that is a valid parentheses sequence.
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import "testing"

func longestValidParentheses(s string) int {
	return 0
}

func TestLongestValidParentheses(t *testing.T) {
	cases := map[string]int{"(()": 2, ")()())": 4, "": 0, "(": 0, "()(())": 6, "())(()": 2, "(()())": 6}
	for in, want := range cases {
		if got := longestValidParentheses(in); got != want {
			t.Errorf("longestValidParentheses(%q) = %d, want %d", in, got, want)
		}
	}
}
