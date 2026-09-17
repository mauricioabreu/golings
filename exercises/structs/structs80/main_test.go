// structs80
// Make the tests pass!

// I AM NOT DONE
//
// fmt does not use String() when printing a value if the method is declared on the pointer.
// Practices the fmt.Stringer interface and method sets.
package main_test

import (
	"fmt"
	"testing"
)

type Card struct {
	Rank string
	Suit string
}

func (c *Card) String() string { return c.Rank + c.Suit }

func TestCardString(t *testing.T) {
	hand := []Card{{"A", "♠"}, {"10", "♥"}}
	if got := fmt.Sprint(hand); got != "[A♠ 10♥]" {
		t.Errorf("Sprint = %q", got)
	}
}
