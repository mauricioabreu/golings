// concurrent19
// Make the tests pass!

// I AM NOT DONE
//
// startWorkers has to start exactly n workers.
// Practices starting goroutines in a loop.
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func startWorkers(n int) int32 {
	var started int32
	var wg sync.WaitGroup
	for i := 1; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&started, 1)
		}()
	}
	wg.Wait()
	return started
}

func TestStartWorkers(t *testing.T) {
	if got := startWorkers(4); got != 4 {
		t.Errorf("started = %d", got)
	}
}
