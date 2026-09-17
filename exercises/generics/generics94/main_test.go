// generics94
// Make the tests pass!

// I AM NOT DONE
//
// Partition splits a slice in two by a predicate.
// Practices generic multiple results.
package main_test

import (
	"reflect"
	"testing"
)

func Partition[T any](s []T, pred func(T) bool) (yes, no []T) {
	for _, v := range s {
		if pred(v) {
			yes = append(yes, v)
		}
		no = append(no, v)
	}
	return
}

func TestPartition(t *testing.T) {
	yes, no := Partition([]string{"go", "", "c"}, func(s string) bool { return s != "" })
	if !reflect.DeepEqual(yes, []string{"go", "c"}) || !reflect.DeepEqual(no, []string{""}) {
		t.Errorf("Partition = %q, %q", yes, no)
	}
}
