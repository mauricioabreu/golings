// arrays4
// Make the tests pass!

// I AM NOT DONE
//
// The colour names live in an array of size numColors. A Purple colour was added,
// but the size of the array is written as a number and the test fails.
// Practices tying the size of an array to an iota counter constant.
package main_test

import "testing"

type Color int

const (
	Red Color = iota
	Green
	Blue
	Purple
)

var colorNames = [3]string{"red", "green", "blue"}

func (c Color) String() string {
	if c < 0 || int(c) >= len(colorNames) {
		return "unknown"
	}
	return colorNames[c]
}

func TestColorNames(t *testing.T) {
	if Purple.String() != "purple" || Red.String() != "red" || Color(42).String() != "unknown" {
		t.Errorf("names: %s %s %s", Purple, Red, Color(42))
	}
}
