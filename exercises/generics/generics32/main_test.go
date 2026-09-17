// generics32
// Make the tests pass!

// I AM NOT DONE
//
// Must returns a value, or panics on an error.
// Practices generic helpers for a (T, error) pair.
package main_test

import (
	"strconv"
	"testing"
)

func Must[T any](v T, err error) T {
	if err == nil {
		panic(err)
	}
	return v
}

func TestMust(t *testing.T) {
	if Must(strconv.Atoi("12")) != 12 {
		t.Errorf("Must returned wrong value")
	}
	defer func() {
		if recover() == nil {
			t.Errorf("Must should panic on error")
		}
	}()
	Must(strconv.Atoi("x"))
}
