// generics102
// Make the tests pass!

// I AM NOT DONE
//
// Reverse returns a new slice in reverse order.
// Practices generic functions over slices.
package main_test

import (
	"reflect"
	"testing"
)

func Reverse[T any](s []T) []T {
	out := make([]T, len(s))
	for i, v := range s {
		out[i] = v
	}
	return out
}

func TestReverse(t *testing.T) {
	if got := Reverse([]rune("abc")); string(got) != "cba" {
		t.Errorf("Reverse = %q", string(got))
	}
	if got := Reverse([]int{1, 2}); !reflect.DeepEqual(got, []int{2, 1}) {
		t.Errorf("Reverse = %v", got)
	}
}
