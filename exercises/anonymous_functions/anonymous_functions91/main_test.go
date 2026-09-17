// anonymous_functions91
// Make the tests pass!

// I AM NOT DONE
//
// safeGo runs a function in a goroutine and has to turn its panic into an error.
// A recover in the calling function does not catch a panic from another goroutine.
// recover only works in the goroutine the panic happened in.
package main_test

import (
	"fmt"
	"testing"
)

func safeGo(f func()) (err error) {
	done := make(chan error, 1)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	go func() {
		f()
		done <- nil
	}()
	return <-done
}

func TestSafeGo(t *testing.T) {
	if err := safeGo(func() {}); err != nil {
		t.Errorf("safeGo(ok) = %v", err)
	}
	if err := safeGo(func() { panic("boom") }); err == nil || err.Error() != "panic: boom" {
		t.Errorf("safeGo(panic) = %v", err)
	}
}
