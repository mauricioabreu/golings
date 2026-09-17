// functions15
// Make the tests pass!

// I AM NOT DONE
//
// flatten must turn nested []interface{} values into a flat []int recursively.
// Practices recursion and a type switch on the empty interface.
package main_test

import (
	"reflect"
	"testing"
)

func flatten(items []interface{}) []int {
	var out []int
	for _, it := range items {
		switch v := it.(type) {
		case int:
			out = append(out, v)
		default:
			continue
		}
	}
	return out
}

func TestFlatten(t *testing.T) {
	in := []interface{}{1, []interface{}{2, []interface{}{3, 4}}, 5}
	if got := flatten(in); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
		t.Errorf("flatten = %v, want [1 2 3 4 5]", got)
	}
}
