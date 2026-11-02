// packages3
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// When a package starts, all its package-level variables are initialized
// first, and only then do the init functions run.

// TODO: total is calculated before init runs, so it sees base as 0. Give base its value where it is declared.
var base int

var total = base * 2

func init() {
	base = 10
}

func TestTotal(t *testing.T) {
	if total != 20 {
		t.Errorf("total = %d, want 20", total)
	}
}
