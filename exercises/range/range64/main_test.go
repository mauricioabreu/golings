// range64
// Make the tests pass!

// I AM NOT DONE
//
// Pipeline reads events from the events channel. When the channel was never built,
// a range over a nil channel blocks forever.
// Operations on a nil channel block for good.
package main_test

import (
	"testing"
	"time"
)

type Pipeline struct {
	events chan string
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Emit(e string) { p.events <- e }
func (p *Pipeline) Close()        { close(p.events) }

func (p *Pipeline) Collect() []string {
	var out []string
	for e := range p.events {
		out = append(out, e)
	}
	return out
}

func TestPipeline(t *testing.T) {
	done := make(chan []string, 1)
	go func() {
		p := NewPipeline()
		p.Emit("start")
		p.Emit("stop")
		p.Close()
		done <- p.Collect()
	}()
	select {
	case got := <-done:
		if len(got) != 2 || got[0] != "start" {
			t.Errorf("Collect = %v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("pipeline is stuck: is the channel created?")
	}
}
