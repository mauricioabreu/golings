// defer3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"fmt"
	"slices"
	"testing"
)

// trace returns what it recorded. lines is a named result, so the calls
// deferred in trace can still add to what is returned.
func trace() (lines []string) {
	record := func(line string) {
		lines = append(lines, line)
	}

	x := 1
	// TODO: The arguments of a deferred call are evaluated right away, when the defer statement runs. Make the deferred call report the value x has when trace returns.
	defer record(fmt.Sprintf("deferred: x = %d", x))

	x = 2
	record(fmt.Sprintf("x = %d", x))

	return lines
}

func TestTrace(t *testing.T) {
	want := []string{"x = 2", "deferred: x = 2"}

	if got := trace(); !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
