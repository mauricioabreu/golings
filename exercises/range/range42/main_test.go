// range42
// Make the tests pass!

// I AM NOT DONE
//
// produce sends numbers into a channel, and total reads them with a range.
// total never finishes.
// A range over a channel waits for the channel to be closed.
package main_test

import (
	"testing"
	"time"
)

func produce(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

func total(ch <-chan int) int {
	s := 0
	for v := range ch {
		s += v
	}
	return s
}

func TestTotal(t *testing.T) {
	done := make(chan int, 1)
	go func() { done <- total(produce([]int{1, 2, 3})) }()
	select {
	case got := <-done:
		if got != 6 {
			t.Errorf("total = %d, want 6", got)
		}
	case <-time.After(time.Second):
		t.Fatal("total is stuck: the channel is never closed")
	}
}
