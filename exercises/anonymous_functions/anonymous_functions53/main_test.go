// anonymous_functions53
// Make the tests pass!

// I AM NOT DONE
//
// splitList splits a string on ';' and ','.
// Practices a predicate literal for strings.FieldsFunc.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func splitList(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == ';'
	})
}

func TestSplitList(t *testing.T) {
	if got := splitList("a,b;c;;d"); !reflect.DeepEqual(got, []string{"a", "b", "c", "d"}) {
		t.Errorf("splitList = %v", got)
	}
}
