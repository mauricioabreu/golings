// functions66
// Make the tests pass!

// I AM NOT DONE
//
// The calculator keeps its functions in a map keyed by the operator symbol.
// Two of the operators are wired to the wrong functions.
// Practices functions as map values.
package main_test

import "testing"

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }

var ops = map[string]func(int, int) int{
	"+": add,
	"-": add,
	"*": add,
}

func calc(a int, op string, b int) int {
	return ops[op](a, b)
}

func TestCalc(t *testing.T) {
	cases := []struct {
		a    int
		op   string
		b    int
		want int
	}{{2, "+", 3, 5}, {9, "-", 4, 5}, {3, "*", 4, 12}}
	for _, c := range cases {
		if got := calc(c.a, c.op, c.b); got != c.want {
			t.Errorf("calc(%d %s %d) = %d, want %d", c.a, c.op, c.b, got, c.want)
		}
	}
	_ = sub
}
