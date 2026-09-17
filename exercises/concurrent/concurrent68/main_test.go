// concurrent68
// Make the tests pass!

// I AM NOT DONE
//
// poll runs a check on a ticker a given number of times and stops
// the ticker.
// Practices time.Ticker and stopping it.
package main_test

import (
	"testing"
	"time"
)

func poll(times int, check func()) {
	ticker := time.NewTicker(2 * time.Millisecond)
	for range ticker.C {
		check()
	}
}

func TestPoll(t *testing.T) {
	n := 0
	done := make(chan struct{})
	go func() {
		poll(3, func() { n++ })
		close(done)
	}()
	select {
	case <-done:
		if n != 3 {
			t.Errorf("checks = %d", n)
		}
	case <-time.After(time.Second):
		t.Fatal("poll never stopped")
	}
}
