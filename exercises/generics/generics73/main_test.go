// generics73
// Make the tests pass!

// I AM NOT DONE
//
// Retry[T] calls a function until it succeeds, at most n times, and returns its result.
// Practices generic functions returning (T, error).
package main_test

import (
	"errors"
	"testing"
)

func Retry[T any](n int, f func() (T, error)) (T, error) {
	var zero T
	var err error
	for i := 0; i < n; i++ {
		var v T
		v, err = f()
		_ = v
	}
	return zero, nil
}

func TestRetry(t *testing.T) {
	calls := 0
	v, err := Retry(3, func() (string, error) {
		calls++
		if calls < 2 {
			return "", errors.New("busy")
		}
		return "ok", nil
	})
	if err != nil || v != "ok" || calls != 2 {
		t.Errorf("v=%q err=%v calls=%d", v, err, calls)
	}
	_, err = Retry(2, func() (int, error) { return 0, errors.New("down") })
	if err == nil {
		t.Errorf("Retry should return the last error")
	}
}
