// switch57
// Make the tests pass!

// I AM NOT DONE
//
// mimeType returns the MIME type for a file extension.
// Practices a switch with an init statement.
package main_test

import (
	"path/filepath"
	"testing"
)

func mimeType(name string) string {
	switch ext := filepath.Ext(name); ext {
	case ".html":
		return "text/html"
	case "json":
		return "application/json"
	default:
		return "application/octet-stream"
	}
}

func TestMimeType(t *testing.T) {
	cases := map[string]string{"index.html": "text/html", "data.json": "application/json", "a.bin": "application/octet-stream"}
	for in, want := range cases {
		if got := mimeType(in); got != want {
			t.Errorf("mimeType(%s) = %s, want %s", in, got, want)
		}
	}
}
