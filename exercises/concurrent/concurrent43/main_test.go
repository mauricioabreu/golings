// concurrent43
// Make the tests pass!

// I AM NOT DONE
//
// listen has to finish total after it started, even while messages
// keep arriving. A time.After in every iteration restarts the timer.
// A shared timeout is created once.
package main_test

import (
	"testing"
	"time"
)

func listen(msgs <-chan int, total time.Duration) int {
	n := 0
	for {
		select {
		case <-msgs:
			n++
		case <-time.After(total):
			return n
		}
	}
}

func TestListen(t *testing.T) {
	msgs := make(chan int)
	stop := make(chan struct{})
	defer close(stop)
	go func() {
		for {
			select {
			case msgs <- 1:
				time.Sleep(5 * time.Millisecond)
			case <-stop:
				return
			}
		}
	}()
	res := make(chan int, 1)
	go func() { res <- listen(msgs, 100*time.Millisecond) }()
	select {
	case n := <-res:
		if n == 0 {
			t.Errorf("no messages received")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("listen never timed out")
	}
}
