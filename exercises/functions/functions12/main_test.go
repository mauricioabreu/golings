// functions12
// Make the tests pass!

// I AM NOT DONE
//
// add(a)(b)(c) must return a+b+c.
// Practices functions returning functions returning functions.
package main_test

import "testing"

func add(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b
		}
	}
}

func TestAdd(t *testing.T) {
	if got := add(1)(2)(3); got != 6 {
		t.Errorf("add(1)(2)(3) = %d, want 6", got)
	}
	plusTen := add(4)(6)
	if got := plusTen(-10); got != 0 {
		t.Errorf("add(4)(6)(-10) = %d, want 0", got)
	}
}
