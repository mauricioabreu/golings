package exercises_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cleonte/golings/golings/exercises"
)

// allSolutionsRequired is switched on once every exercise has a solution.
// From then on a new exercise without one fails the build.
const allSolutionsRequired = false

const repoRoot = "../.."

// TestSolutions proves each exercise can be solved: the unsolved version has
// to fail and the solution has to pass, through the same runner learners use.
func TestSolutions(t *testing.T) {
	if testing.Short() {
		t.Skip("runs every exercise with go run / go test; skipped in -short mode")
	}

	all, err := exercises.List(filepath.Join(repoRoot, "info.toml"))
	if err != nil {
		t.Fatal(err)
	}

	for _, ex := range all {
		solutionPath := ex.SolutionPath()
		if solutionPath == "" {
			t.Errorf("%s: path %q is not laid out as exercises/<topic>/<name>/...", ex.Name, ex.Path)
			continue
		}
		solutionFile := filepath.Join(repoRoot, filepath.FromSlash(solutionPath))

		source, err := os.ReadFile(solutionFile)
		if os.IsNotExist(err) {
			if allSolutionsRequired {
				t.Errorf("%s: missing solution %s", ex.Name, solutionPath)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: %v", ex.Name, err)
			continue
		}

		t.Run(ex.Name, func(t *testing.T) {
			t.Parallel()

			if strings.Contains(string(source), "I AM NOT DONE") {
				t.Errorf("%s still has the 'I AM NOT DONE' marker", solutionPath)
			}

			unsolved := ex
			unsolved.Path = filepath.Join(repoRoot, filepath.FromSlash(ex.Path))
			if result, err := unsolved.Run(); err == nil {
				t.Errorf("the unsolved exercise passes, it should fail:\n%s", result.Out)
			}

			solved := ex
			solved.Path = solutionFile
			if result, err := solved.Run(); err != nil {
				t.Errorf("the solution fails: %v\n%s%s", err, result.Err, result.Out)
			}
		})
	}
}

// TestNoOrphanedSolutions catches solutions left behind after an exercise was
// renamed or removed, and stray files in the solutions tree.
func TestNoOrphanedSolutions(t *testing.T) {
	all, err := exercises.List(filepath.Join(repoRoot, "info.toml"))
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]bool{}
	for _, ex := range all {
		expected[ex.SolutionPath()] = true
	}

	root := filepath.Join(repoRoot, "solutions")
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(repoRoot, path)
		if err != nil {
			return err
		}
		if rel = filepath.ToSlash(rel); !expected[rel] {
			t.Errorf("%s does not belong to any exercise in info.toml", rel)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
