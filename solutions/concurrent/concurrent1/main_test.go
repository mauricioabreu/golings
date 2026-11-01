// concurrent1
// Make the tests pass!

package main_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

func TestPrinter(t *testing.T) {
	var buf bytes.Buffer
	print(&buf)

	out := buf.String()

	for i := 0; i < 3; i++ {
		want := fmt.Sprintf("Hello from goroutine %d!", i)
		if !bytes.Contains([]byte(out), []byte(want)) {
			t.Errorf("Output did not contain expected string. Wanted: %q, Got: %q", want, out)
		}
	}
}

func print(buf *bytes.Buffer) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	goroutines := 3

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			fmt.Fprintf(buf, "Hello from goroutine %d!\n", i)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
}
