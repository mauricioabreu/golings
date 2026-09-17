// concurrent16
// Make the tests pass!

// I AM NOT DONE
//
// callWithTimeout waits for a result no longer than a given time.
// A slow call must not hold it up.
// Practices a select between the result and a timer, plus a buffer for the result.
package main_test

import (
	"errors"
	"testing"
	"time"
)

func callWithTimeout(f func() int, d time.Duration) (int, error) {
	return f(), nil
}

func TestCallWithTimeout(t *testing.T) {
	if v, err := callWithTimeout(func() int { return 5 }, time.Second); err != nil || v != 5 {
		t.Errorf("fast call = %d, %v", v, err)
	}
	start := time.Now()
	_, err := callWithTimeout(func() int { time.Sleep(500 * time.Millisecond); return 1 }, 20*time.Millisecond)
	if err == nil || time.Since(start) > 300*time.Millisecond {
		t.Errorf("slow call: err=%v after %v", err, time.Since(start))
	}
}
