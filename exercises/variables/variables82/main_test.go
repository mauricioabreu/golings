// variables82
// Make the tests pass!

// I AM NOT DONE
//
// Every counter made by newCounter must count on its own.
// Right now all the counters share one value.
// Practices variable lifetime and capture in a closure.
package main_test

import "testing"

var count int

func newCounter() func() int {
	return func() int {
		count++
		return count
	}
}

func TestCounters(t *testing.T) {
	a := newCounter()
	b := newCounter()
	a()
	a()
	if got := a(); got != 3 {
		t.Errorf("a() third call = %d, want 3", got)
	}
	if got := b(); got != 1 {
		t.Errorf("b() first call = %d, want 1", got)
	}
}
