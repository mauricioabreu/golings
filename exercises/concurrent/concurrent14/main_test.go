// concurrent14
// Make the tests pass!

// I AM NOT DONE
//
// The server waits for a stop signal on a done channel, and nobody closes the channel.
// close(done) as a broadcast signal.
package main_test

import (
	"sync"
	"testing"
	"time"
)

func shutdown(workers int) bool {
	done := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-done
		}()
	}
	_ = done
	finished := make(chan struct{})
	go func() { wg.Wait(); close(finished) }()
	select {
	case <-finished:
		return true
	case <-time.After(time.Second):
		return false
	}
}

func TestShutdown(t *testing.T) {
	if !shutdown(3) {
		t.Errorf("workers did not stop")
	}
}
