// anonymous_functions86
// Make the tests pass!

// I AM NOT DONE
//
// The configuration is loaded once through a sync.Once. The first attempt panics,
// and the Once never calls the literal again, leaving the configuration empty.
// Once counts a call as done even when the literal panicked.
package main_test

import (
	"sync"
	"testing"
)

type Loader struct {
	once  sync.Once
	value string
	load  func() string
}

func (l *Loader) Get() string {
	func() {
		defer func() { _ = recover() }()
		l.once.Do(func() {
			l.value = l.load()
		})
	}()
	return l.value
}

func TestLoader(t *testing.T) {
	l := &Loader{load: func() string { panic("no file") }}
	l.Get()
	if got := l.Get(); got != "default" {
		t.Errorf("Get = %q, want default", got)
	}
}
