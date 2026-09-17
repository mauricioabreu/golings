// anonymous_functions45
// Make the tests pass!

// I AM NOT DONE
//
// collector returns an add literal appending to the slice it was given.
// The caller never sees the elements that were added.
// A literal captures the parameter, a copy of the slice header.
package main_test

import (
	"reflect"
	"testing"
)

func collector(dst []string) func(string) {
	return func(s string) { dst = append(dst, s) }
}

func gather() []string {
	var names []string
	add := collector(names)
	add("a")
	add("b")
	return names
}

func TestGather(t *testing.T) {
	if got := gather(); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("gather = %v", got)
	}
}
