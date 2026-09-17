// concurrent47
// Make the tests pass!

// I AM NOT DONE
//
// stopAll has to cancel the context so the worker finishes.
// Practices context.WithCancel.
package main_test

import (
	"context"
	"testing"
	"time"
)

func stopAll() bool {
	ctx, cancel := context.WithCancel(context.Background())
	stopped := make(chan struct{})
	go func() {
		<-ctx.Done()
		close(stopped)
	}()
	_ = cancel
	select {
	case <-stopped:
		return true
	case <-time.After(time.Second):
		return false
	}
}

func TestStopAll(t *testing.T) {
	if !stopAll() {
		t.Errorf("worker was not stopped")
	}
}
