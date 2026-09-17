// concurrent75
// Make the tests pass!

// I AM NOT DONE
//
// Goroutines send errors into a channel that is only read after wg.Wait.
// The unbuffered channel blocks the goroutines forever.
// Practices a buffer sized for the number of senders.
package main_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func validateAll(items []int) int {
	errs := make(chan error)
	var wg sync.WaitGroup
	for _, it := range items {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			if v < 0 {
				errs <- fmt.Errorf("negative %d", v)
			}
		}(it)
	}
	wg.Wait()
	close(errs)
	n := 0
	for range errs {
		n++
	}
	return n
}

func TestValidateAll(t *testing.T) {
	done := make(chan int, 1)
	go func() { done <- validateAll([]int{1, -2, -3}) }()
	select {
	case n := <-done:
		if n != 2 {
			t.Errorf("errors = %d", n)
		}
	case <-time.After(time.Second):
		t.Fatal("validateAll is blocked")
	}
}
