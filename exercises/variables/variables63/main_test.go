// variables63
// Make the tests pass!

// I AM NOT DONE
//
// The functions must return values of type int and float64 respectively.
// The type of a variable is inferred from the literal, and right now it is the wrong one.
// Practices type inference with := from untyped constants.
package main_test

import (
	"fmt"
	"testing"
)

func answer() interface{} {
	v := 42.0
	return v
}

func ratio() interface{} {
	v := 1 / 2
	return v
}

func TestTypes(t *testing.T) {
	if got := fmt.Sprintf("%T %v", answer(), answer()); got != "int 42" {
		t.Errorf("answer() = %s, want int 42", got)
	}
	if got := fmt.Sprintf("%T %v", ratio(), ratio()); got != "float64 0.5" {
		t.Errorf("ratio() = %s, want float64 0.5", got)
	}
}
