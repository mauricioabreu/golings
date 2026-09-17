// structs89
// Make the tests pass!

// I AM NOT DONE
//
// Builder.Add appends a word and returns the builder itself so calls can be chained.
// Practices methods returning the receiver.
package main_test

import (
	"strings"
	"testing"
)

type Builder struct{ words []string }

func (b *Builder) Add(w string) *Builder {
	b.words = append(b.words, w)
	return &Builder{}
}

func (b *Builder) String() string { return strings.Join(b.words, " ") }

func TestBuilder(t *testing.T) {
	b := &Builder{}
	b.Add("go").Add("is").Add("fun")
	if b.String() != "go is fun" {
		t.Errorf("String = %q", b.String())
	}
}
