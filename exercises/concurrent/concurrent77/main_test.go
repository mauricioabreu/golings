// concurrent77
// Make the tests pass!

// I AM NOT DONE
//
// tee copies every value from the input into two outputs.
// Practices sending one value to several channels.
package main_test

import (
	"sync"
	"testing"
	"time"
)

func tee(in <-chan int) (<-chan int, <-chan int) {
	a, b := make(chan int, 10), make(chan int, 10)
	go func() {
		defer close(a)
		for v := range in {
			a <- v
		}
	}()
	return a, b
}

func TestTee(t *testing.T) {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	a, b := tee(in)
	var sa, sb int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range a {
			sa += v
		}
	}()
	go func() {
		defer wg.Done()
		for v := range b {
			sb += v
		}
	}()
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
		if sa != 6 || sb != 6 {
			t.Errorf("sums = %d, %d", sa, sb)
		}
	case <-time.After(time.Second):
		t.Fatal("tee is stuck")
	}
}
