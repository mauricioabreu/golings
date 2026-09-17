// anonymous_functions5
// Make the tests pass!

// I AM NOT DONE
//
// countWords counts the words in several goroutines, each writing into a shared
// captured variable with no synchronization.
// Literals in goroutines changing a shared variable make a data race.
package main_test

import (
	"strings"
	"testing"
)

func countWords(texts []string) int {
	total := 0
	done := make(chan struct{})
	for _, t := range texts {
		go func(t string) {
			total += len(strings.Fields(t))
			done <- struct{}{}
		}(t)
	}
	for range texts {
		<-done
	}
	return total
}

func TestCountWords(t *testing.T) {
	texts := make([]string, 50)
	for i := range texts {
		texts[i] = "one two three"
	}
	if got := countWords(texts); got != 150 {
		t.Errorf("countWords = %d, want 150", got)
	}
}
