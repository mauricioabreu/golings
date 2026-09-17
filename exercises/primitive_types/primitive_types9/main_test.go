// primitive_types9
// Make the tests pass!

// I AM NOT DONE
//
// stripScheme must remove a leading "https://" from a URL when it is there.
// Practices strings.TrimPrefix.
package main_test

import (
	"strings"
	"testing"
)

func stripScheme(url string) string {
	return strings.TrimSuffix(url, "https://")
}

func TestStripScheme(t *testing.T) {
	cases := map[string]string{"https://go.dev": "go.dev", "go.dev": "go.dev"}
	for in, want := range cases {
		if got := stripScheme(in); got != want {
			t.Errorf("stripScheme(%q) = %q, want %q", in, got, want)
		}
	}
}
