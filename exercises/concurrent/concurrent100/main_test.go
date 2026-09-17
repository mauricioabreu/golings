// concurrent100
// Make the tests pass!

// I AM NOT DONE
//
// emit sends events to a channel; when the receiver is gone and the context
// is cancelled, the send blocks forever.
// Practices that blocking sends must also listen to ctx.Done().
package main_test

import (
	"context"
	"testing"
	"time"
)

func emit(ctx context.Context, out chan<- int, n int) error {
	for i := 0; i < n; i++ {
		out <- i
	}
	return nil
}

func TestEmit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	out := make(chan int)
	res := make(chan error, 1)
	go func() { res <- emit(ctx, out, 10) }()
	<-out
	<-out
	cancel()
	select {
	case err := <-res:
		if err != context.Canceled {
			t.Errorf("err = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("emit is blocked after cancellation")
	}
}
