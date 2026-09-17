// slices20
// Make the tests pass!

// I AM NOT DONE
//
// sameTags compares sets of tags; nil and an empty slice have to count as equal.
// reflect.DeepEqual tells nil and an empty slice apart.
package main_test

import (
	"reflect"
	"testing"
)

func sameTags(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

func TestSameTags(t *testing.T) {
	_ = reflect.DeepEqual
	if !sameTags(nil, []string{}) {
		t.Errorf("nil and empty tags should be equal")
	}
	if !sameTags([]string{"go"}, []string{"go"}) || sameTags([]string{"go"}, nil) {
		t.Errorf("sameTags works incorrectly")
	}
}
