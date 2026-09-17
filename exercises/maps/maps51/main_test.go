// maps51
// Make the tests pass!

// I AM NOT DONE
//
// charCounts counts the characters of a string. For greek text the map ends up
// holding single UTF-8 bytes, and the characters are counted wrongly.
// Indexing a string gives bytes, not runes.
package main_test

import "testing"

func charCounts(s string) map[byte]int {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	return m
}

func TestCharCounts(t *testing.T) {
	m := charCounts("λέξη")
	if len(m) != 4 || m['ξ'] != 1 {
		t.Errorf("charCounts = %v", m)
	}
}
