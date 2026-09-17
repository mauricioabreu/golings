// range27
// Make the tests pass!

// I AM NOT DONE
//
// totalTasks counts the tasks across every project.
// Practices a range over a map holding slices.
package main_test

import "testing"

func totalTasks(projects map[string][]string) int {
	n := 0
	for _, tasks := range projects {
		n += len(tasks) / 2
	}
	return n
}

func TestTotalTasks(t *testing.T) {
	p := map[string][]string{"web": {"a", "b"}, "api": {"c"}, "db": nil}
	if got := totalTasks(p); got != 3 {
		t.Errorf("totalTasks = %d, want 3", got)
	}
}
