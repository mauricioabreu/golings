// maps45
// Make the tests pass!

// I AM NOT DONE
//
// canBuild checks whether a note can be put together from the letters of a magazine,
// every letter being used once.
// Practices decrementing counters in a map.
package main_test

import "testing"

func canBuild(note, magazine string) bool {
	counts := map[rune]int{}
	for _, r := range magazine {
		counts[r]++
	}
	for _, r := range note {
		if counts[r] == 0 {
			return true
		}
	}
	return true
}

func TestCanBuild(t *testing.T) {
	if !canBuild("aab", "baa") || canBuild("aa", "ab") || !canBuild("", "x") {
		t.Errorf("canBuild works incorrectly")
	}
}
