package cmd_test

import (
	"bytes"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/cleonte/golings/golings/cmd"
	"github.com/cleonte/golings/golings/exercises"
	"github.com/fatih/color"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commands Suite")
}

var _ = Describe("Commands", func() {
	Describe("List", func() {
		Context("List exercises", func() {
			It("returns a list of exercises", func() {
				list := cmd.ListCmd("../fixtures/info.toml")

				err := list.Execute()

				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
	Describe("Run", func() {
		Context("Running 'compile' mode exercises", func() {
			When("it is compilable", func() {
				It("returns success", func() {
					run := cmd.RunCmd("../fixtures/success1/info.toml")
					run.SetArgs([]string{"success1"})

					err := run.Execute()

					Expect(err).ToNot(HaveOccurred())
				})
			})
			When("it is compilable but it is pending", func() {
				It("returns error", func() {
					run := cmd.RunCmd("../fixtures/pending1/info.toml")
					run.SetArgs([]string{"pending1"})

					err := run.Execute()

					Expect(err).To(HaveOccurred())
				})
			})
			When("it is not compilable", func() {
				It("returns error", func() {
					run := cmd.RunCmd("../fixtures/error1/info.toml")
					run.SetArgs([]string{"error1"})

					err := run.Execute()

					Expect(err).To(HaveOccurred())
				})
			})
			When("it does not exist", func() {
				It("returns an error", func() {
					run := cmd.RunCmd("../fixtures/info.toml")
					run.SetArgs([]string{"404"})

					err := run.Execute()

					Expect(err).To(HaveOccurred())
				})
			})
			When("'next' is passed is argument to the command", func() {
				It("runs the next pending exercise", func() {
					run := cmd.RunCmd("../fixtures/next/info.toml")
					run.SetArgs([]string{"next"})

					err := run.Execute()

					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
	Describe("Exercises that need a newer Go", func() {
		const infoFile = "../fixtures/goversion/cmd/info.toml"
		var out *bytes.Buffer

		BeforeEach(func() {
			out = &bytes.Buffer{}
			prevOutput, prevNoColor := color.Output, color.NoColor
			color.Output, color.NoColor = out, true
			DeferCleanup(func() { color.Output, color.NoColor = prevOutput, prevNoColor })
		})

		It("run refuses to run it and says why", func() {
			run := cmd.RunCmd(infoFile)
			run.SetArgs([]string{"future-broken"})

			err := run.Execute()

			Expect(err).To(MatchError(exercises.ErrGoVersionTooOld))
			Expect(out.String()).To(ContainSubstring("needs Go 99.0"))
			Expect(out.String()).NotTo(ContainSubstring("Failed to compile"))
		})
		It("verify skips it instead of running or passing it", func() {
			verify := cmd.VerifyCmd(infoFile)

			err := verify.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(ContainSubstring("Skipped"))
			Expect(out.String()).To(ContainSubstring("You passed all the exercises"))
		})
		It("watch reports how many exercises were skipped", func() {
			cmd.RunNextExercise(infoFile)

			Expect(out.String()).To(ContainSubstring("1 exercise(s) skipped"))
		})
	})
})
