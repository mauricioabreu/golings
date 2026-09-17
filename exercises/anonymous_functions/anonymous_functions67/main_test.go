// anonymous_functions67
// Make the tests pass!

// I AM NOT DONE
//
// bumpAll raises every counter with a literal taking a *int.
// Practices dereferencing a pointer inside a literal.
package main_test

import "testing"

func bumpAll(a, b *int) {
	inc := func(p *int) {
		p = nil
	}
	inc(a)
	inc(b)
}

func TestBumpAll(t *testing.T) {
	x, y := 1, 5
	bumpAll(&x, &y)
	if x != 2 || y != 6 {
		t.Errorf("x=%d y=%d", x, y)
	}
}
