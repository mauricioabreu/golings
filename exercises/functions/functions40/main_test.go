// functions40
// Make the tests pass!

// I AM NOT DONE
//
// isTimeout must recognize ErrTimeout even when the error is wrapped.
// Practices errors.Is instead of a comparison with ==.
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

var ErrTimeout = errors.New("timeout")

func isTimeout(err error) bool {
	return err == ErrTimeout
}

func TestIsTimeout(t *testing.T) {
	wrapped := fmt.Errorf("call api: %w", ErrTimeout)
	if !isTimeout(ErrTimeout) || !isTimeout(wrapped) {
		t.Errorf("isTimeout should detect direct and wrapped ErrTimeout")
	}
	if isTimeout(errors.New("timeout")) {
		t.Errorf("isTimeout should not match a different error with the same text")
	}
}
