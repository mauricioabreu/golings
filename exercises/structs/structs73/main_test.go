// structs73
// Make the tests pass!

// I AM NOT DONE
//
// Category holds its parent category. The code does not compile:
// a struct cannot hold itself by value.
// Recursive structs call for pointers.
package main_test

import "testing"

type Category struct {
	Name   string
	Parent Category
}

func path(c *Category) string {
	if c.Parent == nil {
		return c.Name
	}
	return path(c.Parent) + "/" + c.Name
}

func TestPath(t *testing.T) {
	root := &Category{Name: "shop"}
	kids := &Category{Name: "toys", Parent: root}
	if got := path(kids); got != "shop/toys" {
		t.Errorf("path = %q", got)
	}
}
