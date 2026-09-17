// anonymous_functions8
// Make the tests pass!

// I AM NOT DONE
//
// defaultConfig has to be a Config value computed by a literal at initialization.
// The code does not compile: the literal is never called.
// func() T { ... } and func() T { ... }() are different things.
package main_test

import "testing"

type Config struct {
	Workers int
	Name    string
}

var defaultConfig Config = func() Config {
	c := Config{Name: "svc"}
	c.Workers = 4
	return c
}

func TestDefaultConfig(t *testing.T) {
	if defaultConfig.Workers != 4 || defaultConfig.Name != "svc" {
		t.Errorf("defaultConfig = %+v", defaultConfig)
	}
}
