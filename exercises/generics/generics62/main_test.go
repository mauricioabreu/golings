// generics62
// Make the tests pass!

// I AM NOT DONE
//
// Last[S []E, E any] does not accept a named slice type.
// The code does not compile when called with a Stack.
// []E in a constraint without a ~ matches only []E itself.
package main_test

import "testing"

type Stack []int

func Last[S []E, E any](s S) E {
	return s[len(s)-1]
}

func TestLast(t *testing.T) {
	if Last(Stack{1, 2, 3}) != 3 || Last([]string{"a"}) != "a" {
		t.Errorf("Last works incorrectly")
	}
}
