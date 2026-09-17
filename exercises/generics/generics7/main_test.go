// generics7
// Make the tests pass!

// I AM NOT DONE
//
// Clear replaces every element of a slice with the zero value of its type.
// Practices var zero T.
package main_test

import "testing"

func Clear[T any](s []T) {
	var zero T
	for i := range s {
		s[i] = s[0]
	}
}

func TestClear(t *testing.T) {
	a := []string{"x", "y"}
	Clear(a)
	b := []*int{new(int)}
	Clear(b)
	if a[0] != "" || a[1] != "" || b[0] != nil {
		t.Errorf("Clear works incorrectly: %q %v", a, b)
	}
}
