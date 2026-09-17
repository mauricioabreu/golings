// structs47
// Make the tests pass!

// I AM NOT DONE
//
// NewServer returns a server on port 8080 with a timeout of 30.
// Practices constructor functions.
package main_test

import "testing"

type Server struct {
	Port    int
	Timeout int
}

func NewServer() *Server {
	return &Server{Port: 8080}
}

func TestNewServer(t *testing.T) {
	s := NewServer()
	if s.Port != 8080 || s.Timeout != 30 {
		t.Errorf("NewServer = %+v", *s)
	}
}
