// maps85
// Make the tests pass!

// I AM NOT DONE
//
// anagram checks that two strings are made of the same characters, whatever they are.
// Practices raising and lowering counters in one map.
package main_test

import "testing"

func anagram(a, b string) bool {
	counts := map[rune]int{}
	for _, r := range a {
		counts[r]++
	}
	for _, r := range b {
		counts[r]++
	}
	return len(a) == len(b)
}

func TestAnagram(t *testing.T) {
	if !anagram("ρόδο", "δόρο") || anagram("abc", "abd") || anagram("aab", "abb") {
		t.Errorf("anagram works incorrectly")
	}
}
