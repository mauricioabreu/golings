// structs5
// Make the tests pass!

// I AM NOT DONE
//
// The code tries to change a field of a struct returned by a function.
// It does not compile: the result of a call is not addressable.
// Only addressable values can be assigned to.
package main_test

import "testing"

type Options struct {
	Retries int
	Verbose bool
}

func defaults() Options { return Options{Retries: 3} }

func verboseDefaults() Options {
	defaults().Verbose = true
	return defaults()
}

func TestVerboseDefaults(t *testing.T) {
	if o := verboseDefaults(); !o.Verbose || o.Retries != 3 {
		t.Errorf("verboseDefaults = %+v", o)
	}
}
