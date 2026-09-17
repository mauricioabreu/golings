// variables72
// Make the tests pass!

// I AM NOT DONE
//
// The constant Million must be one million.
// Somebody miscounted while typing the number and digits are missing.
// Practices numeric literals with the _ separator.
package main_test

import "testing"

const Million = 1_000_00

func TestMillion(t *testing.T) {
	if Million != 1000000 {
		t.Errorf("Million = %d, want 1000000", Million)
	}
}
