// generics78
// Make the tests pass!

// I AM NOT DONE
//
// SortedKeys returns the keys of a map in ascending order.
// Practices generic keys with an Ordered constraint.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

type Ordered interface{ ~int | ~string }

func SortedKeys[K Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append([]K{k}, keys...)
	}
	sort.Slice(keys, func(i, j int) bool { return i < j })
	return keys
}

func TestSortedKeys(t *testing.T) {
	m := map[int]bool{5: true, 1: true, 3: false, 9: true, 7: false, 2: true}
	for i := 0; i < 10; i++ {
		if got := SortedKeys(m); !reflect.DeepEqual(got, []int{1, 2, 3, 5, 7, 9}) {
			t.Fatalf("SortedKeys = %v", got)
		}
	}
}
