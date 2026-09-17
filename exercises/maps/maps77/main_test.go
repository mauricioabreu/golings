// maps77
// Make the tests pass!

// I AM NOT DONE
//
// portFrom pulls the port out of a JSON configuration parsed into a map[string]interface{}.
// The type assertion to int never succeeds.
// encoding/json decodes numbers into an interface{} as float64.
package main_test

import (
	"encoding/json"
	"testing"
)

func portFrom(data []byte) (int, bool) {
	var cfg map[string]interface{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return 0, false
	}
	p, ok := cfg["port"].(int)
	return p, ok
}

func TestPortFrom(t *testing.T) {
	if p, ok := portFrom([]byte(`{"port": 8080}`)); !ok || p != 8080 {
		t.Errorf("portFrom = %d, %v", p, ok)
	}
	if _, ok := portFrom([]byte(`{"port": "80"}`)); ok {
		t.Errorf("string port should be rejected")
	}
}
