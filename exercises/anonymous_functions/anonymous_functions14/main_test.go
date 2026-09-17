// anonymous_functions14
// Make the tests pass!

// I AM NOT DONE
//
// transformAll applies a callback to every string, and the strings do not change:
// the result of the callback is thrown away.
// A literal returning a value does not change its argument.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func transformAll(items []string, f func(string) string) {
	for i := range items {
		f(items[i])
	}
}

func TestTransformAll(t *testing.T) {
	s := []string{"a", "b"}
	transformAll(s, func(x string) string { return strings.ToUpper(x) + "!" })
	if !reflect.DeepEqual(s, []string{"A!", "B!"}) {
		t.Errorf("transformAll = %v", s)
	}
}
