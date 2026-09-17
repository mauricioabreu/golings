// structs55
// Make the tests pass!

// I AM NOT DONE
//
// sameOrder compares two orders. The code does not compile: the struct has a slice field.
// == is only available for structs whose fields are all comparable.
package main_test

import (
	"reflect"
	"testing"
)

type Order struct {
	ID    int
	Items []string
}

func sameOrder(a, b Order) bool {
	return a == b
}

func TestSameOrder(t *testing.T) {
	_ = reflect.DeepEqual
	if !sameOrder(Order{1, []string{"a"}}, Order{1, []string{"a"}}) || sameOrder(Order{1, []string{"a"}}, Order{1, []string{"b"}}) {
		t.Errorf("sameOrder works incorrectly")
	}
}
