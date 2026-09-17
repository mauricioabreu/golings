// structs102
// Make the tests pass!

// I AM NOT DONE
//
// Point.String formats a point as "(x, y)", and fmt uses that method.
// Practices the fmt.Stringer interface.
package main_test

import (
	"fmt"
	"testing"
)

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("[%d; %d]", p.X, p.Y)
}

func TestPointString(t *testing.T) {
	if got := fmt.Sprint(Point{1, 2}); got != "(1, 2)" {
		t.Errorf("Sprint = %q", got)
	}
}
