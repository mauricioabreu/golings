// concurrent30
// Make the tests pass!

// I AM NOT DONE
//
// slow has to return context.DeadlineExceeded when the context has run out.
// Practices ctx.Err().
package main_test

import (
	"context"
	"errors"
	"testing"
	"time"
)

func slow(ctx context.Context) error {
	select {
	case <-time.After(time.Second):
		return nil
	case <-ctx.Done():
		return nil
	}
}

func TestSlow(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	if err := slow(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v", err)
	}
}
