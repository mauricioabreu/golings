// range34
// Make the tests pass!

// I AM NOT DONE
//
// mergeSlots joins sorted busy slots [start, end) that overlap or touch.
// Practices a range that changes the last element of the result.
package main_test

import (
	"reflect"
	"testing"
)

func mergeSlots(slots [][2]int) [][2]int {
	var out [][2]int
	for _, s := range slots {
		n := len(out)
		if n > 0 && s[0] < out[n-1][1] {
			last := out[n-1]
			last[1] = s[1]
			continue
		}
		out = append(out, s)
	}
	return out
}

func TestMergeSlots(t *testing.T) {
	got := mergeSlots([][2]int{{9, 10}, {10, 12}, {11, 11}, {13, 14}})
	if !reflect.DeepEqual(got, [][2]int{{9, 12}, {13, 14}}) {
		t.Errorf("mergeSlots = %v", got)
	}
}
