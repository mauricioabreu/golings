// generics63
// Make the tests pass!

// I AM NOT DONE
//
// BFS returns the vertices in breadth first order for a graph whose vertices
// are of any comparable type.
// Practices generic maps and a queue.
package main_test

import (
	"reflect"
	"testing"
)

func BFS[T comparable](graph map[T][]T, start T) []T {
	visited := map[T]bool{start: true}
	queue := []T{start}
	var order []T
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, n := range graph[v] {
			queue = append(queue, n)
		}
	}
	return order
}

func TestBFS(t *testing.T) {
	g := map[int][]int{1: {2, 3}, 2: {1, 4}, 3: {4}, 4: {1}}
	if got := BFS(g, 1); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("BFS = %v", got)
	}
}
