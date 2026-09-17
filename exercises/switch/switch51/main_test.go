// switch51
// Make the tests pass!

// I AM NOT DONE
//
// detectKind returns the kind of a file through a named result,
// replacing an unknown kind with "unknown". Right now it returns an empty string.
// := in a switch initializer declares a new variable.
package main_test

import (
	"strings"
	"testing"
)

func detect(name string) string {
	if strings.HasSuffix(name, ".go") {
		return "go"
	}
	return ""
}

func detectKind(name string) (kind string) {
	switch kind := detect(name); kind {
	case "":
		kind = "unknown"
	}
	return
}

func TestDetectKind(t *testing.T) {
	if got := detectKind("main.go"); got != "go" {
		t.Errorf("detectKind(main.go) = %q", got)
	}
	if got := detectKind("notes.txt"); got != "unknown" {
		t.Errorf("detectKind(notes.txt) = %q, want unknown", got)
	}
}
