// algorithms21
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: sliding window. Find the length of the longest substring
// (in characters) in which all characters are distinct.
// Expected asymptotics: O(n) time, O(k) space.
package main_test

import "testing"

func lengthOfLongestSubstring(s string) int {
	return 0
}

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
		"":         0,
		"a":        1,
		"abba":     2,
		"γράφωκγ":  6,
	}
	for in, want := range cases {
		if got := lengthOfLongestSubstring(in); got != want {
			t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", in, got, want)
		}
	}
}
