// anonymous_functions64
// Make the tests pass!

// I AM NOT DONE
//
// sortBy sorts by a list of comparison literals: the next one is used
// only when the earlier ones could not tell the elements apart.
// Practices a slice of comparator functions.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

type Emp struct {
	Dept string
	Age  int
	Name string
}

type cmp func(a, b Emp) int

func sortBy(es []Emp, cmps ...cmp) {
	sort.Slice(es, func(i, j int) bool {
		for _, c := range cmps {
			return c(es[i], es[j]) < 0
		}
		return false
	})
}

func byDept(a, b Emp) int {
	if a.Dept == b.Dept {
		return 0
	}
	if a.Dept < b.Dept {
		return -1
	}
	return 1
}

func TestSortBy(t *testing.T) {
	es := []Emp{{"it", 30, "c"}, {"hr", 40, "a"}, {"it", 25, "b"}}
	sortBy(es, byDept, func(a, b Emp) int { return a.Age - b.Age })
	names := []string{es[0].Name, es[1].Name, es[2].Name}
	if !reflect.DeepEqual(names, []string{"a", "b", "c"}) {
		t.Errorf("order = %v", names)
	}
}
