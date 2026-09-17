// concurrent4
// Make the tests pass!

// I AM NOT DONE
//
// worker takes the WaitGroup by value, so Done lowers the counter of a copy.
// A WaitGroup must not be copied.
package main_test

import (
	"sync"
	"testing"
	"time"
)

func worker(id int, out []int, wg sync.WaitGroup) {
	defer wg.Done()
	out[id] = id + 1
}

func runAll(n int) []int {
	out := make([]int, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, out, &wg)
	}
	wg.Wait()
	return out
}

func TestRunAll(t *testing.T) {
	done := make(chan []int, 1)
	go func() { done <- runAll(3) }()
	select {
	case got := <-done:
		if got[2] != 3 {
			t.Errorf("runAll = %v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("runAll is stuck")
	}
}
