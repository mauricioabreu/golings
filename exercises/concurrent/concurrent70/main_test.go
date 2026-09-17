// concurrent70
// Make the tests pass!

// I AM NOT DONE
//
// sumSquares launches goroutines and waits for them with a WaitGroup,
// but Wait does not wait for anything.
// Practices calling wg.Add before launching a goroutine.
package main_test

import (
	"sync"
	"testing"
)

func sumSquares(nums []int) int {
	res := make([]int, len(nums))
	var wg sync.WaitGroup
	for i, n := range nums {
		wg.Add(0)
		go func(i, n int) {
			defer wg.Done()
			res[i] = n * n
		}(i, n)
	}
	wg.Wait()
	total := 0
	for _, v := range res {
		total += v
	}
	return total
}

func TestSumSquares(t *testing.T) {
	if got := sumSquares([]int{1, 2, 3}); got != 14 {
		t.Errorf("sumSquares = %d, want 14", got)
	}
}
