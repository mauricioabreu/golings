// variables59
// Make the tests pass!

// I AM NOT DONE
//
// The constant Timeout must be 2.5 seconds, and the function must return
// it in milliseconds as an int64.
// Practices typed time.Duration constants and the Milliseconds method.
package main_test

import (
	"testing"
	"time"
)

const Timeout = 2 * time.Second

func timeoutMs() int64 {
	return int64(Timeout.Seconds())
}

func TestTimeoutMs(t *testing.T) {
	if Timeout != 2500*time.Millisecond {
		t.Errorf("Timeout = %v, want 2.5s", Timeout)
	}
	if got := timeoutMs(); got != 2500 {
		t.Errorf("timeoutMs() = %d, want 2500", got)
	}
}
