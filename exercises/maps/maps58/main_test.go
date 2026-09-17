// maps58
// Make the tests pass!

// I AM NOT DONE
//
// toBytes turns a size with a unit into bytes using a table of factors.
// Practices a map as a table of constants.
package main_test

import "testing"

var units = map[string]int{
	"B":  1,
	"KB": 1 << 10,
	"MB": 1 << 21,
}

func toBytes(n int, unit string) int {
	return n * units[unit]
}

func TestToBytes(t *testing.T) {
	if toBytes(3, "KB") != 3072 || toBytes(2, "MB") != 2097152 || toBytes(5, "B") != 5 {
		t.Errorf("toBytes works incorrectly")
	}
}
