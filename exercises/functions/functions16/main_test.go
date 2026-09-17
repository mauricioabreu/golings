// functions16
// Make the tests pass!

// I AM NOT DONE
//
// normalized must return the values divided by the maximum WITHOUT changing the caller's slice.
// Right now the original slice changes too.
// In a call f(s...) the parameter shares the array with s.
package main_test

import (
	"reflect"
	"testing"
)

func normalized(vals ...int) []int {
	max := 0
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	out := vals
	for i := range out {
		out[i] = out[i] * 100 / max
	}
	return out
}

func TestNormalized(t *testing.T) {
	data := []int{5, 10, 20}
	got := normalized(data...)
	if !reflect.DeepEqual(got, []int{25, 50, 100}) {
		t.Errorf("normalized = %v, want [25 50 100]", got)
	}
	if !reflect.DeepEqual(data, []int{5, 10, 20}) {
		t.Errorf("input changed to %v", data)
	}
}
