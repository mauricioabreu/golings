// variables86
// Make the tests pass!

// I AM NOT DONE
//
// pageSize takes a *int, where nil means "use 20".
// Right now it panics on nil.
// Practices the zero value of a pointer and the nil check.
package main_test

import "testing"

func pageSize(size *int) int {
	return *size
}

func TestPageSize(t *testing.T) {
	if got := pageSize(nil); got != 20 {
		t.Errorf("pageSize(nil) = %d, want 20", got)
	}
	n := 50
	if got := pageSize(&n); got != 50 {
		t.Errorf("pageSize(&50) = %d, want 50", got)
	}
}
