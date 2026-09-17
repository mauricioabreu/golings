// arrays17
// Make the tests pass!

// I AM NOT DONE
//
// firstOr returns the first element of an array through a pointer, or def when the pointer is nil.
// len of a nil array pointer does not panic, while indexing it does.
// len(p) for a *[N]T is the constant N, even when p == nil.
package main_test

import "testing"

func firstOr(p *[3]int, def int) int {
	if len(p) > 0 {
		return p[0]
	}
	return def
}

func TestFirstOr(t *testing.T) {
	a := [3]int{7, 8, 9}
	if got := firstOr(&a, -1); got != 7 {
		t.Errorf("firstOr(&a) = %d", got)
	}
	if got := firstOr(nil, -1); got != -1 {
		t.Errorf("firstOr(nil) = %d, want -1", got)
	}
}
