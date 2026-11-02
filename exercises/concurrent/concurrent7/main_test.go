// concurrent7
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"context"
	"testing"
	"time"
)

// spin sends 0, 1, 2, ... on ticks for as long as someone listens. It has to
// return as soon as ctx is cancelled.
func spin(ctx context.Context, ticks chan<- int) {
	for i := 0; ; i++ {
		// TODO: Don't wait for a receiver forever: stop when the context is done.
		ticks <- i
	}
}

func TestSpinStopsWhenCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ticks := make(chan int)
	stopped := make(chan struct{})

	go func() {
		defer close(stopped)
		spin(ctx, ticks)
	}()

	for want := 0; want < 3; want++ {
		if got := <-ticks; got != want {
			t.Fatalf("tick = %d, want %d", got, want)
		}
	}
	cancel()

	select {
	case <-stopped:
	case <-time.After(time.Second):
		t.Fatal("spin is still running a second after the context was cancelled")
	}
}
