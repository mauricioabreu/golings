// range71
// Make the tests pass!

// I AM NOT DONE
//
// visit walks the elements through a callback and has to stop the walk at "end".
// The code does not compile: a break inside a function literal does not belong to the loop.
// break is only allowed inside a for, switch or select of the same function.
package main_test

import (
	"reflect"
	"testing"
)

func walk(items []string, f func(string) bool) {
	for _, it := range items {
		if !f(it) {
			return
		}
	}
}

func visit(items []string) []string {
	var seen []string
	walk(items, func(s string) bool {
		if s == "end" {
			break
		}
		seen = append(seen, s)
		return true
	})
	return seen
}

func TestVisit(t *testing.T) {
	if got := visit([]string{"a", "b", "end", "c"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("visit = %v", got)
	}
}
