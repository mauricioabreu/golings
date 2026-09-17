// anonymous_functions26
// Make the tests pass!

// I AM NOT DONE
//
// withContext returns a function adding context to an error.
// Practices a literal returning an error.
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

func withContext(ctx string) func(error) error {
	return func(err error) error {
		return fmt.Errorf("%s: %v", ctx, err)
	}
}

func TestWithContext(t *testing.T) {
	base := errors.New("timeout")
	err := withContext("db")(base)
	if !errors.Is(err, base) || err.Error() != "db: timeout" {
		t.Errorf("err = %v", err)
	}
}
