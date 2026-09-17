// range88
// Make the tests pass!

// I AM NOT DONE
//
// weightedAvg computes a weighted average over (score, weight) pairs.
// Practices a range over a slice of pairs with two accumulators.
package main_test

import "testing"

type Grade struct{ Score, Weight float64 }

func weightedAvg(gs []Grade) float64 {
	var sum, weights float64
	for _, g := range gs {
		sum += g.Score
		weights++
	}
	return sum / weights
}

func TestWeightedAvg(t *testing.T) {
	if got := weightedAvg([]Grade{{5, 3}, {2, 1}}); got != 4.25 {
		t.Errorf("weightedAvg = %v, want 4.25", got)
	}
	if got := weightedAvg(nil); got != 0 {
		t.Errorf("weightedAvg(nil) = %v", got)
	}
}
