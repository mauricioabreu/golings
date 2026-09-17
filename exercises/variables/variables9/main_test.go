// variables9
// Make the tests pass!

// I AM NOT DONE
//
// This function must turn a fractional number of seconds (1.5, say) into a time.Duration.
// Right now the fractional part is lost.
// A time.Duration is an integer number of nanoseconds.
package main_test

import (
	"testing"
	"time"
)

func fromSeconds(secs float64) time.Duration {
	return time.Duration(secs) * time.Second
}

func TestFromSeconds(t *testing.T) {
	if got := fromSeconds(1.5); got != 1500*time.Millisecond {
		t.Errorf("fromSeconds(1.5) = %v, want 1.5s", got)
	}
	if got := fromSeconds(0.25); got != 250*time.Millisecond {
		t.Errorf("fromSeconds(0.25) = %v, want 250ms", got)
	}
}
