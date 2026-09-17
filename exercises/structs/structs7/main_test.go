// structs7
// Make the tests pass!

// I AM NOT DONE
//
// Service embeds Logger and uses its Log method with the name of the service as a prefix.
// Practices embedding to reuse behaviour.
package main_test

import (
	"reflect"
	"testing"
)

type Logger struct {
	prefix string
	lines  []string
}

func (l *Logger) Log(msg string) { l.lines = append(l.lines, l.prefix+msg) }

type Service struct {
	log  Logger
	Name string
}

func NewService(name string) *Service {
	return &Service{Name: name}
}

func TestService(t *testing.T) {
	s := NewService("auth")
	s.Log("started")
	if !reflect.DeepEqual(s.lines, []string{"[auth] started"}) {
		t.Errorf("lines = %v", s.lines)
	}
}
