// concurrent94
// Make the tests pass!

// I AM NOT DONE
//
// waitResult waits for a result for at most 50 ms. The timeout is set much
// too large.
// Practices time.After in a select.
package main_test

import (
	"errors"
	"testing"
	"time"
)

func waitResult(ch <-chan string) (string, error) {
	select {
	case v := <-ch:
		return v, nil
	case <-time.After(50 * time.Hour):
		return "", errors.New("timeout")
	}
}

func TestWaitResult(t *testing.T) {
	done := make(chan error, 1)
	go func() {
		_, err := waitResult(make(chan string))
		done <- err
	}()
	select {
	case err := <-done:
		if err == nil {
			t.Errorf("expected timeout error")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timeout is too long")
	}
}
