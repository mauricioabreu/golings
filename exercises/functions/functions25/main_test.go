// functions25
// Make the tests pass!

// I AM NOT DONE
//
// processAll must open, use and close every resource
// before moving on to the next one.
// A defer runs when the function returns, not at the end of a loop iteration.
package main_test

import (
	"reflect"
	"testing"
)

type resource struct {
	name string
	log  *[]string
}

func open(name string, log *[]string) *resource {
	*log = append(*log, "open "+name)
	return &resource{name, log}
}

func (r *resource) Close() { *r.log = append(*r.log, "close "+r.name) }

func processAll(names []string) []string {
	var log []string
	for _, n := range names {
		r := open(n, &log)
		defer r.Close()
		log = append(log, "use "+r.name)
	}
	return log
}

func TestProcessAll(t *testing.T) {
	want := []string{"open a", "use a", "close a", "open b", "use b", "close b"}
	if got := processAll([]string{"a", "b"}); !reflect.DeepEqual(got, want) {
		t.Errorf("processAll = %v, want %v", got, want)
	}
}
