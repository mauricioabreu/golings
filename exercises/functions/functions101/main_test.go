// functions101
// Make the tests pass!

// I AM NOT DONE
//
// fibGen must return a function that hands out the next Fibonacci number
// on every call: 0, 1, 1, 2, 3, 5 and so on.
// Practices closures that keep state between calls.
package main_test

import "testing"

func fibGen() func() int {
	return func() int {
		a, b := 0, 1
		a, b = b, a+b
		return a
	}
}

func TestFibGen(t *testing.T) {
	next := fibGen()
	for i, want := range []int{0, 1, 1, 2, 3, 5, 8} {
		if got := next(); got != want {
			t.Errorf("call %d = %d, want %d", i, got, want)
		}
	}
}
