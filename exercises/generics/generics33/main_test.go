// generics33
// Make the tests pass!

// I AM NOT DONE
//
// MaxOf returns the maximum of a slice. The code does not compile:
// comparable does not support <.
// comparable only gives == and !=.
package main_test

import "testing"

func MaxOf[T comparable](s []T) T {
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func TestMaxOf(t *testing.T) {
	if MaxOf([]int{3, 9, 2}) != 9 || MaxOf([]string{"b", "c", "a"}) != "c" {
		t.Errorf("MaxOf works incorrectly")
	}
}
