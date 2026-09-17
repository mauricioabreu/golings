// range83
// Make the tests pass!

// I AM NOT DONE
//
// toSlice turns a map from page number to title into a slice of titles
// ordered by page number. The numbers have gaps, and the code panics.
// Map keys are not positions; the order has to come from a sort.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func toSlice(pages map[int]string) []string {
	out := make([]string, len(pages))
	for n, title := range pages {
		out[n] = title
	}
	return out
}

func TestToSlice(t *testing.T) {
	_ = sort.Ints
	got := toSlice(map[int]string{10: "c", 1: "a", 5: "b"})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("toSlice = %v", got)
	}
}
