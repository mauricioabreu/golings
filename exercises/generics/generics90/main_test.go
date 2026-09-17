// generics90
// Make the tests pass!

// I AM NOT DONE
//
// Wrapper[T] tries to embed T so that its methods are promoted.
// The code does not compile: a type parameter cannot be embedded.
// Practices the limits of generic structs.
package main_test

import "testing"

type Named interface{ Name() string }

type User struct{ N string }

func (u User) Name() string { return u.N }

type Wrapper[T Named] struct {
	T
	Tag string
}

func (w Wrapper[T]) Label() string { return w.Tag + ":" + w.Name() }

func TestLabel(t *testing.T) {
	w := Wrapper[User]{User{"ann"}, "admin"}
	if w.Label() != "admin:ann" {
		t.Errorf("Label = %q", w.Label())
	}
}
