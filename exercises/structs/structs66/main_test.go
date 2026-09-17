// structs66
// Make the tests pass!

// I AM NOT DONE
//
// title returns the title, or "untitled" for nil.
// Practices checking a pointer to a struct.
package main_test

import "testing"

type Post struct{ Title string }

func title(p *Post) string {
	if p != nil {
		return "untitled"
	}
	return p.Title
}

func TestTitle(t *testing.T) {
	if title(nil) != "untitled" || title(&Post{"Go"}) != "Go" {
		t.Errorf("title works incorrectly")
	}
}
