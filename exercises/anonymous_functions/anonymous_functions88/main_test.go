// anonymous_functions88
// Make the tests pass!

// I AM NOT DONE
//
// pager returns a literal handing out the next page of size size on every call;
// an empty slice means the pages have run out.
// Practices a closure over a position and slice bounds.
package main_test

import (
	"reflect"
	"testing"
)

func pager(items []int, size int) func() []int {
	pos := 0
	return func() []int {
		if pos >= len(items) {
			return nil
		}
		end := pos + size
		page := items[pos:end]
		pos++
		return page
	}
}

func TestPager(t *testing.T) {
	next := pager([]int{1, 2, 3, 4, 5}, 2)
	var pages [][]int
	for p := next(); p != nil; p = next() {
		pages = append(pages, p)
	}
	if !reflect.DeepEqual(pages, [][]int{{1, 2}, {3, 4}, {5}}) {
		t.Errorf("pages = %v", pages)
	}
}
