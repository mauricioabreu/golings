// concurrent28
// Make the tests pass!

// I AM NOT DONE
//
// pending returns the number of unread messages in a buffered channel.
// Practices len on channels.
package main_test

import "testing"

func pending(ch chan string) int {
	return cap(ch)
}

func TestPending(t *testing.T) {
	ch := make(chan string, 5)
	ch <- "a"
	ch <- "b"
	if pending(ch) != 2 {
		t.Errorf("pending = %d", pending(ch))
	}
}
