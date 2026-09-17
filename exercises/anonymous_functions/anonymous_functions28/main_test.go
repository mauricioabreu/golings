// anonymous_functions28
// Make the tests pass!

// I AM NOT DONE
//
// stats returns an add function and a report function (min, max, average)
// sharing one piece of state.
// Practices several closures over shared state.
package main_test

import "testing"

func stats() (add func(float64), report func() (min, max, mean float64)) {
	var lo, hi, sum float64
	n := 0
	add = func(v float64) {
		if v < lo {
			lo = v
		}
		if v > hi {
			hi = v
		}
		sum += v
		n++
	}
	report = func() (float64, float64, float64) {
		if n == 0 {
			return 0, 0, 0
		}
		return lo, hi, sum / float64(n)
	}
	return
}

func TestStats(t *testing.T) {
	add, report := stats()
	for _, v := range []float64{4, 8, 6} {
		add(v)
	}
	if lo, hi, mean := report(); lo != 4 || hi != 8 || mean != 6 {
		t.Errorf("report = %v %v %v", lo, hi, mean)
	}
}
