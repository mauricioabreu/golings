// slices78
// Make the tests pass!

// I AM NOT DONE
//
// sortNames sorts the names alphabetically.
// Practices sort.Strings.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func sortNames(names []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
}

func TestSortNames(t *testing.T) {
	names := []string{"bob", "ann", "cid"}
	sortNames(names)
	if !reflect.DeepEqual(names, []string{"ann", "bob", "cid"}) {
		t.Errorf("sortNames = %v", names)
	}
}
