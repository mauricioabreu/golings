// functions6
// Make the tests pass!

// I AM NOT DONE
//
// sortPeople must order people by age, and by name when the ages are equal.
// Practices a comparison function with several criteria.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func sortPeople(ps []Person) {
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Name < ps[j].Name
	})
}

func TestSortPeople(t *testing.T) {
	ps := []Person{{"Zoe", 30}, {"Adam", 30}, {"Bob", 25}}
	sortPeople(ps)
	want := []Person{{"Bob", 25}, {"Adam", 30}, {"Zoe", 30}}
	if !reflect.DeepEqual(ps, want) {
		t.Errorf("sortPeople = %v, want %v", ps, want)
	}
}
