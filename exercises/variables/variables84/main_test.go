// variables84
// Make the tests pass!

// I AM NOT DONE
//
// This function must pass the address of the maxRetries setting to a function that changes it.
// The code does not compile: a constant has no address.
// Practices the difference between const and var.
package main_test

import "testing"

const maxRetries = 3

func bump(p *int) {
	*p++
}

func bumpRetries() int {
	bump(&maxRetries)
	return maxRetries
}

func TestBumpRetries(t *testing.T) {
	if got := bumpRetries(); got != 4 {
		t.Errorf("bumpRetries() = %d, want 4", got)
	}
}
