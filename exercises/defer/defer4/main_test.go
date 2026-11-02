// defer4
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
func use(name string)     { events = append(events, "use "+name) }

// processAll opens, uses and closes the resources one after the other: each
// one must be closed before the next one is opened.
func processAll(names []string) {
	for _, name := range names {
		open(name)
		// TODO: A deferred call runs when the function returns, not at the end of the iteration. Close each resource when its own turn is over.
		defer release(name)
		use(name)
	}
}

func TestProcessAll(t *testing.T) {
	events = nil

	processAll([]string{"a", "b"})

	want := []string{"open a", "use a", "close a", "open b", "use b", "close b"}
	if !slices.Equal(events, want) {
		t.Errorf("got %v, want %v", events, want)
	}
}
