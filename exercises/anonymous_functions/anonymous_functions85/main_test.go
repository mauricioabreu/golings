// anonymous_functions85
// Make the tests pass!

// I AM NOT DONE
//
// filter keeps the elements the callback returned true for;
// bigOnes passes a "larger than limit" literal.
// Practices a literal using the parameter of the outer function.
package main_test

import (
	"reflect"
	"testing"
)

func filter(s []int, keep func(int) bool) []int {
	var out []int
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func bigOnes(s []int, limit int) []int {
	return filter(s, func(v int) bool { return v < limit })
}

func TestBigOnes(t *testing.T) {
	if got := bigOnes([]int{1, 7, 3, 9}, 5); !reflect.DeepEqual(got, []int{7, 9}) {
		t.Errorf("bigOnes = %v", got)
	}
}
