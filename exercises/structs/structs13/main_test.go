// structs13
// Make the tests pass!

// I AM NOT DONE
//
// The statistics live in a map with struct keys, and json.Marshal does not
// support such a map and returns an error.
// JSON map keys have to be strings, numbers or a TextMarshaler.
package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Cell struct{ X, Y int }

func encode(counts map[Cell]int) (string, error) {
	b, err := json.Marshal(counts)
	_ = fmt.Sprint
	return string(b), err
}

func TestEncode(t *testing.T) {
	got, err := encode(map[Cell]int{{1, 2}: 3, {0, 0}: 1})
	if err != nil || got != `{"0,0":1,"1,2":3}` {
		t.Errorf("encode = %s, %v", got, err)
	}
}
