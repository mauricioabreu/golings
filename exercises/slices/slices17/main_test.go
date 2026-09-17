// slices17
// Make the tests pass!

// I AM NOT DONE
//
// lengths returns the lengths of the strings.
// Practices building a result slice of the right length.
package main_test

import (
	"reflect"
	"testing"
)

func lengths(words []string) []int {
	out := make([]int, len(words))
	for i, w := range words {
		out[i] = i
	}
	return out
}

func TestLengths(t *testing.T) {
	if got := lengths([]string{"go", "gopher", ""}); !reflect.DeepEqual(got, []int{2, 6, 0}) {
		t.Errorf("lengths = %v", got)
	}
}
