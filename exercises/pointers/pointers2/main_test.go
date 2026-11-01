// pointers2
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// increment adds one to the number.
func increment(n int) {
	// TODO: Change the caller's number, not a copy of it.
	n++
}

// swap exchanges the two numbers.
func swap(a, b int) {
	// TODO: Swap the caller's numbers, not copies of them.
	a, b = b, a
}

func TestIncrement(t *testing.T) {
	n := 1

	increment(n)

	if n != 2 {
		t.Errorf("n should be 2 after increment, but got %d", n)
	}
}

func TestSwap(t *testing.T) {
	x, y := 1, 2

	swap(x, y)

	if x != 2 || y != 1 {
		t.Errorf("x and y should be 2 and 1 after swap, but got %d and %d", x, y)
	}
}
