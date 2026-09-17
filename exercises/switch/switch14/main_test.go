// switch14
// Make the tests pass!

// I AM NOT DONE
//
// features returns the features of a plan: "pro" includes everything in "plus",
// and "plus" everything in "basic". Right now "pro" is missing features.
// fallthrough has to be written out in every branch.
package main_test

import (
	"reflect"
	"testing"
)

func features(plan string) []string {
	var f []string
	switch plan {
	case "pro":
		f = append(f, "api")
	case "plus":
		f = append(f, "export")
		fallthrough
	case "basic":
		f = append(f, "editor")
	}
	return f
}

func TestFeatures(t *testing.T) {
	if got := features("pro"); !reflect.DeepEqual(got, []string{"api", "export", "editor"}) {
		t.Errorf("features(pro) = %v", got)
	}
	if got := features("plus"); !reflect.DeepEqual(got, []string{"export", "editor"}) {
		t.Errorf("features(plus) = %v", got)
	}
}
