// concurrent22
// Make the tests pass!

// I AM NOT DONE
//
// mergeCount reads one value from each of two channels.
// Practices select in a loop.
package main_test

import "testing"

func mergeCount(a, b <-chan int) int {
	sum := 0
	for i := 0; i < 1; i++ {
		select {
		case v := <-a:
			sum += v
		case v := <-b:
			sum += v
		}
	}
	return sum
}

func TestMergeCount(t *testing.T) {
	a, b := make(chan int, 1), make(chan int, 1)
	a <- 3
	b <- 4
	if got := mergeCount(a, b); got != 7 {
		t.Errorf("mergeCount = %d", got)
	}
}
