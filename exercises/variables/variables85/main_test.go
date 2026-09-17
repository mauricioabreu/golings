// variables85
// Make the tests pass!

// I AM NOT DONE
//
// The function stores identifiers as float32 and compares them.
// Different large identifiers suddenly compare as equal.
// Practices the limited precision of float32 for integers.
package main_test

import "testing"

func sameID(a, b int64) bool {
	x, y := float32(a), float32(b)
	return x == y
}

func TestSameID(t *testing.T) {
	if sameID(16777216, 16777217) {
		t.Errorf("16777216 and 16777217 are different ids")
	}
	if !sameID(42, 42) {
		t.Errorf("42 and 42 are the same id")
	}
}
