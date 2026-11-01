package cmd_test

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/cleonte/golings/golings/cmd"
	"github.com/fatih/color"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// captureColor redirects the colored output of the commands into the returned buffer.
func captureColor() *bytes.Buffer {
	out := &bytes.Buffer{}
	prevOutput, prevNoColor := color.Output, color.NoColor
	color.Output, color.NoColor = out, true
	DeferCleanup(func() { color.Output, color.NoColor = prevOutput, prevNoColor })
	return out
}

// writeBrokenPendingInfo creates a pending exercise that does not compile and
// returns the path of an info file listing it.
func writeBrokenPendingInfo() string {
	dir := GinkgoT().TempDir()
	Expect(os.WriteFile(filepath.Join(dir, "main.go"), []byte(`// I AM NOT DONE
package main

func main() { a := 1 }
`), 0o644)).To(Succeed())

	wd, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())
	rel, err := filepath.Rel(wd, filepath.Join(dir, "main.go"))
	Expect(err).NotTo(HaveOccurred())

	info := filepath.Join(dir, "info.toml")
	Expect(os.WriteFile(info, []byte(`[[exercises]]
name = "broken"
path = "`+rel+`"
mode = "compile"
hint = "remove the unused variable"
`), 0o644)).To(Succeed())
	return info
}

var _ = Describe("More commands", func() {
	var out *bytes.Buffer

	BeforeEach(func() {
		out = captureColor()
	})

	Describe("Hint", func() {
		It("prints the hint of a named exercise", func() {
			hint := cmd.HintCmd("../fixtures/info.toml")
			hint.SetArgs([]string{"compile1"})

			Expect(hint.Execute()).To(Succeed())
			Expect(out.String()).To(ContainSubstring("hints are cool"))
		})
		It("prints the hint of the next pending exercise", func() {
			hint := cmd.HintCmd(writeBrokenPendingInfo())
			hint.SetArgs([]string{"next"})

			Expect(hint.Execute()).To(Succeed())
			Expect(out.String()).To(ContainSubstring("remove the unused variable"))
		})
		It("fails for an unknown exercise", func() {
			hint := cmd.HintCmd("../fixtures/info.toml")
			hint.SetArgs([]string{"404"})

			Expect(hint.Execute()).To(HaveOccurred())
			Expect(out.String()).To(ContainSubstring("exercise not found"))
		})
		It("fails when no exercise is pending", func() {
			hint := cmd.HintCmd("../fixtures/success1/info.toml")
			hint.SetArgs([]string{"next"})

			Expect(hint.Execute()).To(HaveOccurred())
		})
	})

	Describe("List", func() {
		It("fails when the info file does not exist", func() {
			list := cmd.ListCmd("../fixtures/info404.toml")

			Expect(list.Execute()).To(HaveOccurred())
		})
	})

	Describe("Verify", func() {
		It("passes when every exercise compiles", func() {
			verify := cmd.VerifyCmd("../fixtures/success1/info.toml")

			Expect(verify.Execute()).To(Succeed())
			Expect(out.String()).To(ContainSubstring("You passed all the exercises"))
		})
		It("fails on the first exercise that does not compile", func() {
			verify := cmd.VerifyCmd("../fixtures/error1/info.toml")

			Expect(verify.Execute()).To(HaveOccurred())
			Expect(out.String()).To(ContainSubstring("Failed to compile the exercise"))
			Expect(out.String()).NotTo(ContainSubstring("You passed all the exercises"))
		})
		It("fails when the info file does not exist", func() {
			verify := cmd.VerifyCmd("../fixtures/info404.toml")

			Expect(verify.Execute()).To(HaveOccurred())
		})
	})

	Describe("RunNextExercise", func() {
		It("congratulates the learner when everything is done", func() {
			cmd.RunNextExercise("../fixtures/success1/info.toml")

			Expect(out.String()).To(ContainSubstring("Progress: 1/1 (100.00%)"))
			Expect(out.String()).To(ContainSubstring("completed all 1 of the currently available exercises"))
		})
		It("explains that a compiling exercise is still pending", func() {
			cmd.RunNextExercise("../fixtures/pending1/info.toml")

			Expect(out.String()).To(ContainSubstring("Progress: 0/1 (0.00%)"))
			Expect(out.String()).To(ContainSubstring("Remove the 'I AM NOT DONE'"))
		})
		It("shows the compiler error and points to the hint", func() {
			cmd.RunNextExercise(writeBrokenPendingInfo())

			Expect(out.String()).To(ContainSubstring("Failed to compile the exercise"))
			Expect(out.String()).To(ContainSubstring("declared and not used"))
			Expect(out.String()).To(ContainSubstring("golings hint broken"))
		})
	})

	Describe("PrintHint", func() {
		It("prints the hint of the next pending exercise", func() {
			cmd.PrintHint(writeBrokenPendingInfo())

			Expect(out.String()).To(ContainSubstring("remove the unused variable"))
		})
		It("reports when there is no next exercise", func() {
			cmd.PrintHint("../fixtures/success1/info.toml")

			Expect(out.String()).To(ContainSubstring("Failed to find next exercises"))
		})
	})

	Describe("Root", func() {
		It("registers every command", func() {
			root := cmd.NewRootCmd("test")

			names := []string{}
			for _, c := range root.Commands() {
				names = append(names, c.Name())
			}
			Expect(names).To(ContainElements("hint", "list", "run", "verify", "watch"))
			Expect(root.Version).To(Equal("test"))
		})
	})
})
