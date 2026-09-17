// range54
// Make the tests pass!

// I AM NOT DONE
//
// countByte counts the occurrences of a given byte in a string.
// The code does not compile: a range over a string yields runes while the wanted value is a byte.
// The value types of a range over a string and over a []byte differ.
package main_test

import "testing"

func countByte(s string, c byte) int {
	n := 0
	for _, r := range s {
		if r == c {
			n++
		}
	}
	return n
}

func TestCountByte(t *testing.T) {
	if got := countByte("a,b,,c", ','); got != 3 {
		t.Errorf("countByte = %d, want 3", got)
	}
}
