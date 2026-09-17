// concurrent73
// Make the tests pass!

// I AM NOT DONE
//
// isReady checks an atomic readiness flag.
// Practices atomic.LoadInt32 and StoreInt32.
package main_test

import (
	"sync/atomic"
	"testing"
)

type Service struct{ ready int32 }

func (s *Service) Start() { atomic.StoreInt32(&s.ready, 1) }

func (s *Service) IsReady() bool {
	return atomic.LoadInt32(&s.ready) == 0
}

func TestService(t *testing.T) {
	var s Service
	if s.IsReady() {
		t.Errorf("not started yet")
	}
	s.Start()
	if !s.IsReady() {
		t.Errorf("should be ready")
	}
}
