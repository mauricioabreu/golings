// anonymous_functions96
// Make the tests pass!

// I AM NOT DONE
//
// startJob starts the work in the background, and the function loading the argument
// runs in the current goroutine and blocks the caller.
// In go f(x()) the call x() happens before the goroutine starts.
package main_test

import (
	"testing"
	"time"
)

func startJob(load func() int, results chan<- int) {
	process := func(v int) { results <- v * 2 }
	go process(load())
}

func TestStartJob(t *testing.T) {
	release := make(chan struct{})
	results := make(chan int, 1)
	load := func() int { <-release; return 21 }
	started := make(chan struct{})
	go func() {
		startJob(load, results)
		close(started)
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		close(release)
		t.Fatal("startJob blocked the caller")
	}
	close(release)
	if got := <-results; got != 42 {
		t.Errorf("result = %d", got)
	}
}
