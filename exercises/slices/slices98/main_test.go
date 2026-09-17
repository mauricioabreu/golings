// slices98
// Make the tests pass!

// I AM NOT DONE
//
// without returns the list without element i, and the original list has to be left untouched.
// append(s[:i], s[i+1:]...) changes the array of the original slice.
package main_test

import (
	"reflect"
	"testing"
)

func without(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func TestWithout(t *testing.T) {
	team := []string{"ann", "bob", "cid"}
	got := without(team, 0)
	if !reflect.DeepEqual(got, []string{"bob", "cid"}) {
		t.Errorf("without = %v", got)
	}
	if !reflect.DeepEqual(team, []string{"ann", "bob", "cid"}) {
		t.Errorf("team modified: %v", team)
	}
}
