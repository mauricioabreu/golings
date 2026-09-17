// maps30
// Make the tests pass!

// I AM NOT DONE
//
// countReadings counts how often every reading occurs; NaN, meaning "no data",
// has to be counted under one key.
// Right now every NaN creates a new key and none of them can be read back.
// NaN != NaN, so NaN keys in a map are distinct and unreachable.
package main_test

import (
	"math"
	"testing"
)

func countReadings(vals []float64) (map[float64]int, int) {
	m := map[float64]int{}
	missing := 0
	for _, v := range vals {
		m[v]++
	}
	return m, missing
}

func TestCountReadings(t *testing.T) {
	nan := math.NaN()
	m, missing := countReadings([]float64{1.5, nan, 1.5, nan, nan})
	if m[1.5] != 2 || missing != 3 || len(m) != 1 {
		t.Errorf("m = %v, missing = %d", m, missing)
	}
}
