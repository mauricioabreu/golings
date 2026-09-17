// maps92
// Make the tests pass!

// I AM NOT DONE
//
// mergeSum joins two maps, adding the values of matching keys,
// and leaves the input maps alone.
// Practices building a new map and accumulating values.
package main_test

import (
	"reflect"
	"testing"
)

func mergeSum(a, b map[string]int) map[string]int {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func TestMergeSum(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 10, "z": 3}
	got := mergeSum(a, b)
	if !reflect.DeepEqual(got, map[string]int{"x": 1, "y": 12, "z": 3}) {
		t.Errorf("mergeSum = %v", got)
	}
	if !reflect.DeepEqual(a, map[string]int{"x": 1, "y": 2}) {
		t.Errorf("input modified: %v", a)
	}
}
