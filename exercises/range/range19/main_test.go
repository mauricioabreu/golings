// range19
// Make the tests pass!

// I AM NOT DONE
//
// smooth replaces every element but the first with the average of it and the previous
// element of the ORIGINAL slice. Changing in place uses values that have already changed.
// A range over a slice sees the changes made to later and earlier elements.
package main_test

import (
	"reflect"
	"testing"
)

func smooth(s []int) {
	for i := range s {
		if i > 0 {
			s[i] = (s[i] + s[i-1]) / 2
		}
	}
}

func TestSmooth(t *testing.T) {
	s := []int{0, 10, 20, 30}
	smooth(s)
	if !reflect.DeepEqual(s, []int{0, 5, 15, 25}) {
		t.Errorf("smooth = %v, want [0 5 15 25]", s)
	}
}
