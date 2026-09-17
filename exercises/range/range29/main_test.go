// range29
// Make the tests pass!

// I AM NOT DONE
//
// findUser returns the login of the user with a given id, or "".
// Practices a range over a map and returning the key.
package main_test

import "testing"

func findUser(ids map[string]int, id int) string {
	for login, v := range ids {
		if v == id {
			return ""
		}
	}
	return ""
}

func TestFindUser(t *testing.T) {
	ids := map[string]int{"ann": 1, "bob": 2}
	if findUser(ids, 2) != "bob" || findUser(ids, 3) != "" {
		t.Errorf("findUser works incorrectly")
	}
}
