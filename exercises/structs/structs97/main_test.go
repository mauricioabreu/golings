// structs97
// Make the tests pass!

// I AM NOT DONE
//
// The Label method with a value receiver is called on a nil pointer and panics,
// unlike methods with a pointer receiver.
// Calling a value method through a pointer dereferences it.
package main_test

import "testing"

type Tag struct{ Name string }

func (t Tag) Label() string {
	return "#" + t.Name
}

func TestLabel(t *testing.T) {
	var none *Tag
	if none.Label() != "#untagged" || (&Tag{"go"}).Label() != "#go" {
		t.Errorf("Label works incorrectly")
	}
}
