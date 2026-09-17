// concurrent54
// Make the tests pass!

// I AM NOT DONE
//
// Two goroutines take turns passing a ball through channels,
// incrementing a counter.
// Practices alternating through unbuffered channels.
package main_test

import (
	"testing"
	"time"
)

func player(in <-chan int, out chan<- int, limit int, done chan<- int) {
	for n := range in {
		if n > limit {
			done <- n
			return
		}
		out <- n
	}
}

func rally(limit int) int {
	a, b := make(chan int), make(chan int)
	done := make(chan int, 2)
	go player(a, b, limit, done)
	go player(b, a, limit, done)
	a <- 0
	return <-done
}

func TestRally(t *testing.T) {
	res := make(chan int, 1)
	go func() { res <- rally(10) }()
	select {
	case got := <-res:
		if got != 10 {
			t.Errorf("rally = %d", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("rally never ended")
	}
}
