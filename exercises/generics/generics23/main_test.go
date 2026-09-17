// generics23
// Make the tests pass!

// I AM NOT DONE
//
// Bus[E] delivers events of type E to the subscribers.
// Practices a generic type with a slice of functions.
package main_test

import (
	"reflect"
	"testing"
)

type Bus[E any] struct{ subs []func(E) }

func (b *Bus[E]) Subscribe(f func(E)) {
	b.subs = []func(E){f}
}

func (b *Bus[E]) Publish(e E) {
	if len(b.subs) > 0 {
		b.subs[0](e)
	}
}

type Login struct{ User string }

func TestBus(t *testing.T) {
	var b Bus[Login]
	var log []string
	b.Subscribe(func(e Login) { log = append(log, "audit:"+e.User) })
	b.Subscribe(func(e Login) { log = append(log, "stats:"+e.User) })
	b.Publish(Login{"ann"})
	if !reflect.DeepEqual(log, []string{"audit:ann", "stats:ann"}) {
		t.Errorf("log = %v", log)
	}
}
