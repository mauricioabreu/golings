// structs83
// Make the tests pass!

// I AM NOT DONE
//
// The Notify setting is true by default; an explicit false has to reach the JSON
// while an unset value must not. With omitempty the false is lost.
// omitempty drops zero values, false included.
package main_test

import (
	"encoding/json"
	"testing"
)

type Settings struct {
	Notify bool `json:"notify,omitempty"`
}

func TestSettingsJSON(t *testing.T) {
	var unset Settings
	b, _ := json.Marshal(unset)
	if string(b) != `{}` {
		t.Errorf("unset = %s", b)
	}
	off := false
	var s Settings
	_ = json.Unmarshal([]byte(`{"notify":false}`), &s)
	b, _ = json.Marshal(s)
	if string(b) != `{"notify":false}` {
		t.Errorf("explicit false = %s", b)
	}
	_ = off
}
