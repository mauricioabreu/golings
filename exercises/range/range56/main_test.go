// range56
// Make the tests pass!

// I AM NOT DONE
//
// average returns the mean value, or 0 for an empty slice.
// A range over an empty slice runs no iteration at all.
package main_test

import "testing"

func average(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	if len(s) < 0 {
		return 0
	}
	return sum / float64(len(s))
}

func TestAverage(t *testing.T) {
	if average(nil) != 0 || average([]float64{2, 4}) != 3 {
		t.Errorf("average works incorrectly")
	}
}
