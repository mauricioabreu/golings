// structs45
// Make the tests pass!

// I AM NOT DONE
//
// Notifier embeds the Sender interface. The struct is built without a Sender,
// and the method call panics.
// An embedded interface is a field, and it is nil by default.
package main_test

import "testing"

type Sender interface{ Send(msg string) bool }

type Notifier struct {
	Sender
	Channel string
}

func (n Notifier) Notify(msg string) bool {
	return n.Send("[" + n.Channel + "] " + msg)
}

type okSender struct{}

func (okSender) Send(string) bool { return true }

func TestNotifier(t *testing.T) {
	if (Notifier{Channel: "ops"}).Notify("hi") {
		t.Errorf("Notify without sender should return false")
	}
	if !(Notifier{okSender{}, "ops"}).Notify("hi") {
		t.Errorf("Notify with sender should return true")
	}
}
