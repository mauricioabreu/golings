// maps38
// Make the tests pass!

// I AM NOT DONE
//
// isBanned checks whether a word is in the set of banned ones.
// Practices map[string]bool as a set.
package main_test

import "testing"

var banned = map[string]bool{"spam": true, "scam": true}

func isBanned(w string) bool {
	return !banned[w]
}

func TestIsBanned(t *testing.T) {
	if !isBanned("spam") || isBanned("hello") {
		t.Errorf("isBanned works incorrectly")
	}
}
