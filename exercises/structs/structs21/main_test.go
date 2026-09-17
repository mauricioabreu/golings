// structs21
// Make the tests pass!

// I AM NOT DONE
//
// Two Event structs with a Payload interface{} field are compared with ==.
// When the payload holds a slice, the comparison panics.
// == on structs with interface fields may panic.
package main_test

import (
	"reflect"
	"testing"
)

type Event struct {
	Kind    string
	Payload interface{}
}

func sameEvent(a, b Event) bool {
	return a == b
}

func TestSameEvent(t *testing.T) {
	_ = reflect.DeepEqual
	if !sameEvent(Event{"x", 1}, Event{"x", 1}) {
		t.Errorf("scalar payloads should match")
	}
	if !sameEvent(Event{"batch", []int{1, 2}}, Event{"batch", []int{1, 2}}) {
		t.Errorf("slice payloads should match")
	}
}
