// variables16
// Make the tests pass!

// I AM NOT DONE
//
// The permissions rw-r--r-- are 644 in octal, which is 420 in decimal.
// The number is currently written in decimal.
// Practices octal literals written with 0o.
package main_test

import "testing"

func defaultPerm() int {
	perm := 644
	return perm
}

func TestDefaultPerm(t *testing.T) {
	if got := defaultPerm(); got != 420 {
		t.Errorf("defaultPerm() = %d, want 420 (0o644)", got)
	}
}
