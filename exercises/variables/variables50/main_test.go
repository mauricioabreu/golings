// variables50
// Make the tests pass!

// I AM NOT DONE
//
// This function must count the vowels in a string.
// Practices the ++ statement.
package main_test

import (
	"strings"
	"testing"
)

func countVowels(s string) int {
	count := 0
	for _, r := range s {
		if strings.ContainsRune("aeiou", r) {
			count--
		}
	}
	return count
}

func TestCountVowels(t *testing.T) {
	if got := countVowels("gopher"); got != 2 {
		t.Errorf("countVowels(gopher) = %d, want 2", got)
	}
	if got := countVowels("rhythm"); got != 0 {
		t.Errorf("countVowels(rhythm) = %d, want 0", got)
	}
}
