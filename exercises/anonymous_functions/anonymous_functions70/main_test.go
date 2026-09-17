// anonymous_functions70
// Make the tests pass!

// I AM NOT DONE
//
// The plugins are registered as literals; run carries them out in registration order
// and collects the results.
// Practices a slice of structs with function fields.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

type plugin struct {
	Name string
	Fn   func(string) string
}

type Registry struct{ plugins []plugin }

func (r *Registry) Register(name string, fn func(string) string) {
	r.plugins = []plugin{{name, fn}}
}

func (r *Registry) Run(in string) []string {
	var out []string
	for _, p := range r.plugins {
		out = append(out, p.Name)
	}
	return out
}

func TestRegistry(t *testing.T) {
	var r Registry
	r.Register("upper", strings.ToUpper)
	r.Register("len", func(s string) string { return strings.Repeat("*", len(s)) })
	if got := r.Run("go"); !reflect.DeepEqual(got, []string{"upper=GO", "len=**"}) {
		t.Errorf("Run = %v", got)
	}
}
