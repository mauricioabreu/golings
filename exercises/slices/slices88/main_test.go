// slices88
// Make the tests pass!

// I AM NOT DONE
//
// zip combines names and ages into a slice of structs, as long as the shorter slice.
// Practices walking two slices side by side.
package main_test

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func zip(names []string, ages []int) []Person {
	out := make([]Person, len(names))
	for i := range names {
		out[i] = Person{names[i], ages[i]}
	}
	return out
}

func TestZip(t *testing.T) {
	got := zip([]string{"ann", "bob", "cid"}, []int{30, 40})
	if !reflect.DeepEqual(got, []Person{{"ann", 30}, {"bob", 40}}) {
		t.Errorf("zip = %v", got)
	}
}
