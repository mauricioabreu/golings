// switch99
// Make the tests pass!

// I AM NOT DONE
//
// double doubles an int or a float64 and returns it as a float64.
// The code does not compile: in a branch listing two types v stays an interface{}.
// In a case with several types the variable keeps the type of the switch expression.
package main_test

import "testing"

func double(x interface{}) float64 {
	switch v := x.(type) {
	case int, float64:
		return float64(v) * 2
	}
	return 0
}

func TestDouble(t *testing.T) {
	if got := double(21); got != 42 {
		t.Errorf("double(21) = %v", got)
	}
	if got := double(1.25); got != 2.5 {
		t.Errorf("double(1.25) = %v", got)
	}
	if got := double("x"); got != 0 {
		t.Errorf("double(x) = %v", got)
	}
}
