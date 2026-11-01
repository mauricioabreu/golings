package cmd

import (
	"bytes"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/fatih/color"
	. "github.com/onsi/gomega"
)

type counters struct{ run, list, hint atomic.Int32 }

func (c *counters) actions() watchActions {
	return watchActions{
		run:  func() { c.run.Add(1) },
		list: func() { c.list.Add(1) },
		hint: func() { c.hint.Add(1) },
	}
}

// startLoop runs watchLoop in the background and returns what the test needs
// to drive it and to know when it has returned.
func startLoop(t *testing.T, c *counters) (typed *io.PipeWriter, changes chan []string, done <-chan struct{}) {
	t.Helper()
	in, typed := io.Pipe()
	changes = make(chan []string)
	finished := make(chan struct{})
	t.Cleanup(func() { _ = typed.Close() })

	go func() {
		defer close(finished)
		watchLoop(in, changes, c.actions())
	}()
	return typed, changes, finished
}

// syncBuffer is a buffer the loop can write to while the test reads it.
type syncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *syncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

func captureOutput(t *testing.T) *syncBuffer {
	t.Helper()
	out := &syncBuffer{}
	prevOutput, prevNoColor := color.Output, color.NoColor
	color.Output, color.NoColor = out, true
	t.Cleanup(func() { color.Output, color.NoColor = prevOutput, prevNoColor })
	return out
}

func TestWatchLoopRunsOnceAtStartAndOncePerChange(t *testing.T) {
	g := NewWithT(t)
	var c counters
	_, changes, _ := startLoop(t, &c)

	g.Eventually(c.run.Load).Should(Equal(int32(1)), "runs the next exercise right away")
	changes <- []string{"exercises/a/main.go"}
	g.Eventually(c.run.Load).Should(Equal(int32(2)))
	changes <- []string{"exercises/a/main.go"}
	g.Eventually(c.run.Load).Should(Equal(int32(3)))
}

// Every command the learner typed used to start one more goroutine waiting for
// file changes, and none of them ever ended.
func TestTypingCommandsDoesNotLeakGoroutines(t *testing.T) {
	g := NewWithT(t)
	captureOutput(t)
	var c counters
	typed, changes, _ := startLoop(t, &c)
	g.Eventually(c.run.Load).Should(Equal(int32(1)))
	before := runtime.NumGoroutine()

	for range 30 {
		_, err := typed.Write([]byte("list\n"))
		g.Expect(err).NotTo(HaveOccurred())
	}
	g.Eventually(c.list.Load).Should(Equal(int32(30)))

	g.Expect(runtime.NumGoroutine()).To(BeNumerically("<", before+10))

	// ...and a change still runs the exercise exactly once.
	runsBefore := c.run.Load()
	changes <- []string{"exercises/a/main.go"}
	g.Eventually(c.run.Load).Should(Equal(runsBefore + 1))
	g.Consistently(c.run.Load, 200*time.Millisecond).Should(Equal(runsBefore + 1))
}

func TestWatchLoopCommands(t *testing.T) {
	g := NewWithT(t)
	out := captureOutput(t)
	var c counters
	typed, _, done := startLoop(t, &c)

	for _, line := range []string{"hint", "list", "  hint  ", "dance"} {
		_, err := typed.Write([]byte(line + "\n"))
		g.Expect(err).NotTo(HaveOccurred())
	}
	g.Eventually(c.hint.Load).Should(Equal(int32(2)))
	g.Eventually(c.list.Load).Should(Equal(int32(1)))
	g.Eventually(func() string { return out.String() }).Should(ContainSubstring("only list or hint commands are available"))
	g.Expect(done).NotTo(BeClosed())

	_, err := typed.Write([]byte("exit\n"))
	g.Expect(err).NotTo(HaveOccurred())
	g.Eventually(done).Should(BeClosed())
	g.Expect(out.String()).To(ContainSubstring("Bye by golings"))
}

func TestWatchLoopStopsOnQuit(t *testing.T) {
	g := NewWithT(t)
	captureOutput(t)
	var c counters
	typed, _, done := startLoop(t, &c)

	_, err := typed.Write([]byte("quit\n"))
	g.Expect(err).NotTo(HaveOccurred())

	g.Eventually(done).Should(BeClosed())
}

func TestWatchLoopStopsWhenInputEnds(t *testing.T) {
	g := NewWithT(t)
	var c counters
	typed, _, done := startLoop(t, &c)

	// Used to spin forever printing the end-of-file error.
	g.Expect(typed.Close()).To(Succeed())

	g.Eventually(done).Should(BeClosed())
}

func TestWatchLoopStopsWhenTheWatcherStops(t *testing.T) {
	g := NewWithT(t)
	var c counters
	_, changes, done := startLoop(t, &c)

	close(changes)

	g.Eventually(done).Should(BeClosed())
}

func TestWatchLoopWorksWithPlainReaders(t *testing.T) {
	g := NewWithT(t)
	captureOutput(t)
	var c counters
	changes := make(chan []string)
	finished := make(chan struct{})

	go func() {
		defer close(finished)
		watchLoop(strings.NewReader("hint\nlist\n"), changes, c.actions())
	}()

	g.Eventually(finished).Should(BeClosed(), "input ended")
	g.Expect([]int32{c.run.Load(), c.hint.Load(), c.list.Load()}).To(Equal([]int32{1, 1, 1}))
}

func TestIsTerminal(t *testing.T) {
	g := NewWithT(t)

	g.Expect(isTerminal(strings.NewReader("not a file"))).To(BeFalse())
	g.Expect(isTerminal(&bytes.Buffer{})).To(BeFalse())

	file, err := os.CreateTemp(t.TempDir(), "plain")
	g.Expect(err).NotTo(HaveOccurred())
	t.Cleanup(func() { _ = file.Close() })
	g.Expect(isTerminal(file)).To(BeFalse(), "a regular file is not a terminal")

	reader, writer, err := os.Pipe()
	g.Expect(err).NotTo(HaveOccurred())
	t.Cleanup(func() { _ = reader.Close(); _ = writer.Close() })
	g.Expect(isTerminal(reader)).To(BeFalse(), "a pipe is not a terminal")
}
