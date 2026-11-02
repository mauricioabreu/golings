package exercises_test

import (
	"os"
	"path/filepath"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Marking an exercise as done", func() {
	write := func(content string, mode os.FileMode) exercises.Exercise {
		file := filepath.Join(GinkgoT().TempDir(), "main.go")
		Expect(os.WriteFile(file, []byte(content), mode)).To(Succeed())
		return exercises.Exercise{Path: file}
	}

	It("removes the marker and leaves the rest of the file alone", func() {
		ex := write("// x1\n// Make me compile!\n\n// I AM NOT DONE\npackage main\n\nfunc main() {}\n", 0o644)

		Expect(ex.MarkDone()).To(Succeed())

		data, err := os.ReadFile(ex.Path)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(data)).To(Equal("// x1\n// Make me compile!\n\npackage main\n\nfunc main() {}\n"))
		Expect(ex.State()).To(Equal(exercises.Done))
	})
	It("understands the variants of the marker the state check accepts", func() {
		for _, marker := range []string{"//I AM NOT DONE", "  // i am not done", "/// I  AM NOT DONE"} {
			ex := write("package main\n"+marker+"\nfunc main() {}\n", 0o644)
			if ex.State() == exercises.Done {
				continue // not a marker the tool recognizes, nothing to remove
			}

			Expect(ex.MarkDone()).To(Succeed(), marker)
			Expect(ex.State()).To(Equal(exercises.Done), marker)
		}
	})
	It("tells when there is no marker, and does not touch the file", func() {
		ex := write("package main\n\nfunc main() {}\n", 0o644)

		Expect(ex.MarkDone()).To(MatchError(exercises.ErrNotMarked))

		data, _ := os.ReadFile(ex.Path)
		Expect(string(data)).To(Equal("package main\n\nfunc main() {}\n"))
	})
	It("fails for a file that does not exist", func() {
		Expect(exercises.Exercise{Path: "does/not/exist.go"}.MarkDone()).To(HaveOccurred())
	})
	It("keeps the permissions of the file", func() {
		ex := write("// I AM NOT DONE\npackage main\n", 0o600)

		Expect(ex.MarkDone()).To(Succeed())

		info, err := os.Stat(ex.Path)
		Expect(err).NotTo(HaveOccurred())
		Expect(info.Mode().Perm()).To(Equal(os.FileMode(0o600)))
	})
	It("works on a file whose last line is the marker, without a final newline", func() {
		ex := write("package main\n// I AM NOT DONE", 0o644)

		Expect(ex.MarkDone()).To(Succeed())

		data, _ := os.ReadFile(ex.Path)
		Expect(string(data)).To(Equal("package main\n"))
	})
})
