// generics95
// Make the tests pass!

// I AM NOT DONE
//
// SwapPtr swaps the values behind two pointers.
// Practices generic pointers *T.
package main_test

import "testing"

func SwapPtr[T any](a, b *T) {
	a, b = b, a
}

func TestSwapPtr(t *testing.T) {
	x, y := "left", "right"
	SwapPtr(&x, &y)
	if x != "right" || y != "left" {
		t.Errorf("x=%s y=%s", x, y)
	}
}
