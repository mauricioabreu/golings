// slices51
// Make the tests pass!

// I AM NOT DONE
//
// parentDir returns every part of a path but the last, joined with "/".
// Practices strings.Split and slicing the result.
package main_test

import (
	"strings"
	"testing"
)

func parentDir(path string) string {
	parts := strings.Split(path, "/")
	return strings.Join(parts[1:], "/")
}

func TestParentDir(t *testing.T) {
	if got := parentDir("usr/local/bin"); got != "usr/local" {
		t.Errorf("parentDir = %q", got)
	}
}
