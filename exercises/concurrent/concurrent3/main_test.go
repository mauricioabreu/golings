// concurrent3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSendAndReceive(t *testing.T) {
	var buf bytes.Buffer

	messages := make(chan string)
	sendAndReceive(&buf, messages)

	got := buf.String()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func sendAndReceive(buf *bytes.Buffer, messages chan string) {
	go func() {
		messages <- "Hello"
		messages <- "World"
		close(messages)
	}()

	greeting := <-messages
	fmt.Fprint(buf, greeting)

	// Here we just receive the first message
	// Consider using a for-range loop to iterate over the messages
	_, ok := <-messages
	if !ok {
		fmt.Fprint(buf, "Channel is closed")
	}
}
