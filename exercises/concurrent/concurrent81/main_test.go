// concurrent81
// Make the tests pass!

// I AM NOT DONE
//
// gather waits for task results up to a deadline and returns what arrived.
// Practices a single shared timer for the whole collection.
package main_test

import (
	"testing"
	"time"
)

func gather(tasks []time.Duration, deadline time.Duration) int {
	ch := make(chan int, len(tasks))
	for i, d := range tasks {
		go func(i int, d time.Duration) {
			time.Sleep(d)
			ch <- i
		}(i, d)
	}
	got := 0
	for got < len(tasks) {
		<-ch
		got++
	}
	return got
}

func TestGather(t *testing.T) {
	start := time.Now()
	got := gather([]time.Duration{time.Millisecond, 2 * time.Millisecond, 2 * time.Second}, 100*time.Millisecond)
	if got != 2 || time.Since(start) > time.Second {
		t.Errorf("gathered %d in %v", got, time.Since(start))
	}
}
