// concurrent31
// Make the tests pass!

// I AM NOT DONE
//
// Two components put values in the context under the string key "id".
// The second value buries the first, and the first component reads somebody else's data.
// Context keys need an unexported type of their own.
package main_test

import (
	"context"
	"testing"
)

func withUser(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, "id", id)
}

func userID(ctx context.Context) string {
	v, _ := ctx.Value("id").(string)
	return v
}

func withTrace(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, "id", id)
}

func TestContextKeys(t *testing.T) {
	ctx := withTrace(withUser(context.Background(), "user-1"), "trace-9")
	if got := userID(ctx); got != "user-1" {
		t.Errorf("userID = %q", got)
	}
}
