// structs24
// Make the tests pass!

// I AM NOT DONE
//
// NewRegistry builds a registry ready to use, and Register refuses duplicates.
// Practices initializing map fields in a constructor.
package main_test

import (
	"errors"
	"testing"
)

type Registry struct{ items map[string]int }

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Register(name string, v int) error {
	r.items[name] = v
	return nil
}

func TestRegistry(t *testing.T) {
	_ = errors.New
	r := NewRegistry()
	if err := r.Register("a", 1); err != nil {
		t.Fatalf("Register(a) = %v", err)
	}
	if err := r.Register("a", 2); err == nil || r.items["a"] != 1 {
		t.Errorf("duplicate registration should fail")
	}
}
