// maps44
// Make the tests pass!

// I AM NOT DONE
//
// seen marks the values it has met in a map[interface{}]bool.
// For slices the program panics at run time.
// An interface{} key has to carry a comparable dynamic type.
package main_test

import (
	"fmt"
	"reflect"
	"testing"
)

func distinct(vals []interface{}) int {
	seen := map[interface{}]bool{}
	for _, v := range vals {
		seen[v] = true
	}
	return len(seen)
}

func TestDistinct(t *testing.T) {
	_, _ = fmt.Sprint, reflect.TypeOf
	vals := []interface{}{1, "1", []int{1, 2}, []int{1, 2}, 1}
	if got := distinct(vals); got != 3 {
		t.Errorf("distinct = %d, want 3", got)
	}
}
