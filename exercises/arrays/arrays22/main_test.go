// arrays22
// Make the tests pass!

// I AM NOT DONE
//
// firstTwo must return a slice holding the first two elements of an array.
// Practices slicing an array.
package main_test

import (
	"reflect"
	"testing"
)

func firstTwo(a *[4]string) []string {
	return a[1:3]
}

func TestFirstTwo(t *testing.T) {
	a := [4]string{"w", "x", "y", "z"}
	if got := firstTwo(&a); !reflect.DeepEqual(got, []string{"w", "x"}) {
		t.Errorf("firstTwo = %v", got)
	}
}
