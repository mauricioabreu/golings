// slices19
// Make the tests pass!

// I AM NOT DONE
//
// removeAt returns the slice without the element at index i; the original slice may be changed.
// Practices append(s[:i], s[i+1:]...).
package main_test

import (
	"reflect"
	"testing"
)

func removeAt(s []string, i int) []string {
	return append(s[:i], s[i:]...)
}

func TestRemoveAt(t *testing.T) {
	if got := removeAt([]string{"a", "b", "c"}, 1); !reflect.DeepEqual(got, []string{"a", "c"}) {
		t.Errorf("removeAt = %v", got)
	}
}
