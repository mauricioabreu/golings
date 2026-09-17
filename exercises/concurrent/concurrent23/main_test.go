// concurrent23
// Make the tests pass!

// I AM NOT DONE
//
// The configuration lives in an atomic.Value; an update stores a value
// of another type, and Store panics.
// atomic.Value asks for one concrete type across every Store.
package main_test

import (
	"sync/atomic"
	"testing"
)

type Config struct{ Level int }

type Holder struct{ v atomic.Value }

func (h *Holder) Set(level int) {
	if level == 0 {
		h.v.Store(Config{})
		return
	}
	h.v.Store(&Config{Level: level})
}

func (h *Holder) Level() int { return h.v.Load().(*Config).Level }

func TestHolder(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic: %v", r)
		}
	}()
	var h Holder
	h.Set(3)
	h.Set(0)
	if h.Level() != 0 {
		t.Errorf("Level = %d", h.Level())
	}
}
