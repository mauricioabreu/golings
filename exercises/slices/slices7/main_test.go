// slices7
// Make the tests pass!

// I AM NOT DONE
//
// countLong counts the words longer than three characters.
// Practices a condition in a loop over a slice.
package main_test

import "testing"

func countLong(words []string) int {
	n := 0
	for _, w := range words {
		if len(w) >= 3 {
			n++
		}
	}
	return n
}

func TestCountLong(t *testing.T) {
	if got := countLong([]string{"go", "java", "c", "rust", "zig"}); got != 2 {
		t.Errorf("countLong = %d, want 2", got)
	}
}
