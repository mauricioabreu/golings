// closures5
// Make the tests pass!

package main_test

import "testing"

// multiplier returns a function that multiplies its argument by factor.
func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// counted wraps fn: the wrapped function does what fn does, and calls tells
// how many times the wrapped function has been called.
func counted(fn func(int) int) (wrapped func(int) int, calls func() int) {
	count := 0
	wrapped = func(n int) int {
		count++
		return fn(n)
	}
	calls = func() int {
		return count
	}
	return wrapped, calls
}

func TestMultiplier(t *testing.T) {
	double := multiplier(2)
	triple := multiplier(3)

	if got := double(4); got != 8 {
		t.Errorf("double(4) = %d, want 8", got)
	}
	if got := triple(4); got != 12 {
		t.Errorf("triple(4) = %d, want 12", got)
	}
}

func TestCounted(t *testing.T) {
	wrapped, calls := counted(multiplier(3))

	if got := calls(); got != 0 {
		t.Errorf("before any call, calls() = %d, want 0", got)
	}

	wrapped(1)
	if got := wrapped(2); got != 6 {
		t.Errorf("wrapped(2) = %d, want 6", got)
	}

	if got := calls(); got != 2 {
		t.Errorf("after two calls, calls() = %d, want 2", got)
	}
}
