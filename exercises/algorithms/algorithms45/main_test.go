// algorithms45
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: stack. Check that the brackets ()[]{} in the string are properly
// nested and closed. No other characters occur.
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import "testing"

func isValid(s string) bool {
	return false
}

func TestIsValid(t *testing.T) {
	cases := map[string]bool{"()": true, "()[]{}": true, "(]": false, "([)]": false, "{[]}": true, "": true, "(": false, "]": false}
	for in, want := range cases {
		if got := isValid(in); got != want {
			t.Errorf("isValid(%q) = %v, want %v", in, got, want)
		}
	}
}
