package workspace_test

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/cleonte/golings"
	"github.com/cleonte/golings/golings/doctor"
	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/workspace"
	. "github.com/onsi/gomega"
)

const unsolved = "// I AM NOT DONE\npackage main\n\nfunc main() {}\n"

func sampleFS() fstest.MapFS {
	return fstest.MapFS{
		"info.toml": {Data: []byte(`[[exercises]]
name = "a1"
path = "exercises/a/a1/main.go"
mode = "compile"
hint = ""
`)},
		"exercises/a/a1/main.go": {Data: []byte(unsolved)},
		"exercises/a/README.md":  {Data: []byte("# a")},
	}
}

func TestInitWritesExercisesAndGoMod(t *testing.T) {
	g := NewWithT(t)
	dest := filepath.Join(t.TempDir(), "ws")

	g.Expect(workspace.Init(sampleFS(), dest)).To(Succeed())

	read := func(name string) string {
		data, err := os.ReadFile(filepath.Join(dest, filepath.FromSlash(name)))
		g.Expect(err).NotTo(HaveOccurred(), name)
		return string(data)
	}
	g.Expect(read("exercises/a/a1/main.go")).To(Equal(unsolved))
	g.Expect(read("exercises/a/README.md")).To(Equal("# a"))
	g.Expect(read("info.toml")).To(ContainSubstring(`name = "a1"`))
	g.Expect(read("go.mod")).To(Equal("module golings-exercises\n\ngo " + doctor.MinGoVersion + "\n"))
}

func TestInitRefusesExistingPathAndLeavesItAlone(t *testing.T) {
	g := NewWithT(t)
	dest := t.TempDir()
	keep := filepath.Join(dest, "mine.txt")
	g.Expect(os.WriteFile(keep, []byte("precious"), 0o644)).To(Succeed())

	err := workspace.Init(sampleFS(), dest)

	g.Expect(err).To(MatchError(workspace.ErrExists))
	data, readErr := os.ReadFile(keep)
	g.Expect(readErr).NotTo(HaveOccurred())
	g.Expect(string(data)).To(Equal("precious"))
	g.Expect(filepath.Join(dest, "info.toml")).NotTo(BeAnExistingFile())
}

// failingFS fails to read one file, to simulate an error half way through.
type failingFS struct {
	fs.FS
	bad string
}

func (f failingFS) Open(name string) (fs.File, error) {
	if name == f.bad {
		return nil, errors.New("disk on fire")
	}
	return f.FS.Open(name)
}

func TestInitCleansUpAfterAFailure(t *testing.T) {
	g := NewWithT(t)
	dest := filepath.Join(t.TempDir(), "ws")

	err := workspace.Init(failingFS{FS: sampleFS(), bad: "exercises/a/a1/main.go"}, dest)

	g.Expect(err).To(MatchError(ContainSubstring("disk on fire")))
	g.Expect(dest).NotTo(BeADirectory())
}

func writeLearnerWork(g *WithT, root string) string {
	file := filepath.Join(root, "exercises", "a", "a1", "main.go")
	g.Expect(os.MkdirAll(filepath.Dir(file), 0o755)).To(Succeed())
	g.Expect(os.WriteFile(file, []byte("package main\n\nfunc main() { println(1) }\n"), 0o644)).To(Succeed())
	return file
}

func TestResetRestoresTheOriginalIncludingTheMarker(t *testing.T) {
	g := NewWithT(t)
	root := t.TempDir()
	file := writeLearnerWork(g, root)

	path, err := workspace.Reset(sampleFS(), root, "a1")

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(path).To(Equal("exercises/a/a1/main.go"))
	data, readErr := os.ReadFile(file)
	g.Expect(readErr).NotTo(HaveOccurred())
	g.Expect(string(data)).To(Equal(unsolved))
	g.Expect(exercises.Exercise{Path: file}.State()).To(Equal(exercises.Pending))
}

func TestResetRecreatesADeletedExercise(t *testing.T) {
	g := NewWithT(t)
	root := t.TempDir()

	_, err := workspace.Reset(sampleFS(), root, "a1")

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(filepath.Join(root, "exercises", "a", "a1", "main.go")).To(BeAnExistingFile())
}

func TestResetUnknownExercise(t *testing.T) {
	g := NewWithT(t)
	root := t.TempDir()
	file := writeLearnerWork(g, root)

	for _, name := range []string{"nope", "", "../a1", "exercises/a/a1/main.go"} {
		_, err := workspace.Reset(sampleFS(), root, name)
		g.Expect(err).To(MatchError(exercises.ErrExerciseNotFound), name)
	}
	data, _ := os.ReadFile(file)
	g.Expect(string(data)).To(ContainSubstring("println(1)"), "learner's work must be untouched")
}

func TestResetRejectsManifestPathsOutsideTheWorkspace(t *testing.T) {
	g := NewWithT(t)
	root := t.TempDir()
	src := fstest.MapFS{
		"info.toml": {Data: []byte(`[[exercises]]
name = "evil"
path = "../outside.go"
mode = "compile"
hint = ""
`)},
	}

	_, err := workspace.Reset(src, root, "evil")

	g.Expect(err).To(MatchError(ContainSubstring("invalid path")))
	g.Expect(filepath.Join(filepath.Dir(root), "outside.go")).NotTo(BeAnExistingFile())
}

func TestEmbeddedExercisesAreComplete(t *testing.T) {
	g := NewWithT(t)

	embedded, err := fs.ReadFile(golings.FS, "info.toml")
	g.Expect(err).NotTo(HaveOccurred())
	listed, err := exercises.ParseInfo(embedded)
	g.Expect(err).NotTo(HaveOccurred())
	onDisk, err := exercises.List("../../info.toml")
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(listed).To(Equal(onDisk))
	for _, ex := range listed {
		original, err := fs.ReadFile(golings.FS, ex.Path)
		g.Expect(err).NotTo(HaveOccurred(), ex.Name)
		g.Expect(string(original)).To(ContainSubstring("I AM NOT DONE"), ex.Name)
	}
}

func TestInitWithTheRealExercises(t *testing.T) {
	g := NewWithT(t)
	dest := filepath.Join(t.TempDir(), "ws")

	g.Expect(workspace.Init(golings.FS, dest)).To(Succeed())

	listed, err := exercises.List(filepath.Join(dest, "info.toml"))
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(listed).NotTo(BeEmpty())
	for _, ex := range listed {
		g.Expect(filepath.Join(dest, filepath.FromSlash(ex.Path))).To(BeAnExistingFile(), ex.Name)
	}
}
