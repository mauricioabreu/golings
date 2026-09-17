// generics80
// Make the tests pass!

// I AM NOT DONE
//
// Percent computes part as a percentage of whole, decimals included.
// For integer types the fractional part is lost.
// Operations on a T happen in the type T.
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func Percent[T Number](part, whole T) float64 {
	return float64(part / whole * 100)
}

func TestPercent(t *testing.T) {
	if got := Percent(1, 8); got != 12.5 {
		t.Errorf("Percent(1, 8) = %v, want 12.5", got)
	}
	if got := Percent(1.0, 4.0); got != 25 {
		t.Errorf("Percent(1.0, 4.0) = %v", got)
	}
}
