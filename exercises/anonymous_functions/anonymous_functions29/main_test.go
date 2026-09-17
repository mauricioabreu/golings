// anonymous_functions29
// Make the tests pass!

// I AM NOT DONE
//
// The literal has to count events in the original statistics map. After a rotation
// the stats variable points at a new map, and the literal writes there instead.
// A literal sees the new value of a captured variable.
package main_test

import "testing"

func run() (old, fresh map[string]int) {
	stats := map[string]int{}
	count := func(e string) { stats[e]++ }
	count("a")
	old = stats
	stats = map[string]int{}
	count("b")
	return old, stats
}

func TestRun(t *testing.T) {
	old, fresh := run()
	if old["a"] != 1 || old["b"] != 1 || len(fresh) != 0 {
		t.Errorf("old = %v, fresh = %v", old, fresh)
	}
}
