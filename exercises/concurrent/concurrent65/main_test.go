// concurrent65
// Make the tests pass!

// I AM NOT DONE
//
// Pipeline: gen emits numbers, square squares them, sum adds them up.
// Practices pipeline stages, each closing its own output channel.
package main_test

import (
	"testing"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + n
		}
	}()
	return out
}

func sum(in <-chan int) int {
	s := 0
	for v := range in {
		s += v
	}
	return s
}

func TestPipeline(t *testing.T) {
	done := make(chan int, 1)
	go func() { done <- sum(square(gen(1, 2, 3))) }()
	select {
	case got := <-done:
		if got != 14 {
			t.Errorf("sum = %d", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("pipeline is stuck")
	}
}
