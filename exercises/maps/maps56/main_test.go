// maps56
// Make the tests pass!

// I AM NOT DONE
//
// renameKeys returns a new map with the keys replaced from a table;
// keys with no replacement are kept as they are.
// Practices comma-ok when substituting.
package main_test

import (
	"reflect"
	"testing"
)

func renameKeys(m map[string]int, names map[string]string) map[string]int {
	out := make(map[string]int, len(m))
	for k, v := range m {
		out[names[k]] = v
	}
	return out
}

func TestRenameKeys(t *testing.T) {
	got := renameKeys(map[string]int{"usr": 1, "pwd": 2, "id": 3}, map[string]string{"usr": "user", "pwd": "password"})
	if !reflect.DeepEqual(got, map[string]int{"user": 1, "password": 2, "id": 3}) {
		t.Errorf("renameKeys = %v", got)
	}
}
