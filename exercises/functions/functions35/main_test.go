// functions35
// Make the tests pass!

// I AM NOT DONE
//
// safeCall must run f and turn a panic into an error.
// Right now the panic flies straight through.
// Practices defer, recover and a named result.
package main_test

import (
	"fmt"
	"testing"
)

func safeCall(f func()) error {
	f()
	return nil
}

func TestSafeCall(t *testing.T) {
	_ = fmt.Sprint
	if err := safeCall(func() {}); err != nil {
		t.Errorf("safeCall(ok) = %v, want nil", err)
	}
	err := safeCall(func() { panic("boom") })
	if err == nil || err.Error() != "panic: boom" {
		t.Errorf("safeCall(panic) = %v, want panic: boom", err)
	}
}
