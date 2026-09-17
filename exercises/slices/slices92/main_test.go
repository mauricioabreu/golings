// slices92
// Make the tests pass!

// I AM NOT DONE
//
// activeOnly filters a list, and the caller expects the original slice
// to stay as it was. Filtering through s[:0] spoils it.
// out := s[:0] reuses the array of the original slice.
package main_test

import (
	"reflect"
	"testing"
)

type Task struct {
	Name   string
	Active bool
}

func activeOnly(tasks []Task) []Task {
	out := tasks[:0]
	for _, t := range tasks {
		if t.Active {
			out = append(out, t)
		}
	}
	return out
}

func TestActiveOnly(t *testing.T) {
	tasks := []Task{{"a", false}, {"b", true}, {"c", true}}
	got := activeOnly(tasks)
	if len(got) != 2 || got[0].Name != "b" {
		t.Errorf("activeOnly = %v", got)
	}
	if !reflect.DeepEqual(tasks, []Task{{"a", false}, {"b", true}, {"c", true}}) {
		t.Errorf("tasks modified: %v", tasks)
	}
}
