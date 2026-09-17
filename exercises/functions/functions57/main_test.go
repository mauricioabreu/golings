// functions57
// Make the tests pass!

// I AM NOT DONE
//
// safely must recover from a panic and return false.
// recover is called in a helper function, so the panic is not caught.
// recover only works when the deferred function calls it directly.
package main_test

import "testing"

func tryRecover(ok *bool) {
	if r := recover(); r != nil {
		*ok = false
	}
}

func safely(f func()) (ok bool) {
	ok = true
	defer func() {
		tryRecover(&ok)
	}()
	f()
	return ok
}

func TestSafely(t *testing.T) {
	if !safely(func() {}) {
		t.Errorf("safely(no panic) = false, want true")
	}
	if safely(func() { panic("oops") }) {
		t.Errorf("safely(panic) = true, want false")
	}
}
