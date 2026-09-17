// concurrent89
// Make the tests pass!

// I AM NOT DONE
//
// retry repeats the operation until it succeeds or the context is cancelled.
// Practices a select between waiting and ctx.Done().
package main_test

import (
	"context"
	"errors"
	"testing"
	"time"
)

func retry(ctx context.Context, op func() error) error {
	for {
		if err := op(); err == nil {
			return nil
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func TestRetry(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	done := make(chan error, 1)
	go func() { done <- retry(ctx, func() error { return errors.New("down") }) }()
	select {
	case err := <-done:
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("err = %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("retry ignored the context")
	}
}
