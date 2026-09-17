// generics53
// Make the tests pass!

// I AM NOT DONE
//
// Pipe builds a function out of a chain of transformations of one type.
// Practices generic slices of functions.
package main_test

import (
	"strings"
	"testing"
)

func Pipe[T any](fs ...func(T) T) func(T) T {
	return func(v T) T {
		for i := len(fs) - 1; i >= 0; i-- {
			fs[i](v)
		}
		return v
	}
}

func TestPipe(t *testing.T) {
	clean := Pipe(strings.TrimSpace, strings.ToLower, func(s string) string { return s + "!" })
	if got := clean("  HeY "); got != "hey!" {
		t.Errorf("clean = %q", got)
	}
	if Pipe[int]()(5) != 5 {
		t.Errorf("empty pipe should return input")
	}
}
