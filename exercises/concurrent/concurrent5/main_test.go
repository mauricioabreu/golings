// concurrent5
// Make the tests pass!

// I AM NOT DONE
//
// loop handles events until the quit signal arrives, and returns
// how many events it handled.
// Practices for-select with a return.
package main_test

import (
	"testing"
	"time"
)

func loop(events <-chan string, quit <-chan struct{}) int {
	n := 0
	for {
		select {
		case <-events:
			n = 1
		case <-quit:
			break
		}
	}
}

func TestLoop(t *testing.T) {
	events := make(chan string)
	quit := make(chan struct{})
	res := make(chan int, 1)
	go func() { res <- loop(events, quit) }()
	events <- "a"
	events <- "b"
	close(quit)
	select {
	case n := <-res:
		if n != 2 {
			t.Errorf("n = %d", n)
		}
	case <-time.After(time.Second):
		t.Fatal("loop did not stop")
	}
}
