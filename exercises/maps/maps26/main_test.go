// maps26
// Make the tests pass!

// I AM NOT DONE
//
// degrees counts the degree of every vertex of an undirected graph from a list of edges.
// Practices incrementing the counters of both ends of an edge.
package main_test

import (
	"reflect"
	"testing"
)

func degrees(edges [][2]string) map[string]int {
	d := map[string]int{}
	for _, e := range edges {
		d[e[0]] += 2
	}
	return d
}

func TestDegrees(t *testing.T) {
	got := degrees([][2]string{{"a", "b"}, {"b", "c"}, {"c", "c"}})
	if !reflect.DeepEqual(got, map[string]int{"a": 1, "b": 2, "c": 3}) {
		t.Errorf("degrees = %v", got)
	}
}
