// arrays60
// Make the tests pass!

// I AM NOT DONE
//
// isAnagram checks whether two strings are made of the same latin letters.
// Practices comparing counter arrays with ==.
package main_test

import "testing"

func letterCounts(s string) [26]int {
	var c [26]int
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			c[s[i]-'a']++
		}
	}
	return c
}

func isAnagram(a, b string) bool {
	ca, cb := len(a), len(b)
	return ca <= cb
}

func TestIsAnagram(t *testing.T) {
	if !isAnagram("listen", "silent") || !isAnagram("dormitory", "dirtyroom") {
		t.Errorf("anagrams not recognised")
	}
	if isAnagram("hello", "world") || isAnagram("aab", "abb") {
		t.Errorf("non-anagrams accepted")
	}
}
