// anonymous_functions13
// Make the tests pass!

// I AM NOT DONE
//
// parallelMap applies f to every element in a goroutine of its own and keeps
// the order of the results.
// Practices goroutine literals with the index passed as an argument.
package main_test

import (
	"reflect"
	"sync"
	"testing"
)

func parallelMap(in []int, f func(int) int) []int {
	out := make([]int, len(in))
	var wg sync.WaitGroup
	for i, v := range in {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out = append(out, f(v))
		}()
	}
	wg.Wait()
	return out
}

func TestParallelMap(t *testing.T) {
	got := parallelMap([]int{1, 2, 3, 4}, func(x int) int { return x * 10 })
	if !reflect.DeepEqual(got, []int{10, 20, 30, 40}) {
		t.Errorf("parallelMap = %v", got)
	}
}
