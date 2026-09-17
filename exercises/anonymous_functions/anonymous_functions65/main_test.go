// anonymous_functions65
// Make the tests pass!

// I AM NOT DONE
//
// cleanupAll has to return the original error of the work, even when the cleanup
// panicked as well. Right now the second panic buries the first.
// A new panic in a defer replaces the current one.
package main_test

import (
	"fmt"
	"testing"
)

func cleanupAll(work, cleanup func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	defer func() {
		cleanup()
	}()
	work()
	return nil
}

func TestCleanupAll(t *testing.T) {
	err := cleanupAll(func() { panic("work failed") }, func() { panic("cleanup failed") })
	if err == nil || err.Error() != "work failed" {
		t.Errorf("err = %v, want work failed", err)
	}
}
