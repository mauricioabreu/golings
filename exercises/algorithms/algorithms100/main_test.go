// algorithms100
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: one-dimensional DP over string positions. Can the string be split
// into a sequence of dictionary words (words may be reused)?
// Expected asymptotics: O(n²·L) time, O(n) space.
package main_test

import "testing"

func wordBreak(s string, dict []string) bool {
	return false
}

func TestWordBreak(t *testing.T) {
	cases := []struct {
		s    string
		dict []string
		want bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"", []string{"a"}, true},
		{"a", nil, false},
		{"aaaaaaa", []string{"aaa", "aaaa"}, true},
	}
	for _, c := range cases {
		if got := wordBreak(c.s, c.dict); got != c.want {
			t.Errorf("wordBreak(%q, %v) = %v, want %v", c.s, c.dict, got, c.want)
		}
	}
}
