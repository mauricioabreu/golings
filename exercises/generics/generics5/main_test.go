// generics5
// Make the tests pass!

// I AM NOT DONE
//
// Max returns the larger of two values of an ordered type.
// Practices a constraint of your own with a union of types.
package main_test

import "testing"

type Ordered interface {
	~int | ~int64 | ~float64 | ~string
}

func Max[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func TestMax(t *testing.T) {
	if Max(3, 7) != 7 || Max("go", "c") != "go" || Max(2.5, -1.0) != 2.5 {
		t.Errorf("Max works incorrectly")
	}
}
