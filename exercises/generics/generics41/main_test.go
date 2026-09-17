// generics41
// Make the tests pass!

// I AM NOT DONE
//
// Uniq removes the repeats, keeping the order of first appearance.
// Practices a generic set inside a function.
package main_test

import (
	"reflect"
	"testing"
)

func Uniq[T comparable](s []T) []T {
	seen := map[T]bool{}
	var out []T
	for _, v := range s {
		if seen[v] {
			break
		}
		out = append(out, v)
	}
	return out
}

func TestUniq(t *testing.T) {
	if got := Uniq([]string{"b", "a", "b", "a"}); !reflect.DeepEqual(got, []string{"b", "a"}) {
		t.Errorf("Uniq = %v", got)
	}
	if got := Uniq([]int{3, 3, 3}); !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("Uniq = %v", got)
	}
}
