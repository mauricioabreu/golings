// maps20
// Make the tests pass!

// I AM NOT DONE
//
// addScore stores a pupil's mark for a subject in a map[pupil]map[subject]mark.
// For a new pupil the write panics.
// Practices initializing a nested map.
package main_test

import "testing"

func addScore(book map[string]map[string]int, student, subject string, score int) {
	book[student][subject] = score
}

func TestAddScore(t *testing.T) {
	book := map[string]map[string]int{}
	addScore(book, "ann", "math", 5)
	addScore(book, "ann", "art", 4)
	addScore(book, "bob", "math", 3)
	if book["ann"]["math"] != 5 || book["ann"]["art"] != 4 || book["bob"]["math"] != 3 {
		t.Errorf("book = %v", book)
	}
}
