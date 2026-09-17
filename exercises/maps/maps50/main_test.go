// maps50
// Make the tests pass!

// I AM NOT DONE
//
// indexByID builds a map from an id to a pointer to the element of the original slice.
// Changes made through the map have to show in the slice.
// Practices map[int]*T and the addresses of slice elements.
package main_test

import "testing"

type Item struct {
	ID    int
	Title string
}

func indexByID(items []Item) map[int]*Item {
	idx := make(map[int]*Item, len(items))
	for _, it := range items {
		copyIt := it
		idx[it.ID] = &copyIt
	}
	return idx
}

func TestIndexByID(t *testing.T) {
	items := []Item{{1, "a"}, {2, "b"}}
	idx := indexByID(items)
	idx[2].Title = "B"
	if items[1].Title != "B" {
		t.Errorf("items[1].Title = %q, want B", items[1].Title)
	}
}
