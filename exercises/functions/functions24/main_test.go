// functions24
// Make the tests pass!

// I AM NOT DONE
//
// maxOf must return the largest of its arguments, of which there is at least one.
// Practices variadic functions with a mandatory first parameter.
package main_test

import "testing"

func maxOf(first int, rest ...int) int {
	m := first
	for _, v := range rest {
		if v < m {
			m = v
		}
	}
	return m
}

func TestMaxOf(t *testing.T) {
	if got := maxOf(3, 9, 2); got != 9 {
		t.Errorf("maxOf(3, 9, 2) = %d, want 9", got)
	}
	if got := maxOf(-5); got != -5 {
		t.Errorf("maxOf(-5) = %d, want -5", got)
	}
}
