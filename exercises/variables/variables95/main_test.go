// variables95
// Make the tests pass!

// I AM NOT DONE
//
// This function must increment the value and return it. The code does not compile.
// := needs at least one new variable on the left.
package main_test

import "testing"

func addTen(start int) int {
	total := start
	total := total + 10
	return total
}

func TestAddTen(t *testing.T) {
	if got := addTen(5); got != 15 {
		t.Errorf("addTen(5) = %d, want 15", got)
	}
}
