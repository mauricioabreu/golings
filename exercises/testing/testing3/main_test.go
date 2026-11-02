// testing3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// createFile makes a file with the content for the test and returns its path.
// When the test is over, the file has to be gone.
func createFile(t *testing.T, content string) string {
	// TODO: A folder made with os.MkdirTemp stays on the disk. Let the testing package make and remove the folder.
	dir, err := os.MkdirTemp("", "golings-testing3-")
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(dir, "data.txt")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestCreateFileCleansUpAfterTheTest(t *testing.T) {
	var path string

	t.Run("uses the file", func(t *testing.T) {
		path = createFile(t, "hello")

		data, err := os.ReadFile(path)
		if err != nil || string(data) != "hello" {
			t.Errorf("read %q, %v; want hello", data, err)
		}
	})

	if _, err := os.Stat(path); !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("%s is still there after the test that used it", path)
		_ = os.RemoveAll(filepath.Dir(path)) // don't leave it behind while you work on this
	}
}
