// concurrent56
// Make the tests pass!

// I AM NOT DONE
//
// seen marks identifiers in a sync.Map and reports whether one has
// been seen before.
// Practices LoadOrStore.
package main_test

import (
	"sync"
	"testing"
)

func seen(m *sync.Map, id string) bool {
	_, loaded := m.LoadOrStore(id, true)
	return !loaded
}

func TestSeen(t *testing.T) {
	var m sync.Map
	if seen(&m, "a") || !seen(&m, "a") || seen(&m, "b") {
		t.Errorf("seen works incorrectly")
	}
}
