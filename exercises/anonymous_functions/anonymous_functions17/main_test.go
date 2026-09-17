// anonymous_functions17
// Make the tests pass!

// I AM NOT DONE
//
// processAll has to give the records consecutive numbers, and they all
// come out as 1: the generator is built afresh for every record.
// The state of a closure lives as long as the literal itself.
package main_test

import (
	"reflect"
	"testing"
)

func counter() func() int {
	n := 0
	return func() int { n++; return n }
}

func processAll(names []string) []int {
	var ids []int
	for range names {
		next := counter()
		ids = append(ids, next())
	}
	return ids
}

func TestProcessAll(t *testing.T) {
	if got := processAll([]string{"a", "b", "c"}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("processAll = %v", got)
	}
}
