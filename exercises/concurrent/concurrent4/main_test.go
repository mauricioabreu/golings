// concurrent4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"slices"
	"testing"
	"time"
)

// roundTrip sends the numbers into a channel and then reads them back, all in
// the same goroutine.
func roundTrip(numbers []int) []int {
	// TODO: A send on an unbuffered channel waits for a receiver, and nobody is receiving yet. Give the channel room for the numbers.
	ch := make(chan int)

	for _, n := range numbers {
		ch <- n
	}
	close(ch)

	var out []int
	for n := range ch {
		out = append(out, n)
	}
	return out
}

func TestRoundTrip(t *testing.T) {
	numbers := []int{1, 2, 3}
	result := make(chan []int, 1)

	go func() { result <- roundTrip(numbers) }()

	select {
	case got := <-result:
		if !slices.Equal(got, numbers) {
			t.Errorf("roundTrip(%v) = %v", numbers, got)
		}
	case <-time.After(time.Second):
		t.Fatal("roundTrip is stuck: a send waits for a receiver, and nobody is receiving yet (a deadlock)")
	}
}
