// functions44
// Make the tests pass!

// I AM NOT DONE
//
// setTags with no arguments leaves the tags alone, while an explicitly passed
// empty slice (setTags(empty...)) clears them.
// Called with no arguments, a variadic parameter is nil.
package main_test

import (
	"reflect"
	"testing"
)

type Item struct{ Tags []string }

func (it *Item) setTags(tags ...string) {
	if len(tags) == 0 {
		return
	}
	it.Tags = tags
}

func TestSetTags(t *testing.T) {
	it := &Item{Tags: []string{"go"}}
	it.setTags()
	if !reflect.DeepEqual(it.Tags, []string{"go"}) {
		t.Errorf("setTags() changed tags to %v", it.Tags)
	}
	it.setTags([]string{}...)
	if len(it.Tags) != 0 {
		t.Errorf("setTags(empty...) should clear tags, got %v", it.Tags)
	}
	it.setTags("a", "b")
	if !reflect.DeepEqual(it.Tags, []string{"a", "b"}) {
		t.Errorf("setTags(a, b) = %v", it.Tags)
	}
}
