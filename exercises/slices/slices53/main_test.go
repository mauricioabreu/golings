// slices53
// Make the tests pass!

// I AM NOT DONE
//
// partition splits a slice into the elements that satisfy a predicate and those that do not.
// Practices two result slices.
package main_test

import (
	"reflect"
	"testing"
)

func partition(s []int, pred func(int) bool) (yes, no []int) {
	for _, v := range s {
		if pred(v) {
			yes = append(yes, v)
		}
		no = append(no, v)
	}
	return
}

func TestPartition(t *testing.T) {
	yes, no := partition([]int{1, 2, 3, 4, 5}, func(v int) bool { return v > 2 })
	if !reflect.DeepEqual(yes, []int{3, 4, 5}) || !reflect.DeepEqual(no, []int{1, 2}) {
		t.Errorf("partition = %v, %v", yes, no)
	}
}
