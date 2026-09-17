// anonymous_functions15
// Make the tests pass!

// I AM NOT DONE
//
// table holds the factorial of 5, computed by a literal with a loop inside.
// Practices an IIFE for initializing a value.
package main_test

import "testing"

var fact5 = func() int {
	r := 1
	for i := 2; i < 5; i++ {
		r *= i
	}
	return r
}()

func TestFact5(t *testing.T) {
	if fact5 != 120 {
		t.Errorf("fact5 = %d, want 120", fact5)
	}
}
