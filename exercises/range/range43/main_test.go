// range43
// Make the tests pass!

// I AM NOT DONE
//
// expensive returns a new map holding only the items above a threshold.
// Practices a range over a map while building a new one.
package main_test

import (
	"reflect"
	"testing"
)

func expensive(prices map[string]int, min int) map[string]int {
	out := map[string]int{}
	for k, v := range prices {
		if v < min {
			out[k] = v
		}
	}
	return prices
}

func TestExpensive(t *testing.T) {
	got := expensive(map[string]int{"tv": 500, "pen": 5, "car": 9000}, 100)
	if !reflect.DeepEqual(got, map[string]int{"tv": 500, "car": 9000}) {
		t.Errorf("expensive = %v", got)
	}
}
