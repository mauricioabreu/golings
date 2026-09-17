// if51
// Make the tests pass!

// I AM NOT DONE
//
// sameRoute compares two routes, which are slices of stop names.
// The code does not compile.
// Slices can only be compared against nil.
package main_test

import "testing"

func sameRoute(a, b []string) bool {
	if a == b {
		return true
	}
	return false
}

func TestSameRoute(t *testing.T) {
	if !sameRoute([]string{"A", "B"}, []string{"A", "B"}) {
		t.Errorf("identical routes should match")
	}
	if sameRoute([]string{"A", "B"}, []string{"B", "A"}) || sameRoute([]string{"A"}, nil) {
		t.Errorf("different routes should not match")
	}
	if !sameRoute(nil, []string{}) {
		t.Errorf("two empty routes should match")
	}
}
