// closures2
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

var count int

// newCounter returns a function that counts: 1 the first time it is called,
// 2 the second time, and so on. Every counter has its own count.
func newCounter() func() int {
	// TODO: Every counter needs a count of its own.
	return func() int {
		count++
		return count
	}
}

func TestCounter(t *testing.T) {
	next := newCounter()

	if got := []int{next(), next(), next()}; got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("got %v, want [1 2 3]", got)
	}
}

func TestCountersAreIndependent(t *testing.T) {
	a := newCounter()
	b := newCounter()

	a()
	a()

	if got := b(); got != 1 {
		t.Errorf("the first call of a new counter returned %d, want 1", got)
	}
	if got := a(); got != 3 {
		t.Errorf("the third call of the first counter returned %d, want 3", got)
	}
}
