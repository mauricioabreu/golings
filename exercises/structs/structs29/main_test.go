// structs29
// Make the tests pass!

// I AM NOT DONE
//
// AddTag adds a tag to a task.
// Practices working with a slice field through a pointer.
package main_test

import (
	"reflect"
	"testing"
)

type Task struct{ Tags []string }

func (t *Task) AddTag(tag string) {
	_ = append(t.Tags, tag)
}

func TestAddTag(t *testing.T) {
	var task Task
	task.AddTag("bug")
	task.AddTag("ui")
	if !reflect.DeepEqual(task.Tags, []string{"bug", "ui"}) {
		t.Errorf("Tags = %v", task.Tags)
	}
}
