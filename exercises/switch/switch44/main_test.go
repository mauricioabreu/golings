// switch44
// Make the tests pass!

// I AM NOT DONE
//
// group returns "image", "text", "json" or "binary" for a MIME type.
// Practices a tagless switch with functions from strings.
package main_test

import (
	"strings"
	"testing"
)

func group(mime string) string {
	switch {
	case strings.HasSuffix(mime, "image/"):
		return "image"
	case strings.HasPrefix(mime, "text/"):
		return "text"
	}
	return "binary"
}

func TestGroup(t *testing.T) {
	cases := map[string]string{
		"image/png": "image", "text/html": "text", "application/json": "json",
		"application/ld+json": "json", "application/zip": "binary",
	}
	for in, want := range cases {
		if got := group(in); got != want {
			t.Errorf("group(%s) = %s, want %s", in, got, want)
		}
	}
}
