// structs69
// Make the tests pass!

// I AM NOT DONE
//
// Report embeds Author and Editor, and both have a Name field.
// The code does not compile: r.Name is ambiguous.
// Promoted fields at the same depth clash.
package main_test

import "testing"

type Author struct{ Name string }
type Editor struct{ Name string }

type Report struct {
	Author
	Editor
}

func byline(r Report) string {
	return "by " + r.Name
}

func TestByline(t *testing.T) {
	r := Report{Author{"ann"}, Editor{"bob"}}
	if byline(r) != "by ann" {
		t.Errorf("byline = %q", byline(r))
	}
}
