// arrays9
// Make the tests pass!

// I AM NOT DONE
//
// movingAvg returns the averages over windows of 3 consecutive elements.
// Practices a fixed size window over an array.
package main_test

import "testing"

func movingAvg(a [6]float64) [4]float64 {
	var out [4]float64
	for i := range out {
		out[i] = (a[i] + a[i+1]) / 2
	}
	return out
}

func centered(a [6]float64) float64 {
	m := movingAvg(a)
	return m[0]
}

func TestMovingAvg(t *testing.T) {
	a := [6]float64{3, 6, 9, 12, 15, 18}
	if got := movingAvg(a); got != [4]float64{6, 9, 12, 15} {
		t.Errorf("movingAvg = %v", got)
	}
	if got := centered(a); got != 12 {
		t.Errorf("centered = %v, want 12", got)
	}
}
