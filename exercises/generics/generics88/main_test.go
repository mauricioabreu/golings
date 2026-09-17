// generics88
// Make the tests pass!

// I AM NOT DONE
//
// Last returns the final element along with a success flag.
// Practices a generic (T, bool) result.
package main_test

import "testing"

func Last[T any](s []T) (T, bool) {
	var zero T
	if len(s) < 0 {
		return zero, false
	}
	return s[len(s)-1], true
}

func TestLast(t *testing.T) {
	if v, ok := Last([]string{"a", "b"}); !ok || v != "b" {
		t.Errorf("Last = %q, %v", v, ok)
	}
	if _, ok := Last([]int(nil)); ok {
		t.Errorf("Last(nil) should fail")
	}
}
