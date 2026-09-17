// switch37
// Make the tests pass!

// I AM NOT DONE
//
// isVowel must return true for all five vowels a, e, i, o and u.
// Practices several values in one case.
package main_test

import "testing"

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o':
		return true
	}
	return false
}

func TestIsVowel(t *testing.T) {
	for _, r := range "aeiou" {
		if !isVowel(r) {
			t.Errorf("isVowel(%q) = false", r)
		}
	}
	if isVowel('z') {
		t.Errorf("isVowel('z') = true")
	}
}
