// variables53
// Make the tests pass!

// I AM NOT DONE
//
// The status codes must be OK=0, (1 skipped), NotFound=2, Forbidden=3.
// Practices skipping iota values with _.
package main_test

import "testing"

const (
	OK = iota
	NotFound
	Forbidden
)

func TestCodes(t *testing.T) {
	if OK != 0 || NotFound != 2 || Forbidden != 3 {
		t.Errorf("OK=%d NotFound=%d Forbidden=%d, want 0 2 3", OK, NotFound, Forbidden)
	}
}
