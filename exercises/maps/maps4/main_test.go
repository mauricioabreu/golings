// maps4
// Make the tests pass!

// I AM NOT DONE
//
// wordCount counts the words of a text.
// Practices m[k]++ in a loop.
package main_test

import (
	"strings"
	"testing"
)

func wordCount(text string) map[string]int {
	counts := map[string]int{}
	for _, w := range strings.Fields(text) {
		counts[text]++
	}
	return counts
}

func TestWordCount(t *testing.T) {
	c := wordCount("go is go")
	if c["go"] != 2 || c["is"] != 1 || len(c) != 2 {
		t.Errorf("wordCount = %v", c)
	}
}
