// algorithms149
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: DFS over a grid. Recolor the region of the same color connected
// to the cell (sr, sc) with newColor.
// Expected asymptotics: O(r·c) time, O(r·c) space.
package main_test

import (
	"reflect"
	"testing"
)

func floodFill(image [][]int, sr, sc, newColor int) [][]int {
	return nil
}

func TestFloodFill(t *testing.T) {
	img := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	want := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
	if got := floodFill(img, 1, 1, 2); !reflect.DeepEqual(got, want) {
		t.Errorf("floodFill = %v, want %v", got, want)
	}
	same := [][]int{{0, 0}, {0, 0}}
	if got := floodFill(same, 0, 0, 0); !reflect.DeepEqual(got, [][]int{{0, 0}, {0, 0}}) {
		t.Errorf("same color fill = %v", got)
	}
	single := [][]int{{5}}
	if got := floodFill(single, 0, 0, 7); got[0][0] != 7 {
		t.Errorf("single cell = %v", got)
	}
}
