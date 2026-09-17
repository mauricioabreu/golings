// anonymous_functions20
// Make the tests pass!

// I AM NOT DONE
//
// multiplier(k) returns a literal returning a literal: m(k)(x)(y) = k*x*y.
// Practices nested function literals.
package main_test

import "testing"

func multiplier(k int) func(int) func(int) int {
	return func(x int) func(int) int {
		return func(y int) int {
			return k * x
		}
	}
}

func TestMultiplier(t *testing.T) {
	if multiplier(2)(3)(4) != 24 {
		t.Errorf("multiplier = %d", multiplier(2)(3)(4))
	}
}
