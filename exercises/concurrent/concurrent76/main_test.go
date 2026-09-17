// concurrent76
// Make the tests pass!

// I AM NOT DONE
//
// Snapshot copies the struct along with its locked mutex; trying to
// lock the copy hangs forever.
// Practices that structs with a mutex must not be copied.
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Profile struct {
	mu   sync.Mutex
	Name string
	Age  int
}

type ProfileData struct {
	Name string
	Age  int
}

func (p *Profile) Snapshot() ProfileData {
	p.mu.Lock()
	defer p.mu.Unlock()
	cp := *p
	cp.mu.Lock()
	defer cp.mu.Unlock()
	return ProfileData{cp.Name, cp.Age}
}

func TestSnapshot(t *testing.T) {
	p := &Profile{Name: "ann", Age: 30}
	res := make(chan ProfileData, 1)
	go func() { res <- p.Snapshot() }()
	select {
	case d := <-res:
		if d.Name != "ann" || d.Age != 30 {
			t.Errorf("Snapshot = %+v", d)
		}
	case <-time.After(time.Second):
		t.Fatal("Snapshot deadlocked on a copied locked mutex")
	}
}
