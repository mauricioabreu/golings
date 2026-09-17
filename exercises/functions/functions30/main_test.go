// functions30
// Make the tests pass!

// I AM NOT DONE
//
// run turns any panic into an error. A panic may carry an error or a string.
// With a string the function panics inside its own defer.
// recover returns an interface{}; the type of the value has to be checked.
package main_test

import (
	"fmt"
	"testing"
)

func run(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	f()
	return nil
}

func TestRun(t *testing.T) {
	diskErr := fmt.Errorf("disk full")
	if err := run(func() { panic(diskErr) }); err != diskErr {
		t.Errorf("run(error panic) = %v, want disk full", err)
	}
	if err := run(func() { panic("oops") }); err == nil || err.Error() != "panic: oops" {
		t.Errorf("run(string panic) = %v, want panic: oops", err)
	}
}
