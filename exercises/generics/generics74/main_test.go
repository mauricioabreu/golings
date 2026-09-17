// generics74
// Make the tests pass!

// I AM NOT DONE
//
// Parse[T] builds a value of the wanted type from a string. A call without a type
// argument does not compile: the compiler does not infer T from the result.
// Inference works from the arguments only.
package main_test

import (
	"fmt"
	"testing"
)

func Parse[T any](s string) (T, error) {
	var v T
	_, err := fmt.Sscan(s, &v)
	return v, err
}

func port() int {
	p, _ := Parse("8080")
	return p
}

func TestPort(t *testing.T) {
	if port() != 8080 {
		t.Errorf("port = %d", port())
	}
}
