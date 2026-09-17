// structs94
// Make the tests pass!

// I AM NOT DONE
//
// A Clone of the configuration has to be independent, and changes to the copy show in the original.
// Copying a struct copies a map by reference.
package main_test

import "testing"

type Config struct {
	Name string
	Opts map[string]string
}

func (c Config) Clone() Config {
	return c
}

func TestClone(t *testing.T) {
	c := Config{"base", map[string]string{"debug": "off"}}
	d := c.Clone()
	d.Opts["debug"] = "on"
	if c.Opts["debug"] != "off" {
		t.Errorf("original changed: %v", c.Opts)
	}
}
