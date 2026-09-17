// concurrent90
// Make the tests pass!

// I AM NOT DONE
//
// consumer receives a read-only channel and tries to close it.
// The code does not compile.
// Practices that only a channel you may send to can be closed.
package main_test

import "testing"

func consumer(in <-chan int) int {
	s := 0
	for v := range in {
		s += v
	}
	close(in)
	return s
}

func producer(out chan<- int) {
	for i := 1; i <= 3; i++ {
		out <- i
	}
}

func TestConsumer(t *testing.T) {
	ch := make(chan int, 3)
	producer(ch)
	if got := consumer(ch); got != 6 {
		t.Errorf("sum = %d", got)
	}
}
