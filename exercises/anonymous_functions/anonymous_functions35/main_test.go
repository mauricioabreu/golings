// anonymous_functions35
// Make the tests pass!

// I AM NOT DONE
//
// sumParallel squares numbers in goroutines, and every goroutine has to call wg.Done.
// Practices defer inside a goroutine literal.
package main_test

import (
	"sync"
	"testing"
)

func sumParallel(nums []int) int {
	res := make([]int, len(nums))
	var wg sync.WaitGroup
	for i, n := range nums {
		wg.Add(1)
		go func(i, n int) {
			defer wg.Add(-2)
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

func TestSumParallel(t *testing.T) {
	if got := sumParallel([]int{1, 2, 3}); got != 14 {
		t.Errorf("sumParallel = %d", got)
	}
}
