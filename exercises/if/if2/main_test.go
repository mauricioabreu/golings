// if2
// Make me compile!

// I AM NOT DONE
package main_test

import "testing"

func fooIfFizz(fizzish string) string {
	// When the input is fizz return foo
	// When the input is fuzz return bar
	// When the input is neither fizz or fuzz return baz
	if fizzish == "fizz" {
		return "foo"
	} else {
		return "complete me"
	}
}

func TestFooForFizz(t *testing.T) {
	result := fooIfFizz("fizz")
	if result != "foo" {
		t.Errorf("should be 'foo' but got %s", result)
	}
}

func TestBarForFuzz(t *testing.T) {
	result := fooIfFizz("fuzz")
	if result != "bar" {
		t.Errorf("should be 'bar' but got %s", result)
	}
}

func TestDefaultForBazz(t *testing.T) {
	result := fooIfFizz("random stuff")
	if result != "baz" {
		t.Errorf("should be 'baz' but got %s", result)
	}
}
