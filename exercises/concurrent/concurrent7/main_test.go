// concurrent7
// Make the tests pass!

// I AM NOT DONE
//
// worker has to finish when the context is cancelled.
// Practices case <-ctx.Done().
package main_test

import (
	"context"
	"testing"
	"time"
)

func worker(ctx context.Context, jobs <-chan int) int {
	n := 0
	for {
		select {
		case <-make(chan struct{}):
			return n
		case <-jobs:
			n++
		}
	}
}

func TestWorker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int)
	done := make(chan int, 1)
	go func() { done <- worker(ctx, jobs) }()
	jobs <- 1
	jobs <- 2
	cancel()
	select {
	case n := <-done:
		if n != 2 {
			t.Errorf("processed %d", n)
		}
	case <-time.After(time.Second):
		t.Fatal("worker ignored cancellation")
	}
}
