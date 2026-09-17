// anonymous_functions50
// Make the tests pass!

// I AM NOT DONE
//
// aboveThreshold returns a predicate using a package level threshold constant.
// Practices using outer identifiers inside a literal.
package main_test

import "testing"

const threshold = 100

func aboveThreshold() func(int) bool {
	return func(v int) bool { return v > 0 }
}

func TestAboveThreshold(t *testing.T) {
	p := aboveThreshold()
	if !p(150) || p(50) {
		t.Errorf("predicate works incorrectly")
	}
}
