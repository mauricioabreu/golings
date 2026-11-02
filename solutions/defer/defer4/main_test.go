// defer4
// Make the tests pass!

package main_test

import (
	"slices"
	"testing"
)

var events []string

func open(name string)    { events = append(events, "open "+name) }
func release(name string) { events = append(events, "close "+name) }
func use(name string)     { events = append(events, "use "+name) }

// processAll opens, uses and closes the resources one after the other: each
// one must be closed before the next one is opened.
func processAll(names []string) {
	for _, name := range names {
		processOne(name)
	}
}

// processOne is a function of its own so that its deferred call runs when it
// returns, which is the end of the iteration.
func processOne(name string) {
	open(name)
	defer release(name)
	use(name)
}

func TestProcessAll(t *testing.T) {
	events = nil

	processAll([]string{"a", "b"})

	want := []string{"open a", "use a", "close a", "open b", "use b", "close b"}
	if !slices.Equal(events, want) {
		t.Errorf("got %v, want %v", events, want)
	}
}
