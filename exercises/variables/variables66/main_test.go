// variables66
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the discount: 10 for a total of 1000 or more, 0 otherwise.
// The variable is declared inside the if block and is out of reach outside it.
// Practices variable scope.
package main_test

import "testing"

func discount(total int) int {
	if total >= 1000 {
		d := 10
	}
	return d
}

func TestDiscount(t *testing.T) {
	if got := discount(1500); got != 10 {
		t.Errorf("discount(1500) = %d, want 10", got)
	}
	if got := discount(999); got != 0 {
		t.Errorf("discount(999) = %d, want 0", got)
	}
}
