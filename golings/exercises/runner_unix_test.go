//go:build unix

package exercises_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/gomega"
)

// stuckExercise writes its process id to pidFile, prints "started" and then
// never finishes (a main loop, or a test that sleeps).
func stuckExercise(t *testing.T, mode string) (exercises.Exercise, string) {
	t.Helper()
	dir := t.TempDir()
	pidFile := filepath.Join(dir, "pid")

	var name, source string
	if mode == "compile" {
		name, source = "main.go", fmt.Sprintf(`package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	os.WriteFile(%q, []byte(strconv.Itoa(os.Getpid())), 0o644)
	fmt.Println("started")
	for {
	}
}
`, pidFile)
	} else {
		name, source = "x_test.go", fmt.Sprintf(`package x

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestStuck(t *testing.T) {
	os.WriteFile(%q, []byte(strconv.Itoa(os.Getpid())), 0o644)
	fmt.Println("started")
	time.Sleep(time.Hour)
}
`, pidFile)
	}

	file := filepath.Join(dir, name)
	if err := os.WriteFile(file, []byte(source), 0o644); err != nil {
		t.Fatal(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	rel, err := filepath.Rel(wd, file)
	if err != nil {
		t.Fatal(err)
	}
	return exercises.Exercise{Name: "stuck", Path: rel, Mode: mode}, pidFile
}

func processGone(t *testing.T, pidFile string) bool {
	t.Helper()
	data, err := os.ReadFile(pidFile)
	if err != nil {
		t.Fatalf("the exercise never wrote its pid: %v", err)
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	return errors.Is(syscall.Kill(pid, 0), syscall.ESRCH)
}

func TestStuckExerciseIsStopped(t *testing.T) {
	for _, mode := range []string{"compile", "test"} {
		t.Run(mode, func(t *testing.T) {
			g := NewWithT(t)
			ex, pidFile := stuckExercise(t, mode)

			started := time.Now()
			result, err := ex.RunContext(context.Background(), 2*time.Second)
			elapsed := time.Since(started)

			g.Expect(err).To(MatchError(exercises.ErrTimeout))
			g.Expect(result.TimedOut).To(BeTrue())
			g.Expect(result.Out).To(ContainSubstring("started"), "output printed before the timeout is kept")
			g.Expect(result.Failure().Kind).To(Equal(exercises.TimeoutFailure))
			g.Expect(elapsed).To(BeNumerically("<", 20*time.Second), "build plus a 2s timeout")
			g.Expect(processGone(t, pidFile)).To(BeTrue(), "the program must not keep running after the timeout")
		})
	}
}

func TestCancellingStopsARunningExercise(t *testing.T) {
	g := NewWithT(t)
	ex, pidFile := stuckExercise(t, "compile")
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel once the program is up, the way a new save would.
	go func() {
		for {
			if _, err := os.Stat(pidFile); err == nil {
				cancel()
				return
			}
			time.Sleep(20 * time.Millisecond)
		}
	}()

	result, err := ex.RunContext(ctx, time.Minute)

	g.Expect(err).To(MatchError(context.Canceled))
	g.Expect(result.TimedOut).To(BeFalse(), "cancelled is not the same as timed out")
	g.Expect(processGone(t, pidFile)).To(BeTrue())
}
