// generics37
// Make the tests pass!

// I AM NOT DONE
//
// Frequencies counts how often every value occurs.
// Practices map[T]int in a generic function.
package main_test

import (
	"reflect"
	"testing"
)

func Frequencies[T comparable](s []T) map[T]int {
	var m map[T]int
	for _, v := range s {
		m[v] = 1
	}
	return m
}

func TestFrequencies(t *testing.T) {
	if got := Frequencies([]rune("abca")); !reflect.DeepEqual(got, map[rune]int{'a': 2, 'b': 1, 'c': 1}) {
		t.Errorf("Frequencies = %v", got)
	}
}
