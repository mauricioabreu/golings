// range21
// Make the tests pass!

// I AM NOT DONE
//
// lowerKeys must return a map with the keys folded to lower case.
// Changing the variable k inside the loop does not change the keys of the map.
// The key variable of a range is a copy, not a reference to the key.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func lowerKeys(m map[string]int) map[string]int {
	for k := range m {
		k = strings.ToLower(k)
		_ = k
	}
	return m
}

func TestLowerKeys(t *testing.T) {
	got := lowerKeys(map[string]int{"Go": 1, "RUST": 2})
	if !reflect.DeepEqual(got, map[string]int{"go": 1, "rust": 2}) {
		t.Errorf("lowerKeys = %v", got)
	}
}
