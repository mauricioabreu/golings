// maps39
// Make the tests pass!

// I AM NOT DONE
//
// lookup pulls a value out of nested map[string]interface{} values by the path "a.b.c".
// Practices type assertions while walking down the levels.
package main_test

import (
	"strings"
	"testing"
)

func lookup(doc map[string]interface{}, path string) (interface{}, bool) {
	var cur interface{} = doc
	for _, part := range strings.Split(path, ".") {
		cur = doc[part]
	}
	return cur, true
}

func TestLookup(t *testing.T) {
	doc := map[string]interface{}{
		"db":    map[string]interface{}{"host": "localhost", "port": 5432},
		"debug": true,
	}
	if v, ok := lookup(doc, "db.port"); !ok || v != 5432 {
		t.Errorf("db.port = %v, %v", v, ok)
	}
	if _, ok := lookup(doc, "db.user"); ok {
		t.Errorf("db.user should be missing")
	}
	if _, ok := lookup(doc, "debug.level"); ok {
		t.Errorf("debug.level should be missing")
	}
}
