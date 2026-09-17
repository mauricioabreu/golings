// range93
// Make the tests pass!

// I AM NOT DONE
//
// countPositive counts the positive elements of an array.
// Practices a range over an array.
package main_test

import "testing"

func countPositive(a [5]int) int {
	n := 0
	for _, v := range a {
		if v >= 0 {
			n++
		}
	}
	return n
}

func TestCountPositive(t *testing.T) {
	if got := countPositive([5]int{-1, 0, 3, 4, -5}); got != 2 {
		t.Errorf("countPositive = %d, want 2", got)
	}
}
