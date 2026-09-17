// concurrent63
// Make the tests pass!

// I AM NOT DONE
//
// merge combines several channels into one and closes it once all the
// input channels are closed.
// Practices fan-in with a WaitGroup.
package main_test

import (
	"sync"
	"testing"
	"time"
)

func merge(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, in := range ins {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(in)
	}
	close(out)
	return out
}

func source(vals ...int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, v := range vals {
			c <- v
		}
	}()
	return c
}

func TestMerge(t *testing.T) {
	done := make(chan int, 1)
	go func() {
		s := 0
		for v := range merge(source(1, 2), source(10), source()) {
			s += v
		}
		done <- s
	}()
	select {
	case got := <-done:
		if got != 13 {
			t.Errorf("sum = %d", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("merge is stuck")
	}
}
