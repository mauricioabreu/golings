package watcher_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cleonte/golings/golings/watcher"
	. "github.com/onsi/gomega"
)

const quiet = 100 * time.Millisecond

// start watches a fresh folder holding exercises/a/main.go.
func start(t *testing.T) (root, file string, changes <-chan []string, stop context.CancelFunc) {
	t.Helper()
	root = filepath.Join(t.TempDir(), "exercises")
	dir := filepath.Join(root, "a")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	file = filepath.Join(dir, "main.go")
	write(t, file, "package main\n")

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	changes, err := watcher.Watch(ctx, root, quiet)
	if err != nil {
		t.Fatal(err)
	}
	return root, file, changes, cancel
}

func write(t *testing.T, name, content string) {
	t.Helper()
	if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func next(t *testing.T, changes <-chan []string) []string {
	t.Helper()
	select {
	case batch, ok := <-changes:
		if !ok {
			t.Fatal("the channel was closed")
		}
		return batch
	case <-time.After(5 * time.Second):
		t.Fatal("no change reported")
		return nil
	}
}

func expectQuiet(g *WithT, changes <-chan []string) {
	g.Consistently(changes, 5*quiet).ShouldNot(Receive())
}

func TestReportsAWrite(t *testing.T) {
	g := NewWithT(t)
	_, file, changes, _ := start(t)

	write(t, file, "package main // edited\n")

	g.Expect(next(t, changes)).To(Equal([]string{file}))
}

func TestOneSaveIsOneReport(t *testing.T) {
	g := NewWithT(t)
	_, file, changes, _ := start(t)

	for i := 0; i < 5; i++ {
		write(t, file, "package main // edit\n")
	}

	g.Expect(next(t, changes)).To(Equal([]string{file}))
	expectQuiet(g, changes)
}

func TestSeesEditorsThatSaveByRenaming(t *testing.T) {
	g := NewWithT(t)
	_, file, changes, _ := start(t)

	// vim and JetBrains write a temporary file and rename it over the original.
	temp := file + ".tmp"
	write(t, temp, "package main // atomic\n")
	if err := os.Rename(temp, file); err != nil {
		t.Fatal(err)
	}

	g.Expect(next(t, changes)).To(Equal([]string{file}))
}

func TestWatchesFoldersCreatedLater(t *testing.T) {
	g := NewWithT(t)
	root, _, changes, _ := start(t)

	newDir := filepath.Join(root, "b", "b1")
	if err := os.MkdirAll(newDir, 0o755); err != nil {
		t.Fatal(err)
	}
	// Give the watcher a moment to register the folders before writing in them.
	time.Sleep(3 * quiet)
	for len(changes) > 0 {
		<-changes
	}
	file := filepath.Join(newDir, "main.go")
	write(t, file, "package main\n")

	g.Eventually(func() []string {
		select {
		case batch := <-changes:
			return batch
		default:
			return nil
		}
	}, 5*time.Second, 20*time.Millisecond).Should(ContainElement(file))
}

func TestIgnoresFilesThatAreNotGoSources(t *testing.T) {
	g := NewWithT(t)
	_, file, changes, _ := start(t)
	dir := filepath.Dir(file)

	write(t, filepath.Join(dir, "notes.txt"), "x")
	write(t, filepath.Join(dir, ".main.go.swp"), "x")
	write(t, filepath.Join(dir, "main.go~"), "x")
	write(t, filepath.Join(dir, "#main.go#"), "x")

	expectQuiet(g, changes)
}

func TestClosesTheChannelWhenCancelled(t *testing.T) {
	_, _, changes, stop := start(t)

	stop()

	select {
	case _, ok := <-changes:
		if ok {
			t.Error("got a change after cancelling")
		}
	case <-time.After(5 * time.Second):
		t.Error("the channel was not closed")
	}
}

func TestFailsForAFolderThatDoesNotExist(t *testing.T) {
	g := NewWithT(t)

	_, err := watcher.Watch(context.Background(), filepath.Join(t.TempDir(), "nope"), quiet)

	g.Expect(err).To(HaveOccurred())
}
