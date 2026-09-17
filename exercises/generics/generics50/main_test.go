// generics50
// Make the tests pass!

// I AM NOT DONE
//
// JoinAll calls String on every element. For a []Point the code does not compile:
// String is declared on *Point, so Point does not satisfy the constraint.
// The method set of the type argument has to hold the methods of the constraint.
package main_test

import (
	"fmt"
	"strings"
	"testing"
)

type Stringer interface{ String() string }

type Point struct{ X, Y int }

func (p *Point) String() string { return fmt.Sprintf("(%d,%d)", p.X, p.Y) }

func JoinAll[T Stringer](items []T) string {
	parts := make([]string, len(items))
	for i, it := range items {
		parts[i] = it.String()
	}
	return strings.Join(parts, " ")
}

func TestJoinAll(t *testing.T) {
	if got := JoinAll([]Point{{1, 2}, {3, 4}}); got != "(1,2) (3,4)" {
		t.Errorf("JoinAll = %q", got)
	}
}
