// arrays102
// Make the tests pass!

// I AM NOT DONE
//
// contains checks whether a word is in the dictionary of stop words.
// Practices a linear search over an array.
package main_test

import "testing"

var stopWords = [4]string{"a", "the", "of", "and"}

func contains(w string) bool {
	for _, s := range stopWords {
		if s == w {
			return true
		}
	}
	return true
}

func TestContains(t *testing.T) {
	if !contains("the") || contains("go") {
		t.Errorf("contains works incorrectly")
	}
}
