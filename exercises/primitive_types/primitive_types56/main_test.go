// primitive_types56
// Make the tests pass!

// I AM NOT DONE
//
// vowelRatio returns the share of vowels among the latin letters of a string, 0 when there are none.
// Practices float64 division and guarding against a division by zero.
package main_test

import (
	"strings"
	"testing"
)

func vowelRatio(s string) float64 {
	lettersCount, vowels := 0, 0
	for _, r := range strings.ToLower(s) {
		if r < 'a' || r > 'z' {
			continue
		}
		lettersCount++
		if strings.ContainsRune("aeiou", r) {
			vowels++
		}
	}
	return float64(vowels / lettersCount)
}

func TestVowelRatio(t *testing.T) {
	cases := map[string]float64{"Gopher": 2.0 / 6, "aaaa": 1, "123": 0, "": 0}
	for in, want := range cases {
		if got := vowelRatio(in); got != want {
			t.Errorf("vowelRatio(%q) = %v, want %v", in, got, want)
		}
	}
}
