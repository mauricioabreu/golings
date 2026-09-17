// functions81
// Make the tests pass!

// I AM NOT DONE
//
// statusOf must pull the code out of an *HTTPError somewhere in the chain.
// When there is no such error it returns 0.
// Practices errors.As and custom error types.
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

type HTTPError struct{ Code int }

func (e *HTTPError) Error() string { return fmt.Sprintf("http %d", e.Code) }

func statusOf(err error) int {
	if he, ok := err.(*HTTPError); ok {
		return he.Code
	}
	return 0
}

func TestStatusOf(t *testing.T) {
	err := fmt.Errorf("fetch: %w", &HTTPError{Code: 404})
	if got := statusOf(err); got != 404 {
		t.Errorf("statusOf(wrapped 404) = %d, want 404", got)
	}
	if got := statusOf(errors.New("boom")); got != 0 {
		t.Errorf("statusOf(boom) = %d, want 0", got)
	}
}
