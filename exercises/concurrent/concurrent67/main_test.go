// concurrent67
// Make the tests pass!

// I AM NOT DONE
//
// runTasks collects task errors into a buffered channel and counts them.
// The error count is incremented incorrectly.
// Practices reading all values out of a channel.
package main_test

import (
	"errors"
	"sync"
	"testing"
)

func runTasks(tasks []func() error) int {
	errs := make(chan error, len(tasks))
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(f func() error) {
			defer wg.Done()
			if err := f(); err != nil {
				errs <- err
			}
		}(task)
	}
	wg.Wait()
	close(errs)
	n := 0
	for range errs {
		n = 1
	}
	return n
}

func TestRunTasks(t *testing.T) {
	fail := func() error { return errors.New("x") }
	ok := func() error { return nil }
	if got := runTasks([]func() error{fail, ok, fail, fail}); got != 3 {
		t.Errorf("errors = %d", got)
	}
}
