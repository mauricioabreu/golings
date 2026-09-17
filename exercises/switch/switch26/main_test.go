// switch26
// Make the tests pass!

// I AM NOT DONE
//
// httpCode returns the code of a *StatusError in the error chain, or 500.
// The type switch does not see an error wrapped with %w.
// A type switch only looks at the top level of an error.
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

type StatusError struct{ Code int }

func (e *StatusError) Error() string { return fmt.Sprintf("status %d", e.Code) }

func httpCode(err error) int {
	switch e := err.(type) {
	case *StatusError:
		return e.Code
	}
	return 500
}

func TestHTTPCode(t *testing.T) {
	_ = errors.As
	if got := httpCode(&StatusError{404}); got != 404 {
		t.Errorf("httpCode(404) = %d", got)
	}
	if got := httpCode(fmt.Errorf("handler: %w", &StatusError{403})); got != 403 {
		t.Errorf("httpCode(wrapped 403) = %d, want 403", got)
	}
	if got := httpCode(errors.New("boom")); got != 500 {
		t.Errorf("httpCode(boom) = %d", got)
	}
}
