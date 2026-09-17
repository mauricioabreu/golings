// generics44
// Make the tests pass!

// I AM NOT DONE
//
// Container[T].Format has to print int values in a special way.
// The code does not compile: a method cannot be declared for Container[int] alone.
// Methods are declared for every instantiation at once.
package main_test

import (
	"fmt"
	"testing"
)

type Container[T any] struct{ V T }

func (c Container[T]) Format() string { return fmt.Sprint(c.V) }

func (c Container[int]) Format() string { return fmt.Sprintf("#%d", c.V) }

func TestFormat(t *testing.T) {
	if (Container[int]{7}).Format() != "#7" || (Container[string]{"x"}).Format() != "x" {
		t.Errorf("Format works incorrectly")
	}
}
