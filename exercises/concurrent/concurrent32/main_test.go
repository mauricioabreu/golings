// concurrent32
// Make the tests pass!

// I AM NOT DONE
//
// Client opens a connection lazily on first use from any goroutine;
// the connection has to be opened exactly once.
// Practices sync.Once in a method.
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

var dials int32

type Client struct {
	once sync.Once
	conn string
}

func (c *Client) Conn() string {
	if c.conn == "" {
		atomic.AddInt32(&dials, 1)
		c.conn = "tcp://db"
	}
	return c.conn
}

func TestClientConn(t *testing.T) {
	var c Client
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Conn()
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&dials) != 1 || c.Conn() != "tcp://db" {
		t.Errorf("dials = %d", dials)
	}
}
