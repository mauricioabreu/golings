// packages3
// Make the tests pass!

package main_test

import "testing"

// When a package starts, all its package-level variables are initialized
// first, and only then do the init functions run.

var base = 10

var total = base * 2

func TestTotal(t *testing.T) {
	if total != 20 {
		t.Errorf("total = %d, want 20", total)
	}
}
