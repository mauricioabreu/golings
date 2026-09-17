// generics17
// Make the tests pass!

// I AM NOT DONE
//
// ParallelMap applies a function in goroutines and keeps the order.
// Practices generic functions with goroutines.
package main_test

import (
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func ParallelMap[T, U any](s []T, f func(T) U) []U {
	out := make([]U, len(s))
	var wg sync.WaitGroup
	for i, v := range s {
		wg.Add(1)
		go func(i int, v T) {
			defer wg.Done()
			out[i] = f(v)
		}(i, v)
	}
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	return out
}

func TestParallelMap(t *testing.T) {
	got := ParallelMap([]int{1, 2, 3}, func(x int) string {
		return strconv.Itoa(x * x)
	})
	if !reflect.DeepEqual(got, []string{"1", "4", "9"}) {
		t.Errorf("ParallelMap = %v", got)
	}
}
