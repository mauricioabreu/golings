// structs43
// Make the tests pass!

// I AM NOT DONE
//
// After the Unmarshal the token field is still empty, although the key is in the JSON.
// encoding/json ignores unexported fields, tag or no tag.
package main_test

import (
	"encoding/json"
	"testing"
)

type Session struct {
	User  string `json:"user"`
	token string `json:"token"`
}

func tokenOf(data []byte) string {
	var s Session
	_ = json.Unmarshal(data, &s)
	return s.Token
}

func TestTokenOf(t *testing.T) {
	if got := tokenOf([]byte(`{"user":"ann","token":"xyz"}`)); got != "xyz" {
		t.Errorf("tokenOf = %q", got)
	}
}
