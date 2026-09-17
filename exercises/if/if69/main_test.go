// if69
// Make the tests pass!

// I AM NOT DONE
//
// tooSlow must return true when a request took longer than 5 seconds.
// Right now almost every request counts as slow.
// In a comparison against a Duration the untyped constant 5 means 5 nanoseconds.
package main_test

import (
	"testing"
	"time"
)

func tooSlow(d time.Duration) bool {
	if d > 5 {
		return true
	}
	return false
}

func TestTooSlow(t *testing.T) {
	if tooSlow(300 * time.Millisecond) {
		t.Errorf("300ms is not slow")
	}
	if tooSlow(5 * time.Second) {
		t.Errorf("exactly 5s is not slow")
	}
	if !tooSlow(6 * time.Second) {
		t.Errorf("6s is slow")
	}
}
