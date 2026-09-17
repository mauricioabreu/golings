// anonymous_functions27
// Make the tests pass!

// I AM NOT DONE
//
// closeAll defers the closing of every resource with a literal.
// Every deferred literal closes the last resource.
// Deferred literals read the variable at the moment they run.
package main_test

import (
	"reflect"
	"testing"
)

func closeAll(names []string) (closed []string) {
	var current string
	func() {
		for _, n := range names {
			current = n
			defer func() { closed = append(closed, current) }()
		}
	}()
	return closed
}

func TestCloseAll(t *testing.T) {
	if got := closeAll([]string{"db", "cache", "file"}); !reflect.DeepEqual(got, []string{"file", "cache", "db"}) {
		t.Errorf("closeAll = %v", got)
	}
}
