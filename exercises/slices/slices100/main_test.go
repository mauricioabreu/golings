// slices100
// Make the tests pass!

// I AM NOT DONE
//
// reverseInPlace has to reverse a slice, and the slice comes back unchanged.
// Every pair is swapped twice when the loop runs over the whole length.
package main_test

import (
	"reflect"
	"testing"
)

func reverseInPlace(s []int) {
	n := len(s)
	for i := 0; i < n; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

func TestReverseInPlace(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	reverseInPlace(s)
	if !reflect.DeepEqual(s, []int{5, 4, 3, 2, 1}) {
		t.Errorf("reverseInPlace = %v", s)
	}
}
