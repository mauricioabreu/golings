// concurrent44
// Make the tests pass!

// I AM NOT DONE
//
// The configuration is published through an atomic.Value, and Current returns it.
// Practices Store and Load with a type assertion.
package main_test

import (
	"sync/atomic"
	"testing"
)

type Settings struct{ Level int }

var current atomic.Value

func Current() Settings {
	return Settings{}
}

func TestCurrent(t *testing.T) {
	current.Store(Settings{Level: 3})
	if Current().Level != 3 {
		t.Errorf("Current = %+v", Current())
	}
}
