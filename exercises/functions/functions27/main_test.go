// functions27
// Make the tests pass!

// I AM NOT DONE
//
// processItems must report how many items were processed when it returns.
// The report always says 0.
// With defer s.Method() a value receiver is copied straight away.
package main_test

import (
	"fmt"
	"testing"
)

type Stats struct{ Count int }

func (s Stats) Report(out *string) {
	*out = fmt.Sprintf("processed %d", s.Count)
}

func processItems(items []string, out *string) {
	var s Stats
	defer s.Report(out)
	for range items {
		s.Count++
	}
}

func TestProcessItems(t *testing.T) {
	var report string
	processItems([]string{"a", "b", "c"}, &report)
	if report != "processed 3" {
		t.Errorf("report = %q, want %q", report, "processed 3")
	}
}
