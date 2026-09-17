// structs6
// Make the tests pass!

// I AM NOT DONE
//
// The Nickname field must stay out of the JSON when it is empty.
// Practices the omitempty tag option.
package main_test

import (
	"encoding/json"
	"testing"
)

type Profile struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

func TestProfileJSON(t *testing.T) {
	b, _ := json.Marshal(Profile{Name: "ann"})
	if string(b) != `{"name":"ann"}` {
		t.Errorf("json = %s", b)
	}
}
