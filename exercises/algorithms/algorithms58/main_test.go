// algorithms58
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: topological sort (Kahn's algorithm). Course a requires course b
// (pair [a, b]). Return an order for taking courses 0..n-1, or nil on a cycle.
// If several variants exist, take the course with the smallest number.
// Expected asymptotics: O(V + E) time, O(V + E) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func findOrder(n int, prereqs [][2]int) []int {
	return nil
}

func TestFindOrder(t *testing.T) {
	_ = sort.Ints
	if got := findOrder(2, [][2]int{{1, 0}}); !reflect.DeepEqual(got, []int{0, 1}) {
		t.Errorf("findOrder = %v", got)
	}
	if got := findOrder(4, [][2]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}); !reflect.DeepEqual(got, []int{0, 1, 2, 3}) {
		t.Errorf("findOrder = %v", got)
	}
	if got := findOrder(2, [][2]int{{1, 0}, {0, 1}}); got != nil {
		t.Errorf("cycle should give nil, got %v", got)
	}
	if got := findOrder(1, nil); !reflect.DeepEqual(got, []int{0}) {
		t.Errorf("single course = %v", got)
	}
}
