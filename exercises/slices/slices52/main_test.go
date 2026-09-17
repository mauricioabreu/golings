// slices52
// Make the tests pass!

// I AM NOT DONE
//
// expand walks a queue of tasks and appends subtasks to its end;
// the subtasks have to be processed too. Right now they are skipped.
// A range evaluates the length of the slice once, before the loop starts.
package main_test

import (
	"reflect"
	"testing"
)

var children = map[string][]string{
	"root": {"a", "b"},
	"a":    {"a1"},
}

func expand(start string) []string {
	queue := []string{start}
	var visited []string
	for _, n := range queue {
		visited = append(visited, n)
		queue = append(queue, children[n]...)
	}
	return visited
}

func TestExpand(t *testing.T) {
	if got := expand("root"); !reflect.DeepEqual(got, []string{"root", "a", "b", "a1"}) {
		t.Errorf("expand = %v", got)
	}
}
