// functions23
// Make the tests pass!

// I AM NOT DONE
//
// timed must run f and store the difference of the now readings in *elapsed,
// even when f panics. The clock is passed in as a function to keep this deterministic.
// Practices a defer with a closure for measuring a duration.
package main_test

import (
	"testing"
	"time"
)

func timed(now func() time.Time, elapsed *time.Duration, f func()) {
	start := now()
	f()
	*elapsed = now().Sub(start)
}

func TestTimed(t *testing.T) {
	base := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tick := 0
	clock := func() time.Time {
		tick++
		return base.Add(time.Duration(tick) * time.Second)
	}
	var d time.Duration
	timed(clock, &d, func() { clock(); clock() })
	if d != 3*time.Second {
		t.Errorf("elapsed = %v, want 3s", d)
	}
	var p time.Duration
	func() {
		defer func() { _ = recover() }()
		timed(clock, &p, func() { clock(); panic("fail") })
	}()
	if p != 2*time.Second {
		t.Errorf("elapsed after panic = %v, want 2s", p)
	}
}
