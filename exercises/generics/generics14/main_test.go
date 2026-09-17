// generics14
// Make the tests pass!

// I AM NOT DONE
//
// CountEqual counts the elements equal to x. Instantiated with any,
// comparing slices inside interfaces panics at run time.
// comparable allows interface types, and == may still panic.
package main_test

import (
	"reflect"
	"testing"
)

func CountEqual[T comparable](s []T, x T) int {
	n := 0
	for _, v := range s {
		if v == x {
			n++
		}
	}
	return n
}

func TestCountEqual(t *testing.T) {
	_ = reflect.DeepEqual
	vals := []any{1, []int{1}, "a", []int{1}}
	if got := CountEqual(vals, any([]int{1})); got != 2 {
		t.Errorf("CountEqual = %d, want 2", got)
	}
}
