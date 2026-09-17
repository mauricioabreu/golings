// anonymous_functions98
// Make the tests pass!

// I AM NOT DONE
//
// A background job has to send its result into a channel. The code does not compile.
// The expression in a go statement has to be a function call.
package main_test

import "testing"

func background() int {
	ch := make(chan int, 1)
	go func() {
		ch <- 42
	}
	return <-ch
}

func TestBackground(t *testing.T) {
	if background() != 42 {
		t.Errorf("background = %d", background())
	}
}
