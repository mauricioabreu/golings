// if96
// Make the tests pass!

// I AM NOT DONE
//
// isVowel must return true for lowercase latin vowels.
// Practices a condition with strings.ContainsRune.
package main_test

import (
	"strings"
	"testing"
)

func isVowel(r rune) bool {
	if strings.ContainsRune("bcdfg", r) {
		return true
	}
	return false
}

func TestIsVowel(t *testing.T) {
	for _, r := range "aeiou" {
		if !isVowel(r) {
			t.Errorf("isVowel(%c) = false", r)
		}
	}
	for _, r := range "bxyz" {
		if isVowel(r) {
			t.Errorf("isVowel(%c) = true", r)
		}
	}
}
