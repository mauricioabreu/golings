// algorithms101
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: Tarjan's algorithm. Find all edges of an undirected graph
// whose removal increases the number of connected components.
// Return the edges with the smaller vertex first, sorted.
// Expected asymptotics: O(V + E) time, O(V + E) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func criticalConnections(n int, connections [][2]int) [][2]int {
	return nil
}

func TestCriticalConnections(t *testing.T) {
	_ = sort.Ints
	got := criticalConnections(4, [][2]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}})
	if !reflect.DeepEqual(got, [][2]int{{1, 3}}) {
		t.Errorf("criticalConnections = %v", got)
	}
	got = criticalConnections(2, [][2]int{{0, 1}})
	if !reflect.DeepEqual(got, [][2]int{{0, 1}}) {
		t.Errorf("single edge = %v", got)
	}
	got = criticalConnections(3, [][2]int{{0, 1}, {1, 2}, {2, 0}})
	if len(got) != 0 {
		t.Errorf("cycle has no bridges, got %v", got)
	}
	if got := criticalConnections(1, nil); len(got) != 0 {
		t.Errorf("single node = %v", got)
	}
}
