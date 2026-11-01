package exercises_test

import (
	"os"
	"path/filepath"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// writeExercise creates a throwaway exercise and returns it with a path
// relative to the working directory, which is how the runner is invoked.
func writeExercise(mode, name, source string) exercises.Exercise {
	dir := GinkgoT().TempDir()
	file := filepath.Join(dir, name)
	Expect(os.WriteFile(file, []byte(source), 0o644)).To(Succeed())

	wd, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())
	rel, err := filepath.Rel(wd, file)
	Expect(err).NotTo(HaveOccurred())

	return exercises.Exercise{Name: "tmp", Path: rel, Mode: mode}
}

var _ = Describe("Running exercises", func() {
	Describe("Run", func() {
		It("captures the output of a compile-mode exercise that works", func() {
			ex := writeExercise("compile", "main.go", `package main

import "fmt"

func main() { fmt.Println("hello golings") }
`)

			result, err := ex.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Out).To(ContainSubstring("hello golings"))
			Expect(result.Exercise).To(Equal(ex))
		})
		It("reports the compiler error of a compile-mode exercise that is broken", func() {
			ex := writeExercise("compile", "main.go", `package main

func main() { a := 1 }
`)

			result, err := ex.Run()

			Expect(err).To(HaveOccurred())
			Expect(result.BuildFailed).To(BeTrue())
			Expect(result.Err).To(ContainSubstring("declared and not used"))
		})
		It("reports a test-mode exercise that does not compile as a build failure", func() {
			ex := writeExercise("test", "x_test.go", `package x

import "testing"

func TestBroken(t *testing.T) { undefinedThing() }
`)

			result, err := ex.Run()

			Expect(err).To(HaveOccurred())
			Expect(result.BuildFailed).To(BeTrue())
			Expect(result.Failure().Kind).To(Equal(exercises.CompileFailure))
		})
		It("reports a program that crashes while running as a runtime failure", func() {
			ex := writeExercise("compile", "main.go", `package main

func main() { panic("boom") }
`)

			result, err := ex.Run()

			Expect(err).To(HaveOccurred())
			Expect(result.BuildFailed).To(BeFalse())
			Expect(result.Failure().Kind).To(Equal(exercises.RuntimeFailure))
			Expect(result.Err).To(ContainSubstring("boom"))
		})
		It("passes a test-mode exercise whose tests pass", func() {
			ex := writeExercise("test", "x_test.go", `package x

import "testing"

func TestOK(t *testing.T) {}
`)

			result, err := ex.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Out).To(ContainSubstring("PASS"))
		})
		It("fails a test-mode exercise whose tests fail", func() {
			ex := writeExercise("test", "x_test.go", `package x

import "testing"

func TestBroken(t *testing.T) { t.Fatal("not yet") }
`)

			result, err := ex.Run()

			Expect(err).To(HaveOccurred())
			Expect(result.BuildFailed).To(BeFalse(), "it compiled, its tests failed")
			Expect(result.Out).To(ContainSubstring("not yet"))
		})
	})
})
