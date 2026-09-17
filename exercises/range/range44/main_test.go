// range44
// Make the tests pass!

// I AM NOT DONE
//
// evenPositions returns the indexes 0, 2, 4 and so on of a slice.
// Practices a range with a single variable, the index.
package main_test

import (
	"reflect"
	"testing"
)

func evenPositions(s []string) []int {
	var out []int
	for i := range s {
		if i%2 == 1 {
			out = append(out, i)
		}
	}
	return out
}

func TestEvenPositions(t *testing.T) {
	if got := evenPositions([]string{"a", "b", "c", "d", "e"}); !reflect.DeepEqual(got, []int{0, 2, 4}) {
		t.Errorf("evenPositions = %v", got)
	}
}
