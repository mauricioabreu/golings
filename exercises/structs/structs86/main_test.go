// structs86
// Make the tests pass!

// I AM NOT DONE
//
// Record embeds Meta and Audit, and both have an ID field tagged "id".
// encoding/json quietly drops the clashing fields.
// Practices the field promotion rules of encoding/json.
package main_test

import (
	"encoding/json"
	"testing"
)

type Meta struct {
	ID int `json:"id"`
}

type Audit struct {
	ID int `json:"id"`
}

type Record struct {
	Meta
	Audit
}

func TestRecordJSON(t *testing.T) {
	b, _ := json.Marshal(Record{Meta{1}, Audit{2}})
	if string(b) != `{"meta_id":1,"audit_id":2}` {
		t.Errorf("json = %s", b)
	}
}
