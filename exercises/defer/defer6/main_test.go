// defer6
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"testing"
)

func readFile(name string) error {
	if name == "missing.toml" {
		return errors.New("file does not exist")
	}
	return nil
}

// loadConfig reads the file. When that fails, the error has to say which file
// it was: "loading missing.toml: file does not exist". err is a named result,
// so a deferred function can still change it after the return statement.
func loadConfig(name string) (err error) {
	// TODO: Add the file name to the error on the way out, using a deferred function. Do not touch it when there is no error.
	return readFile(name)
}

func TestLoadConfig(t *testing.T) {
	if err := loadConfig("app.toml"); err != nil {
		t.Errorf("loadConfig(app.toml) = %v, want nil", err)
	}
}

func TestLoadConfigAddsTheFileName(t *testing.T) {
	err := loadConfig("missing.toml")

	if err == nil || err.Error() != "loading missing.toml: file does not exist" {
		t.Errorf("got %v, want %q", err, "loading missing.toml: file does not exist")
	}
}
