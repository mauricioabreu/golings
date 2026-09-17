// functions29
// Make the tests pass!

// I AM NOT DONE
//
// counted(f) must return a wrapper around f plus a function reporting the call count.
// Practices two functions sharing one piece of state.
package main_test

import "testing"

func counted(f func(int) int) (func(int) int, func() int) {
	calls := 0
	wrapped := func(x int) int {
		return f(x)
	}
	count := func() int { return 0 }
	return wrapped, count
}

func TestCounted(t *testing.T) {
	neg, count := counted(func(x int) int { return -x })
	neg(1)
	neg(2)
	if got := neg(3); got != -3 {
		t.Errorf("neg(3) = %d, want -3", got)
	}
	if got := count(); got != 3 {
		t.Errorf("count() = %d, want 3", got)
	}
}
