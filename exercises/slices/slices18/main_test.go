// slices18
// Make the tests pass!

// I AM NOT DONE
//
// splitAlternate sorts the elements at even and odd indexes into two slices.
// Practices the index of a range and two results.
package main_test

import (
	"reflect"
	"testing"
)

func splitAlternate(s []string) (even, odd []string) {
	for _, v := range s {
		if len(v)%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return
}

func TestSplitAlternate(t *testing.T) {
	e, o := splitAlternate([]string{"a", "bb", "c", "dd", "e"})
	if !reflect.DeepEqual(e, []string{"a", "c", "e"}) || !reflect.DeepEqual(o, []string{"bb", "dd"}) {
		t.Errorf("splitAlternate = %v, %v", e, o)
	}
}
