// if98
// Make the tests pass!

// I AM NOT DONE
//
// isZero must check that a number is zero. The code does not compile.
// An if condition needs the comparison operator ==.
package main_test

import "testing"

func isZero(x int) bool {
	if x = 0 {
		return true
	}
	return false
}

func TestIsZero(t *testing.T) {
	if !isZero(0) || isZero(3) {
		t.Errorf("isZero works incorrectly")
	}
}
