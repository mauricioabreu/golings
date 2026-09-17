// generics4
// Make the tests pass!

// I AM NOT DONE
//
// Describe returns a description of a value depending on its type.
// The code does not compile: a type switch cannot be used on a type parameter value.
// For a type switch the value has to be converted to an interface{}.
package main_test

import (
	"strconv"
	"testing"
)

func Describe[T any](v T) string {
	switch x := v.(type) {
	case int:
		return "int " + strconv.Itoa(x)
	case string:
		return "string " + x
	}
	return "other"
}

func TestDescribe(t *testing.T) {
	if Describe(5) != "int 5" || Describe("go") != "string go" || Describe(1.5) != "other" {
		t.Errorf("Describe works incorrectly")
	}
}
