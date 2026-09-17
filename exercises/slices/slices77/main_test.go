// slices77
// Make the tests pass!

// I AM NOT DONE
//
// equal compares two string slices element by element.
// Practices comparing slices by hand.
package main_test

import "testing"

func equal(a, b []string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return len(a) > 0
}

func TestEqual(t *testing.T) {
	if !equal([]string{"a", "b"}, []string{"a", "b"}) || !equal(nil, []string{}) {
		t.Errorf("equal slices should match")
	}
	if equal([]string{"a"}, []string{"a", "b"}) || equal([]string{"a", "b"}, []string{"a"}) {
		t.Errorf("slices of different length should not match")
	}
}
