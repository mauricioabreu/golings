// generics99
// Make the tests pass!

// I AM NOT DONE
//
// MinMaxBy returns the elements with the smallest and the largest key.
// Practices a generic function with a key function and an Ordered constraint.
package main_test

import "testing"

type Ordered interface{ ~int | ~float64 | ~string }

func MinMaxBy[T any, K Ordered](s []T, key func(T) K) (lo, hi T) {
	lo, hi = s[0], s[0]
	for _, v := range s[1:] {
		if key(v) < key(lo) {
			hi = v
		}
	}
	return
}

type City struct {
	Name string
	Pop  int
}

func TestMinMaxBy(t *testing.T) {
	cities := []City{{"b", 5}, {"a", 1}, {"c", 9}}
	lo, hi := MinMaxBy(cities, func(c City) int { return c.Pop })
	if lo.Name != "a" || hi.Name != "c" {
		t.Errorf("lo=%v hi=%v", lo, hi)
	}
}
