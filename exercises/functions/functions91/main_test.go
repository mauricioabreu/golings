// functions91
// Make the tests pass!

// I AM NOT DONE
//
// retry must call f until it succeeds, but no more than attempts times,
// and return the last error when every attempt fails.
// Practices passing functions and driving a loop by a result.
package main_test

import (
	"errors"
	"testing"
)

func retry(attempts int, f func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		f()
	}
	return err
}

func TestRetry(t *testing.T) {
	calls := 0
	err := retry(5, func() error {
		calls++
		if calls < 3 {
			return errors.New("temporary")
		}
		return nil
	})
	if err != nil || calls != 3 {
		t.Errorf("retry: err=%v calls=%d, want nil and 3", err, calls)
	}
	calls = 0
	err = retry(2, func() error { calls++; return errors.New("down") })
	if err == nil || calls != 2 {
		t.Errorf("retry: err=%v calls=%d, want error and 2", err, calls)
	}
}
