// structs63
// Make the tests pass!

// I AM NOT DONE
//
// Job holds a channel for cancellation, and json.Marshal fails because of it.
// The json:"-" tag keeps a field out of the serialization.
package main_test

import (
	"encoding/json"
	"testing"
)

type Job struct {
	ID     int           `json:"id"`
	Cancel chan struct{} `json:"cancel"`
}

func TestJobJSON(t *testing.T) {
	b, err := json.Marshal(Job{ID: 7, Cancel: make(chan struct{})})
	if err != nil || string(b) != `{"id":7}` {
		t.Errorf("json = %s, %v", b, err)
	}
}
