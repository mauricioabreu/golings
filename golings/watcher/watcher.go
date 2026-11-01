// Package watcher reports when exercise files change.
package watcher

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Watch reports changes to Go files below root until ctx is done, then closes
// the returned channel. Each value is the sorted list of files that changed
// since the previous one: a burst of events is held back until nothing has
// happened for the quiet period, so one save produces one value.
//
// Folders created later are watched too. Creating or renaming a file counts as
// a change as well as writing it, because many editors (vim, JetBrains) save by
// writing a temporary file and renaming it over the original.
func Watch(ctx context.Context, root string, quiet time.Duration) (<-chan []string, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	if err := addTree(w, root); err != nil {
		_ = w.Close()
		return nil, err
	}

	changes := make(chan []string)
	go func() {
		defer close(changes)
		defer w.Close() // nolint

		pending := map[string]bool{}
		// The timer only runs while there are pending changes.
		timer := time.NewTimer(quiet)
		timer.Stop()
		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case event, ok := <-w.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) {
					if info, err := os.Stat(event.Name); err == nil && info.IsDir() {
						_ = addTree(w, event.Name)
						continue
					}
				}
				if !isExerciseFile(event.Name) || !event.Has(fsnotify.Write|fsnotify.Create|fsnotify.Rename) {
					continue
				}
				pending[event.Name] = true
				timer.Reset(quiet)

			case _, ok := <-w.Errors:
				// Nothing useful to do about a failed event; keep watching.
				if !ok {
					return
				}

			case <-timer.C:
				batch := make([]string, 0, len(pending))
				for name := range pending {
					batch = append(batch, name)
				}
				slices.Sort(batch)
				clear(pending)

				select {
				case changes <- batch:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return changes, nil
}

func addTree(w *fsnotify.Watcher, root string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return w.Add(path)
		}
		return nil
	})
}

// isExerciseFile is true for Go sources, leaving out the backups, swap files
// and temporary files editors create next to them.
func isExerciseFile(name string) bool {
	base := filepath.Base(name)
	return strings.HasSuffix(base, ".go") && !strings.HasPrefix(base, ".") && !strings.HasPrefix(base, "#")
}
