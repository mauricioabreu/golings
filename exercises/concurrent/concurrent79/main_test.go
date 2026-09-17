// concurrent79
// Make the tests pass!

// I AM NOT DONE
//
// freeSlots returns how many more messages fit in the buffer.
// Practices cap and len for channels.
package main_test

import "testing"

func freeSlots(ch chan int) int {
	return len(ch) - cap(ch)
}

func TestFreeSlots(t *testing.T) {
	ch := make(chan int, 4)
	ch <- 1
	if freeSlots(ch) != 3 {
		t.Errorf("freeSlots = %d", freeSlots(ch))
	}
}
