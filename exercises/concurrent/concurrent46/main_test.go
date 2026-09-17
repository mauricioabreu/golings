// concurrent46
// Make the tests pass!

// I AM NOT DONE
//
// wg.Add is called inside the goroutine that was started, and Wait returns before
// the work has even begun.
// Add has to be called before the goroutine starts.
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func runJobs(n int, gate <-chan struct{}) int32 {
	var done int32
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		go func() {
			<-gate
			wg.Add(1)
			defer wg.Done()
			atomic.AddInt32(&done, 1)
		}()
	}
	finished := make(chan struct{})
	go func() {
		wg.Wait()
		close(finished)
	}()
	<-finished
	return atomic.LoadInt32(&done)
}

func TestRunJobs(t *testing.T) {
	gate := make(chan struct{})
	res := make(chan int32, 1)
	go func() { res <- runJobs(5, gate) }()
	time.Sleep(50 * time.Millisecond)
	close(gate)
	if got := <-res; got != 5 {
		t.Errorf("done = %d, want 5", got)
	}
}
