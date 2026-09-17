// functions80
// Make the tests pass!

// I AM NOT DONE
//
// compose(f, g) must return the function x -> f(g(x)).
// Practices function composition.
package main_test

import "testing"

func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return g(f(x))
	}
}

func TestCompose(t *testing.T) {
	inc := func(x int) int { return x + 1 }
	dbl := func(x int) int { return x * 2 }
	if got := compose(inc, dbl)(5); got != 11 {
		t.Errorf("compose(inc, dbl)(5) = %d, want 11", got)
	}
}
