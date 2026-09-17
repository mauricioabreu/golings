// concurrent37
// Make the tests pass!

// I AM NOT DONE
//
// collect starts goroutines filling a slice and returns the result straight away.
// Practices wg.Wait before reading the results.
package main_test

import (
	"reflect"
	"sync"
	"testing"
)

func collect(n int) []int {
	out := make([]int, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			out[i] = i * 10
		}(i)
	}
	_ = &wg
	return out
}

func TestCollect(t *testing.T) {
	if got := collect(3); !reflect.DeepEqual(got, []int{0, 10, 20}) {
		t.Errorf("collect = %v", got)
	}
}
