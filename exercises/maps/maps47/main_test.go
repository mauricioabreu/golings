// maps47
// Make the tests pass!

// I AM NOT DONE
//
// The task groups live in a map[string][]string. After the last task is removed
// the group is left empty and the number of groups is wrong.
// An empty slice in a map is still a key that exists.
package main_test

import "testing"

type Board map[string][]string

func (b Board) Done(group, task string) {
	tasks := b[group]
	for i, t := range tasks {
		if t == task {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	b[group] = tasks
}

func TestBoard(t *testing.T) {
	b := Board{"todo": {"a"}, "doing": {"b", "c"}}
	b.Done("todo", "a")
	b.Done("doing", "b")
	if len(b) != 1 || len(b["doing"]) != 1 {
		t.Errorf("board = %v", b)
	}
}
