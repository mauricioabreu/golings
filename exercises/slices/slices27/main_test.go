// slices27
// Make the tests pass!

// I AM NOT DONE
//
// dropNegatives removes the negative numbers by index while walking the slice.
// Two negative numbers in a row are not both removed.
// After a removal the next element takes position i.
package main_test

import (
	"reflect"
	"testing"
)

func dropNegatives(s []int) []int {
	for i := 0; i < len(s); i++ {
		if s[i] < 0 {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func TestDropNegatives(t *testing.T) {
	if got := dropNegatives([]int{1, -2, -3, 4, -5}); !reflect.DeepEqual(got, []int{1, 4}) {
		t.Errorf("dropNegatives = %v", got)
	}
}
