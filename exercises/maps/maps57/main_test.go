// maps57
// Make the tests pass!

// I AM NOT DONE
//
// withoutSecrets must return the configuration without the secret keys
// and leave the original map alone. Right now the secrets vanish for the caller too.
// A map is passed to a function by reference.
package main_test

import (
	"strings"
	"testing"
)

func withoutSecrets(cfg map[string]string) map[string]string {
	for k := range cfg {
		if strings.HasPrefix(k, "secret_") {
			delete(cfg, k)
		}
	}
	return cfg
}

func TestWithoutSecrets(t *testing.T) {
	cfg := map[string]string{"host": "db", "secret_pass": "123"}
	public := withoutSecrets(cfg)
	if len(public) != 1 || public["host"] != "db" {
		t.Errorf("public = %v", public)
	}
	if cfg["secret_pass"] != "123" {
		t.Errorf("original config modified: %v", cfg)
	}
}
