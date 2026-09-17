// primitive_types72
// Make the tests pass!

// I AM NOT DONE
//
// elapsedMs returns the difference between two moments in milliseconds.
// For an interval longer than 25 days the result turns negative.
// An int32 only holds about 2.1 billion, which is some 24.8 days in milliseconds.
package main_test

import (
	"testing"
	"time"
)

func elapsedMs(from, to time.Time) int32 {
	return int32(to.Sub(from).Milliseconds())
}

func TestElapsedMs(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if got := elapsedMs(start, start.Add(1500*time.Millisecond)); got != 1500 {
		t.Errorf("elapsedMs(1.5s) = %d", got)
	}
	if got := elapsedMs(start, start.Add(30*24*time.Hour)); got != 2592000000 {
		t.Errorf("elapsedMs(30 days) = %d, want 2592000000", got)
	}
}
