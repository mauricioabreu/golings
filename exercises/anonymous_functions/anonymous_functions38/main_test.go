// anonymous_functions38
// Make the tests pass!

// I AM NOT DONE
//
// limiter returns a literal allowing no more than n calls in a window of per.
// The time is passed in explicitly.
// Practices a closure holding window state.
package main_test

import (
	"testing"
	"time"
)

func limiter(n int, per time.Duration) func(now time.Time) bool {
	var start time.Time
	count := 0
	return func(now time.Time) bool {
		count++
		return count <= n
	}
}

func TestLimiter(t *testing.T) {
	allow := limiter(2, time.Second)
	t0 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	got := []bool{allow(t0), allow(t0), allow(t0.Add(500 * time.Millisecond)), allow(t0.Add(time.Second)), allow(t0.Add(1100 * time.Millisecond))}
	want := []bool{true, true, false, true, true}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("call %d = %v, want %v", i, got[i], want[i])
		}
	}
}
