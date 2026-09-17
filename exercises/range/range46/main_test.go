// range46
// Make the tests pass!

// I AM NOT DONE
//
// squares starts a goroutine sending the squares of 1..n and closing the channel;
// sumSquares reads them with a range.
// Practices a range over a channel the sender closes.
package main_test

import (
	"testing"
	"time"
)

func squares(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i * i
		}
	}()
	return ch
}

func sumSquares(n int) int {
	total := 0
	for v := range squares(n) {
		total += v
	}
	return total
}

func TestSumSquares(t *testing.T) {
	done := make(chan int, 1)
	go func() { done <- sumSquares(3) }()
	select {
	case got := <-done:
		if got != 14 {
			t.Errorf("sumSquares(3) = %d, want 14", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("sumSquares did not finish: is the channel closed?")
	}
}
