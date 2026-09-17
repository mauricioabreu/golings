// primitive_types54
// Make the tests pass!

// I AM NOT DONE
//
// sortNames sorts names alphabetically, ignoring case.
// Right now every capital letter comes before the lowercase ones.
// Strings compare byte by byte: 'Z' (90) < 'a' (97).
package main_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func sortNames(names []string) {
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
}

func TestSortNames(t *testing.T) {
	_ = strings.ToLower
	names := []string{"bob", "Alice", "carol", "Dave"}
	sortNames(names)
	if want := []string{"Alice", "bob", "carol", "Dave"}; !reflect.DeepEqual(names, want) {
		t.Errorf("sortNames = %v, want %v", names, want)
	}
}
