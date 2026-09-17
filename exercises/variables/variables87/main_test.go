// variables87
// Make the tests pass!

// I AM NOT DONE
//
// This function must return a timeout of seconds seconds.
// Values of different types are being multiplied.
// Practices converting to the named type time.Duration.
package main_test

import (
	"testing"
	"time"
)

func timeout(seconds int) time.Duration {
	return seconds * time.Second
}

func TestTimeout(t *testing.T) {
	if got := timeout(3); got != 3*time.Second {
		t.Errorf("timeout(3) = %v, want 3s", got)
	}
}
