// concurrent102
// Make the tests pass!

// I AM NOT DONE
//
// countHits increments a counter from many goroutines without synchronization.
// Practices sync/atomic.
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func countHits(n int) int64 {
	var hits int64
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			hits++
		}()
	}
	wg.Wait()
	return atomic.LoadInt64(&hits)
}

func TestCountHits(t *testing.T) {
	if got := countHits(200); got != 200 {
		t.Errorf("hits = %d", got)
	}
}
