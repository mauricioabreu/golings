// range101
// Make the tests pass!

// I AM NOT DONE
//
// report prints the marks of the pupils alphabetically, with the marks of each in ascending order.
// Practices a range over sorted keys plus sorting the values.
package main_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func report(grades map[string][]int) string {
	names := make([]string, 0, len(grades))
	for n := range grades {
		names = append(names, n)
	}
	var lines []string
	for n, g := range grades {
		lines = append(lines, fmt.Sprintf("%s:%v", n, g))
	}
	return strings.Join(lines, ";")
}

func TestReport(t *testing.T) {
	_ = sort.Ints
	grades := map[string][]int{"bob": {5, 3}, "ann": {4, 2, 5}, "cid": {1}}
	for i := 0; i < 20; i++ {
		if got := report(grades); got != "ann:[2 4 5];bob:[3 5];cid:[1]" {
			t.Fatalf("report = %q", got)
		}
	}
	if grades["bob"][0] != 5 {
		t.Errorf("input modified: %v", grades["bob"])
	}
}
