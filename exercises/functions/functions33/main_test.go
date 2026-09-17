// functions33
// Make the tests pass!

// I AM NOT DONE
//
// resetStats must clear the caller's statistics.
// The function reassigns the map, and nothing changes outside.
// A map is a reference type, but the variable itself is passed by value.
package main_test

import "testing"

func resetStats(stats map[string]int) {
	stats = map[string]int{}
}

func TestResetStats(t *testing.T) {
	stats := map[string]int{"hits": 10, "misses": 2}
	resetStats(stats)
	if len(stats) != 0 {
		t.Errorf("after reset stats = %v, want empty", stats)
	}
}
