// concurrent38
// Make the tests pass!

// I AM NOT DONE
//
// sumUntilClosed reads from a channel in a select loop and adds the values up.
// Once the channel is closed a receive returns zero at once, and the loop never ends.
// Practices the ok flag when receiving in a select.
package main_test

import (
	"testing"
	"time"
)

func sumUntilClosed(ch <-chan int, abort <-chan struct{}) (int, bool) {
	sum := 0
	for {
		select {
		case v := <-ch:
			sum += v
		case <-abort:
			return sum, false
		}
	}
}

func TestSumUntilClosed(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)
	abort := make(chan struct{})
	go func() {
		time.Sleep(200 * time.Millisecond)
		close(abort)
	}()
	sum, completed := sumUntilClosed(ch, abort)
	if sum != 3 || !completed {
		t.Errorf("sum = %d, completed = %v", sum, completed)
	}
}
