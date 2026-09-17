// functions42
// Make the tests pass!

// I AM NOT DONE
//
// NewServer takes option functions that change the defaults.
// The options are built but never applied, and one of them writes to the wrong field.
// Practices the functional options pattern.
package main_test

import "testing"

type Server struct {
	Port    int
	Timeout int
}

type Option func(*Server)

func WithPort(p int) Option { return func(s *Server) { s.Port = p } }

func WithTimeout(t int) Option { return func(s *Server) { s.Port = t } }

func NewServer(opts ...Option) *Server {
	s := &Server{Port: 80, Timeout: 30}
	return s
}

func TestNewServer(t *testing.T) {
	s := NewServer(WithPort(8080), WithTimeout(5))
	if s.Port != 8080 || s.Timeout != 5 {
		t.Errorf("NewServer = %+v, want Port 8080 Timeout 5", *s)
	}
	d := NewServer()
	if d.Port != 80 || d.Timeout != 30 {
		t.Errorf("default server = %+v", *d)
	}
}
