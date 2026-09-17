// variables24
// Make the tests pass!

// I AM NOT DONE
//
// The constant White must be 0xFFFFFF (16777215).
// Practices hexadecimal integer literals.
package main_test

import "testing"

const White = 0xFFFF

func TestWhite(t *testing.T) {
	if White != 16777215 {
		t.Errorf("White = %d, want 16777215", White)
	}
}
