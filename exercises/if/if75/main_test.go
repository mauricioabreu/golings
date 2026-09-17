// if75
// Make the tests pass!

// I AM NOT DONE
//
// isImage must return true for files with the .png extension.
// Practices strings.HasSuffix in a condition.
package main_test

import (
	"strings"
	"testing"
)

func isImage(name string) bool {
	if strings.Contains(name, "png") {
		return true
	}
	return false
}

func TestIsImage(t *testing.T) {
	cases := map[string]bool{"cat.png": true, "png_notes.txt": false, "photo.jpg": false, "a.png.bak": false}
	for in, want := range cases {
		if got := isImage(in); got != want {
			t.Errorf("isImage(%q) = %v, want %v", in, got, want)
		}
	}
}
