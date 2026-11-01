package exercises_test

import (
	"os"
	"path/filepath"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Listing edge cases", func() {
	It("returns the first pending exercise", func() {
		ex, err := exercises.NextPending("../fixtures/next/info.toml")

		Expect(err).NotTo(HaveOccurred())
		Expect(ex.Path).To(Equal("../fixtures/next/next2/main.go"))
	})
	It("reports when no exercise is pending", func() {
		_, err := exercises.NextPending("../fixtures/success1/info.toml")

		Expect(err).To(MatchError(exercises.ErrNoPendingExercises))
	})
	It("returns an error when the info file can't be read", func() {
		_, err := exercises.NextPending("../fixtures/info404.toml")

		Expect(err).To(HaveOccurred())
	})
	It("returns an error when the info file is not valid TOML", func() {
		file := filepath.Join(GinkgoT().TempDir(), "info.toml")
		Expect(os.WriteFile(file, []byte("[[exercises\nname = "), 0o644)).To(Succeed())

		_, err := exercises.List(file)

		Expect(err).To(HaveOccurred())
	})
	It("treats a missing exercise file as pending", func() {
		Expect(exercises.Exercise{Path: "does/not/exist.go"}.State()).To(Equal(exercises.Pending))
	})
})
