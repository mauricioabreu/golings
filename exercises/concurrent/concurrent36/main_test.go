// concurrent36
// Make the tests pass!

// I AM NOT DONE
//
// fetch starts a child operation with a timeout, and builds its context
// from context.Background, so a cancellation of the parent never reaches it.
// Derived contexts have to inherit from the parent.
package main_test

import (
	"context"
	"testing"
	"time"
)

func fetch(ctx context.Context) error {
	child, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	<-child.Done()
	return child.Err()
}

func TestFetchCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	res := make(chan error, 1)
	go func() { res <- fetch(ctx) }()
	cancel()
	select {
	case err := <-res:
		if err != context.Canceled {
			t.Errorf("err = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("cancellation did not reach the child context")
	}
}
