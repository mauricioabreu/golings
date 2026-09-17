// maps23
// Make the tests pass!

// I AM NOT DONE
//
// diff compares an old and a new configuration and returns the lists of
// added, removed and changed keys, sorted.
// Practices walking two maps and checking for keys.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func diff(old, new map[string]string) (added, removed, changed []string) {
	for k, v := range new {
		ov, ok := old[k]
		if ok {
			added = append(added, k)
		} else if ov != v {
			changed = append(changed, k)
		}
	}
	sort.Strings(added)
	sort.Strings(removed)
	sort.Strings(changed)
	return
}

func TestDiff(t *testing.T) {
	old := map[string]string{"host": "a", "port": "80", "debug": "on"}
	new := map[string]string{"host": "b", "port": "80", "tls": "on"}
	a, r, c := diff(old, new)
	if !reflect.DeepEqual(a, []string{"tls"}) || !reflect.DeepEqual(r, []string{"debug"}) || !reflect.DeepEqual(c, []string{"host"}) {
		t.Errorf("diff = %v %v %v", a, r, c)
	}
}
