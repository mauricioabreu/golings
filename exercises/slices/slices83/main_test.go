// slices83
// Make the tests pass!

// I AM NOT DONE
//
// The employees are already sorted by name. byDept sorts them by department,
// and the alphabetical order has to survive inside every department.
// sort.Slice is not guaranteed to be stable.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

type Emp struct {
	Name string
	Dept int
}

func byDept(es []Emp) {
	sort.Slice(es, func(i, j int) bool { return es[i].Dept < es[j].Dept })
}

func TestByDept(t *testing.T) {
	var es []Emp
	for i := 0; i < 40; i++ {
		es = append(es, Emp{Name: string(rune('A' + i)), Dept: i % 3})
	}
	byDept(es)
	var names []string
	for _, e := range es {
		if e.Dept == 1 {
			names = append(names, e.Name)
		}
	}
	want := []string{"B", "E", "H", "K", "N", "Q", "T", "W", "Z", "]", "`", "c", "f"}
	if !reflect.DeepEqual(names, want) {
		t.Errorf("dept 1 order = %v, want %v", names, want)
	}
}
