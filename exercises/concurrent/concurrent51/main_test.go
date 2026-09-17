// concurrent51
// Make the tests pass!

// I AM NOT DONE
//
// numbers sends the numbers and has to close the channel, otherwise the range on the
// receiving side never ends.
// close as the signal that the data has run out.
package main_test

import (
	"testing"
	"time"
)

func numbers(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func TestNumbers(t *testing.T) {
	done := make(chan int, 1)
	go func() {
		s := 0
		for v := range numbers(4) {
			s += v
		}
		done <- s
	}()
	select {
	case got := <-done:
		if got != 10 {
			t.Errorf("sum = %d", got)
		}
	case <-time.After(time.Second):
		t.Fatal("range never finished")
	}
}
