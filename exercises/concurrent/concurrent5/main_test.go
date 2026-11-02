// concurrent5
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"slices"
	"sync"
	"testing"
)

// squares squares every number, each one in its own goroutine.
func squares(numbers []int) []int {
	results := make([]int, len(numbers))
	var wg sync.WaitGroup

	for i, n := range numbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = n * n
		}()
	}

	// TODO: The goroutines may not have finished yet. Wait for all of them before returning the results.
	return results
}

func TestSquares(t *testing.T) {
	want := []int{1, 4, 9, 16, 25}

	if got := squares([]int{1, 2, 3, 4, 5}); !slices.Equal(got, want) {
		t.Errorf("squares returned %v, want %v: did it wait for every goroutine to finish?", got, want)
	}
}
