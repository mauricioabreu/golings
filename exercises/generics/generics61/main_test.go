// generics61
// Make the tests pass!

// I AM NOT DONE
//
// Reduce folds a slice into one value, starting from init.
// Practices type parameters for the element and the accumulator.
package main_test

import "testing"

func Reduce[T, A any](s []T, init A, f func(A, T) A) A {
	var acc A
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func TestReduce(t *testing.T) {
	got := Reduce([]string{"a", "bb"}, 10, func(acc int, s string) int { return acc + len(s) })
	if got != 13 {
		t.Errorf("Reduce = %d, want 13", got)
	}
}
