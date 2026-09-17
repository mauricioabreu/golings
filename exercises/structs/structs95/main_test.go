// structs95
// Make the tests pass!

// I AM NOT DONE
//
// newBook builds a book with the given fields.
// Practices a struct literal with field names.
package main_test

import "testing"

type Book struct {
	Title string
	Pages int
}

func newBook() Book {
	return Book{Title: "Go", Pages: 0}
}

func TestNewBook(t *testing.T) {
	if b := newBook(); b.Title != "Go" || b.Pages != 380 {
		t.Errorf("newBook = %+v", b)
	}
}
