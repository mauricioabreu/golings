// anonymous_functions84
// Make the tests pass!

// I AM NOT DONE
//
// Option has to hold a function computing the default value rather than the result.
// Right now the expensive function is called when the option is built, even when a value is set.
// A literal defers a computation; a call does not.
package main_test

import "testing"

type Option struct {
	set   bool
	value int
	def   int
}

func NewOption(value int, set bool, compute func() int) Option {
	return Option{set: set, value: value, def: compute()}
}

func (o Option) Value() int {
	if o.set {
		return o.value
	}
	return o.def
}

func TestOption(t *testing.T) {
	calls := 0
	compute := func() int { calls++; return 99 }
	a := NewOption(5, true, compute)
	b := NewOption(0, false, compute)
	if a.Value() != 5 || calls != 0 {
		t.Errorf("explicit value: %d, calls = %d", a.Value(), calls)
	}
	if b.Value() != 99 || calls != 1 {
		t.Errorf("default value: %d, calls = %d", b.Value(), calls)
	}
}
