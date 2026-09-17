// generics3
// Make the tests pass!

// I AM NOT DONE
//
// Map applies a function to every element of a slice of any type.
// The code does not compile: the type parameter is declared without a constraint.
// Practices the [T any] syntax.
package main_test

import (
	"reflect"
	"strconv"
	"testing"
)

func Map[T, U](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

func TestMap(t *testing.T) {
	if got := Map([]int{1, 2}, strconv.Itoa); !reflect.DeepEqual(got, []string{"1", "2"}) {
		t.Errorf("Map = %v", got)
	}
}
