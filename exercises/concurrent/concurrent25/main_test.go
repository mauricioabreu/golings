// concurrent25
// Make the tests pass!

// I AM NOT DONE
//
// store sends a value into a channel and reads it back in the same goroutine.
// With no buffer the send blocks forever.
// Practices buffered channels.
package main_test

import (
	"testing"
	"time"
)

func store(v int) int {
	ch := make(chan int)
	ch <- v
	return <-ch
}

func TestStore(t *testing.T) {
	done := make(chan int, 1)
	go func() { done <- store(7) }()
	select {
	case got := <-done:
		if got != 7 {
			t.Errorf("store = %d", got)
		}
	case <-time.After(time.Second):
		t.Fatal("store is blocked")
	}
}
