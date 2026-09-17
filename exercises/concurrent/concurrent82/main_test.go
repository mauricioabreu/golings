// concurrent82
// Make the tests pass!

// I AM NOT DONE
//
// compute starts a computation in a goroutine but never reads the result
// from the channel.
// Practices receiving a value from a channel.
package main_test

import "testing"

func compute(a, b int) int {
	ch := make(chan int, 1)
	go func() { ch <- a * b }()
	return len(ch)
}

func TestCompute(t *testing.T) {
	if got := compute(6, 7); got != 42 {
		t.Errorf("compute = %d", got)
	}
}
