// variables15
// Make the tests pass!

// I AM NOT DONE
//
// This function must log the FINAL value of the counter when it returns.
// The initial value is logged instead.
// The arguments of a deferred call are evaluated at the defer statement.
package main_test

import "testing"

func process(items []string, log *[]int) {
	count := 0
	defer record(log, count)
	for range items {
		count++
	}
}

func record(log *[]int, v int) {
	*log = append(*log, v)
}

func TestProcess(t *testing.T) {
	var log []int
	process([]string{"a", "b", "c"}, &log)
	if len(log) != 1 || log[0] != 3 {
		t.Errorf("log = %v, want [3]", log)
	}
}
