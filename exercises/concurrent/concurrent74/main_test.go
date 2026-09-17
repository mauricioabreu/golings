// concurrent74
// Make the tests pass!

// I AM NOT DONE
//
// handle passes a request identifier through the context deep into
// the calls.
// Practices context.WithValue and passing ctx down the stack.
package main_test

import (
	"context"
	"testing"
)

type key int

const userKey key = 1

func audit(ctx context.Context) string {
	u, _ := ctx.Value(userKey).(string)
	return "audit:" + u
}

func handle(ctx context.Context, user string) string {
	context.WithValue(ctx, userKey, user)
	return audit(context.Background())
}

func TestHandle(t *testing.T) {
	if got := handle(context.Background(), "ann"); got != "audit:ann" {
		t.Errorf("handle = %q", got)
	}
}
