// slices15
// Make the tests pass!

// I AM NOT DONE
//
// addTag adds a tag to the list of tags of an article held in a map.
// The tag that was added does not stick.
// append returns a new slice header, and it has to be written back into the map.
package main_test

import (
	"reflect"
	"testing"
)

func addTag(m map[string][]string, post, tag string) {
	tags := m[post]
	tags = append(tags, tag)
}

func TestAddTag(t *testing.T) {
	m := map[string][]string{}
	addTag(m, "p1", "go")
	addTag(m, "p1", "slices")
	if !reflect.DeepEqual(m["p1"], []string{"go", "slices"}) {
		t.Errorf("m[p1] = %v", m["p1"])
	}
}
