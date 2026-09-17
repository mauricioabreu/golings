// anonymous_functions95
// Make the tests pass!

// I AM NOT DONE
//
// apply carries out an operation from a map of literals.
// Practices function literals as map values.
package main_test

import "testing"

var ops = map[string]func(int) int{
	"inc":    func(x int) int { return x + 1 },
	"double": func(x int) int { return x + 2 },
}

func TestOps(t *testing.T) {
	if ops["inc"](4) != 5 || ops["double"](4) != 8 || ops["double"](7) != 14 {
		t.Errorf("ops work incorrectly")
	}
}
