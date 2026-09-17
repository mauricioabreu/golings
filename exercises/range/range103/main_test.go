// range103
// Make the tests pass!

// I AM NOT DONE
//
// resetAll clears the counters in every map of a slice by replacing them with new empty maps.
// Assigning to the loop variable does not change the slice elements.
// The v of a range is a copy of the map reference, and reassigning it is invisible outside.
package main_test

import "testing"

func resetAll(stats []map[string]int) {
	for _, m := range stats {
		m = map[string]int{}
		_ = m
	}
}

func TestResetAll(t *testing.T) {
	a := map[string]int{"x": 1}
	stats := []map[string]int{a, {"y": 2}}
	resetAll(stats)
	if len(stats[0]) != 0 || len(stats[1]) != 0 {
		t.Errorf("stats = %v", stats)
	}
	if a["x"] != 1 {
		t.Errorf("old map should stay untouched: %v", a)
	}
}
