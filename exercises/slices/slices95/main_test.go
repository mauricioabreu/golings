// slices95
// Make the tests pass!

// I AM NOT DONE
//
// collectIDs fills a slice built with a capacity by index and panics.
// A capacity does not allow reaching elements past the length.
package main_test

import (
	"reflect"
	"testing"
)

func collectIDs(users []map[string]int) []int {
	ids := make([]int, 0, len(users))
	for i, u := range users {
		ids[i] = u["id"]
	}
	return ids
}

func TestCollectIDs(t *testing.T) {
	got := collectIDs([]map[string]int{{"id": 3}, {"id": 7}})
	if !reflect.DeepEqual(got, []int{3, 7}) {
		t.Errorf("collectIDs = %v", got)
	}
}
