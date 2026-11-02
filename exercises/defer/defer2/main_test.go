// defer2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"slices"
	"testing"
)

var events []string

func open(name string)    { events = append(events, "open "+name) }
func release(name string) { events = append(events, "close "+name) }

// work opens the file, then the connection. Whatever was opened last has to be
// closed first: the connection before the file.
func work() {
	open("file")
	open("connection")

	// TODO: Deferred calls run in the opposite order they were deferred. Arrange them so that the connection is closed first.
	defer release("connection")
	defer release("file")
}

func TestWorkClosesInReverseOrder(t *testing.T) {
	events = nil

	work()

	want := []string{"open file", "open connection", "close connection", "close file"}
	if !slices.Equal(events, want) {
		t.Errorf("got %v, want %v", events, want)
	}
}
