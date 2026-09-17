// structs87
// Make the tests pass!

// I AM NOT DONE
//
// The email field does not reach the JSON because it is unexported.
// encoding/json only sees exported fields.
package main_test

import (
	"encoding/json"
	"testing"
)

type Contact struct {
	Name  string `json:"name"`
	email string `json:"email"`
}

func TestContactJSON(t *testing.T) {
	var c Contact
	_ = json.Unmarshal([]byte(`{"name":"ann","email":"a@x"}`), &c)
	b, _ := json.Marshal(c)
	if string(b) != `{"name":"ann","email":"a@x"}` {
		t.Errorf("json = %s", b)
	}
}
