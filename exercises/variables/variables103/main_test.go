// variables103
// Make the tests pass!

// I AM NOT DONE
//
// swap must exchange two values and return them.
// Practices multiple assignment in Go.
package main_test

import "testing"

func swap(a, b string) (string, string) {
	a = b
	return a, b
}

func TestSwap(t *testing.T) {
	x, y := swap("left", "right")
	if x != "right" || y != "left" {
		t.Errorf("swap(left, right) = (%s, %s), want (right, left)", x, y)
	}
}
