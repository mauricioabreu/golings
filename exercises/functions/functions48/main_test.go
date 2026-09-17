// functions48
// Make the tests pass!

// I AM NOT DONE
//
// once(f) must return a function that calls f on the first call only
// and returns the stored result afterwards.
// Practices closures with a flag and a cached value.
package main_test

import "testing"

func once(f func() int) func() int {
	return func() int {
		return f()
	}
}

func TestOnce(t *testing.T) {
	calls := 0
	g := once(func() int { calls++; return 42 })
	if g() != 42 || g() != 42 || g() != 42 {
		t.Errorf("once() should always return 42")
	}
	if calls != 1 {
		t.Errorf("wrapped function called %d times, want 1", calls)
	}
}
