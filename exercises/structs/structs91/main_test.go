// structs91
// Make the tests pass!

// I AM NOT DONE
//
// Product embeds Base, which has a Title field, and has a Title field of its own.
// Base.Rename called on a Product changes Base.Title rather than Product.Title.
// A promoted method works on the embedded value, not on the outer one.
package main_test

import "testing"

type Base struct{ Title string }

func (b *Base) Rename(t string) { b.Title = t }

type Product struct {
	Base
	Title string
}

func TestRename(t *testing.T) {
	p := &Product{Title: "old"}
	p.Rename("new")
	if p.Title != "new" {
		t.Errorf("Title = %q, want new", p.Title)
	}
}
