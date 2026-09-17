// anonymous_functions41
// Make the tests pass!

// I AM NOT DONE
//
// sortByPriority sorts tasks by a priority from a map ("high" < "mid" < "low");
// unknown priorities go last.
// Practices a comparison literal using a captured table.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

type Task struct{ Name, Priority string }

func sortByPriority(ts []Task) {
	rank := map[string]int{"high": 0, "mid": 1, "low": 2}
	weight := func(p string) int {
		return rank[p]
	}
	sort.SliceStable(ts, func(i, j int) bool {
		return weight(ts[i].Priority) < weight(ts[j].Priority)
	})
}

func TestSortByPriority(t *testing.T) {
	ts := []Task{{"a", "?"}, {"b", "low"}, {"c", "high"}, {"d", "mid"}}
	sortByPriority(ts)
	var names []string
	for _, x := range ts {
		names = append(names, x.Name)
	}
	if !reflect.DeepEqual(names, []string{"c", "d", "b", "a"}) {
		t.Errorf("order = %v", names)
	}
}
