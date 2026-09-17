// structs101
// Make the tests pass!

// I AM NOT DONE
//
// ByAge implements sort.Interface for sorting people by age.
// Practices the Len, Less and Swap methods on a named slice.
package main_test

import (
	"sort"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByAge) Swap(i, j int)      { a[i] = a[j] }

func TestByAge(t *testing.T) {
	ps := []Person{{"a", 40}, {"b", 20}, {"c", 30}}
	sort.Sort(ByAge(ps))
	if ps[0].Name != "b" || ps[1].Name != "c" || ps[2].Name != "a" {
		t.Errorf("sorted = %v", ps)
	}
}
