// if82
// Make the tests pass!

// I AM NOT DONE
//
// isDigitCode must check that the code of a character is between '0' and '9'.
// The code does not compile: a < x < b is not valid Go.
// Comparisons do not chain the way they do in mathematics.
package main_test

import "testing"

func isDigitCode(c byte) bool {
	if '0' <= c <= '9' {
		return true
	}
	return false
}

func TestIsDigitCode(t *testing.T) {
	for _, c := range []byte("0459") {
		if !isDigitCode(c) {
			t.Errorf("isDigitCode(%q) = false", c)
		}
	}
	for _, c := range []byte("a/:") {
		if isDigitCode(c) {
			t.Errorf("isDigitCode(%q) = true", c)
		}
	}
}
