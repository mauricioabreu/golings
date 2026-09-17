// concurrent88
// Make the tests pass!

// I AM NOT DONE
//
// gather appends results to a shared slice from many goroutines;
// some results get lost.
// Practices that append is not thread-safe.
package main_test

import (
	"sync"
	"testing"
)

func gather(n int) []int {
	var out []int
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			out = append(out, i)
		}(i)
	}
	wg.Wait()
	_ = &mu
	return out
}

func TestGather(t *testing.T) {
	if got := gather(300); len(got) != 300 {
		t.Errorf("len = %d, want 300", len(got))
	}
}
