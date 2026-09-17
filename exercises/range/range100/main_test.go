// range100
// Make the tests pass!

// I AM NOT DONE
//
// sumChannel adds up the numbers from a channel until it is closed.
// Practices a range over a channel.
package main_test

import "testing"

func sumChannel(ch <-chan int) int {
	total := 0
	for v := range ch {
		total = v
	}
	return total
}

func TestSumChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 4; i++ {
			ch <- i
		}
		close(ch)
	}()
	if got := sumChannel(ch); got != 10 {
		t.Errorf("sumChannel = %d, want 10", got)
	}
}
