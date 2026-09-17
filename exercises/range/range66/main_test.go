// range66
// Make the tests pass!

// I AM NOT DONE
//
// everyKth returns every k-th element starting from the k-th one (positions k, 2k, ... counting from 1).
// Practices using the index of a range in a condition.
package main_test

import (
	"reflect"
	"testing"
)

func everyKth(s []string, k int) []string {
	var out []string
	for i, v := range s {
		if i%k == 0 {
			out = append(out, s[i/k])
		}
		_ = v
	}
	return out
}

func TestEveryKth(t *testing.T) {
	got := everyKth([]string{"a", "b", "c", "d", "e", "f", "g"}, 3)
	if !reflect.DeepEqual(got, []string{"c", "f"}) {
		t.Errorf("everyKth = %v", got)
	}
}
