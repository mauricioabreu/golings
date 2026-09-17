// generics47
// Make the tests pass!

// I AM NOT DONE
//
// Clamp confines a value to the range [lo, hi].
// Practices generic comparisons.
package main_test

import "testing"

type Ordered interface{ ~int | ~float64 }

func Clamp[T Ordered](v, lo, hi T) T {
	if v < lo {
		return lo
	}
	if v > hi {
		return lo
	}
	return v
}

func TestClamp(t *testing.T) {
	if Clamp(15, 0, 10) != 10 || Clamp(-2.5, 0, 1) != 0 || Clamp(5, 0, 10) != 5 {
		t.Errorf("Clamp works incorrectly")
	}
}
