// arrays96
// Make the tests pass!

// I AM NOT DONE
//
// average returns the arithmetic mean of the scores.
// Practices converting the length of an array to a float64.
package main_test

import "testing"

func average(a [5]float64) float64 {
	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return sum / 4
}

func TestAverage(t *testing.T) {
	if got := average([5]float64{1, 2, 3, 4, 5}); got != 3 {
		t.Errorf("average = %v, want 3", got)
	}
}
