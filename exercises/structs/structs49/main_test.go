// structs49
// Make the tests pass!

// I AM NOT DONE
//
// Temperature.String prints the value with one decimal and a unit, and "-" when there is none.
// Practices String() and fmt with different values.
package main_test

import (
	"fmt"
	"testing"
)

type Temperature struct {
	Value float64
	Valid bool
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v°C", t.Value)
}

func TestTemperatureString(t *testing.T) {
	got := fmt.Sprintf("%v %v", Temperature{21.456, true}, Temperature{})
	if got != "21.5°C -" {
		t.Errorf("got %q", got)
	}
}
