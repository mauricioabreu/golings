package cmd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mauricioabreu/golings/golings/cmd"
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
})
