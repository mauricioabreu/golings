// concurrent8
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"slices"
	"sync"
	"testing"
	"time"
)

// process runs work on every job using the given number of workers, and
// returns the results in any order.
func process(jobs []int, workers int, work func(int) int) []int {
	// TODO: Start the workers: goroutines that take jobs from a channel until it is closed and send the results back.
	var results []int
	for _, job := range jobs {
		results = append(results, work(job))
	}
	return results
}

func TestProcessUsesAllTheWorkers(t *testing.T) {
	var mu sync.Mutex
	active, mostActive := 0, 0

	work := func(job int) int {
		mu.Lock()
		active++
		mostActive = max(mostActive, active)
		mu.Unlock()

		time.Sleep(100 * time.Millisecond)

		mu.Lock()
		active--
		mu.Unlock()
		return job * job
	}

	got := process([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, work)

	slices.Sort(got)
	if want := []int{1, 4, 9, 16, 25, 36, 49, 64}; !slices.Equal(got, want) {
		t.Errorf("results = %v, want %v", got, want)
	}
	if mostActive != 4 {
		t.Errorf("%d jobs ran at the same time, want 4 workers busy together", mostActive)
	}
}
