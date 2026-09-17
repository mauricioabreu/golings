// variables49
// Make the tests pass!

// I AM NOT DONE
//
// hours must convert a number of minutes into (fractional) hours.
// Practices time.Duration, the time.Minute constant and the Hours method.
package main_test

import (
	"testing"
	"time"
)

func hours(minutes int) float64 {
	d := time.Duration(minutes) * time.Hour
	return d.Minutes()
}

func TestHours(t *testing.T) {
	cases := map[int]float64{90: 1.5, 60: 1, 15: 0.25, 0: 0}
	for in, want := range cases {
		if got := hours(in); got != want {
			t.Errorf("hours(%d) = %v, want %v", in, got, want)
		}
	}
}
