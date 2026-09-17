// anonymous_functions81
// Make the tests pass!

// I AM NOT DONE
//
// debounce returns a literal: a call goes through only when at least wait has
// passed since the previous one. The time is passed in explicitly.
// Practices a closure remembering the moment of the last call.
package main_test

import (
	"testing"
	"time"
)

func debounce(wait time.Duration) func(now time.Time) bool {
	var last time.Time
	return func(now time.Time) bool {
		fire := now.Sub(last) > wait
		return fire
	}
}

func TestDebounce(t *testing.T) {
	d := debounce(time.Second)
	t0 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	at := func(ms int) time.Time { return t0.Add(time.Duration(ms) * time.Millisecond) }
	got := []bool{d(at(0)), d(at(500)), d(at(1200)), d(at(2200))}
	want := []bool{true, false, false, true}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("call %d = %v, want %v", i, got[i], want[i])
		}
	}
}
