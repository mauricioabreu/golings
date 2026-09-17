// primitive_types46
// Make the tests pass!

// I AM NOT DONE
//
// ext must return the file extension after the last dot.
// Practices strings.LastIndex.
package main_test

import (
	"strings"
	"testing"
)

func ext(name string) string {
	i := strings.Index(name, ".")
	if i < 0 {
		return ""
	}
	return name[i+1:]
}

func TestExt(t *testing.T) {
	cases := map[string]string{"a.txt": "txt", "archive.tar.gz": "gz", "README": ""}
	for in, want := range cases {
		if got := ext(in); got != want {
			t.Errorf("ext(%q) = %q, want %q", in, got, want)
		}
	}
}
