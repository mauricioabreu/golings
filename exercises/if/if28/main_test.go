// if28
// Make the tests pass!

// I AM NOT DONE
//
// verbose takes a *bool, where nil means "unset" and counts as false.
// Right now it panics on nil.
// Practices dereferencing a nil pointer in a condition.
package main_test

import "testing"

func verbose(flag *bool) string {
	if *flag {
		return "loud"
	}
	return "quiet"
}

func TestVerbose(t *testing.T) {
	yes, no := true, false
	if verbose(&yes) != "loud" || verbose(&no) != "quiet" {
		t.Errorf("verbose with explicit flag is wrong")
	}
	if got := verbose(nil); got != "quiet" {
		t.Errorf("verbose(nil) = %q, want quiet", got)
	}
}
