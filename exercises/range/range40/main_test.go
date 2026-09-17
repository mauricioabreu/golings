// range40
// Make the tests pass!

// I AM NOT DONE
//
// lastNonEmpty must return the last non-empty string.
// The function always returns an empty string.
// := in a range header declares new variables.
package main_test

import "testing"

func lastNonEmpty(lines []string) string {
	var last string
	for _, last := range lines {
		if last == "" {
			continue
		}
	}
	return last
}

func TestLastNonEmpty(t *testing.T) {
	if got := lastNonEmpty([]string{"a", "b", ""}); got != "b" {
		t.Errorf("lastNonEmpty = %q, want b", got)
	}
}
