// maps32
// Make the tests pass!

// I AM NOT DONE
//
// adjacency builds the adjacency list of a directed graph,
// with the neighbours of every vertex sorted.
// Practices map[string][]string.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func adjacency(edges [][2]string) map[string][]string {
	adj := map[string][]string{}
	for _, e := range edges {
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	return adj
}

func TestAdjacency(t *testing.T) {
	got := adjacency([][2]string{{"a", "c"}, {"a", "b"}, {"b", "c"}})
	if !reflect.DeepEqual(got, map[string][]string{"a": {"b", "c"}, "b": {"c"}}) {
		t.Errorf("adjacency = %v", got)
	}
}
