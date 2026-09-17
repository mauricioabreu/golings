// slices89
// Make the tests pass!

// I AM NOT DONE
//
// Stack is built on a slice. Its Push method has to add an element,
// yet the stack stays empty.
// An append in a method with a value receiver changes a copy of the header.
package main_test

import (
	"reflect"
	"testing"
)

type Stack []int

func (s Stack) Push(v int) {
	s = append(s, v)
}

func TestStackPush(t *testing.T) {
	var s Stack
	s.Push(1)
	s.Push(2)
	if !reflect.DeepEqual([]int(s), []int{1, 2}) {
		t.Errorf("stack = %v, want [1 2]", s)
	}
}
