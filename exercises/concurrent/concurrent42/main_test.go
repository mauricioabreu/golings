// concurrent42
// Make the tests pass!

// I AM NOT DONE
//
// Inc takes the mutex and forgets to release it; the second call hangs.
// Practices defer mu.Unlock().
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.n++
}

func TestCounter(t *testing.T) {
	var c Counter
	done := make(chan struct{})
	go func() {
		c.Inc()
		c.Inc()
		close(done)
	}()
	select {
	case <-done:
		if c.n != 2 {
			t.Errorf("n = %d", c.n)
		}
	case <-time.After(time.Second):
		t.Fatal("second Inc is blocked: mutex never unlocked")
	}
}
