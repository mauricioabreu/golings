// concurrent62
// Make the tests pass!

// I AM NOT DONE
//
// Each goroutine must record its own number. The number is passed as
// a parameter.
// Practices passing a value via go func(v int){...}(v).
package main_test

import (
	"sort"
	"sync"
	"testing"
)

func ids(n int) []int {
	var mu sync.Mutex
	var out []int
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			out = append(out, id)
			mu.Unlock()
		}(n)
	}
	wg.Wait()
	sort.Ints(out)
	return out
}

func TestIDs(t *testing.T) {
	got := ids(3)
	if len(got) != 3 || got[0] != 0 || got[2] != 2 {
		t.Errorf("ids = %v", got)
	}
}
