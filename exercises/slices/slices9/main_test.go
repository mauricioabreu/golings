// slices9
// Make the tests pass!

// I AM NOT DONE
//
// minOf returns the smallest element of a non-empty slice.
// Practices comparing elements while walking a slice.
package main_test

import "testing"

func minOf(s []float64) float64 {
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func TestMinOf(t *testing.T) {
	if got := minOf([]float64{2.5, -1, 7}); got != -1 {
		t.Errorf("minOf = %v", got)
	}
}
