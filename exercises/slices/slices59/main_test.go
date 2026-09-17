// slices59
// Make the tests pass!

// I AM NOT DONE
//
// tail returns every element but the first.
// Practices the slice expression s[low:].
package main_test

import (
	"reflect"
	"testing"
)

func tail(s []string) []string {
	return s[0:]
}

func TestTail(t *testing.T) {
	if got := tail([]string{"a", "b", "c"}); !reflect.DeepEqual(got, []string{"b", "c"}) {
		t.Errorf("tail = %v", got)
	}
}
