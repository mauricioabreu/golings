// structs57
// Make the tests pass!

// I AM NOT DONE
//
// decodeAll parses JSON records into one and the same variable.
// Fields missing from a record keep the values of the previous one.
// json.Unmarshal does not zero the fields that are absent from the input.
package main_test

import (
	"encoding/json"
	"testing"
)

type Item struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func decodeAll(lines []string) []Item {
	var out []Item
	var it Item
	for _, l := range lines {
		if err := json.Unmarshal([]byte(l), &it); err == nil {
			out = append(out, it)
		}
	}
	return out
}

func TestDecodeAll(t *testing.T) {
	got := decodeAll([]string{`{"name":"car","color":"red"}`, `{"name":"box"}`})
	if len(got) != 2 || got[1].Color != "" {
		t.Errorf("decodeAll = %+v", got)
	}
}
