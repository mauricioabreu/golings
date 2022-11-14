package cmd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mauricioabreu/golings/golings/cmd"
)

func TestList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "List Suite")
}

var _ = Describe("List", func() {
	Context("List exercises", func() {
		It("returns a list of exercises", func() {
			list := cmd.ListCmd("../fixtures/info.toml")

			err := list.Execute()

			Expect(err).To(BeNil())
		})
	})
})
