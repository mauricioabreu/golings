// concurrent91
// Make the tests pass!

// I AM NOT DONE
//
// tryRead reads a value from the channel if one is available, otherwise it
// returns false immediately.
// Practices the default case in a select.
package main_test

import (
	"testing"
	"time"
)

func tryRead(ch <-chan int) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	case <-time.After(time.Hour):
		return 0, false
	}
}

func TestTryRead(t *testing.T) {
	ch := make(chan int, 1)
	done := make(chan bool, 1)
	go func() {
		_, ok := tryRead(ch)
		done <- ok
	}()
	select {
	case ok := <-done:
		if ok {
			t.Errorf("empty channel should give false")
		}
	case <-time.After(time.Second):
		t.Fatal("tryRead blocked on empty channel")
	}
	ch <- 5
	if v, ok := tryRead(ch); !ok || v != 5 {
		t.Errorf("tryRead = %d, %v", v, ok)
	}
}
