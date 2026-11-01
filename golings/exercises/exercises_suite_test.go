package exercises_test

import (
	"os"
	"testing"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestExercises(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Exercises Suite")
}

var _ = Describe("Exercises", func() {
	Describe("Checking exercise state", func() {
		When("'I AM NOT DONE' comment is still there", func() {
			It("has the Pending state", func() {
				file, err := os.CreateTemp("/tmp", "exercise*.go")
				Expect(err).NotTo(HaveOccurred())
				_, err = file.Write([]byte(`// exercise1.go
				// I AM NOT DONE
				package main

				func main() {

				}
				`))
				Expect(err).NotTo(HaveOccurred())

				defer func() { _ = os.Remove(file.Name()) }()

				ex := exercises.Exercise{Path: file.Name()}

				Expect(ex.State()).To(Equal(exercises.Pending))
				Expect(ex.State().String()).To(Equal("Pending"))
			})
		})

		When("'I AM NOT DONE' comment is not there", func() {
			It("has the Done state", func() {
				file, err := os.CreateTemp("/tmp", "exercise*.go")
				Expect(err).NotTo(HaveOccurred())

				defer func() { _ = os.Remove(file.Name()) }()

				ex := exercises.Exercise{Path: file.Name()}

				Expect(ex.State()).To(Equal(exercises.Done))
				Expect(ex.State().String()).To(Equal("Done"))
			})
		})
	})
	Describe("Listing exercises", func() {
		When("info file exists", func() {
			It("returns a list with the exercises in it", func() {
				list, err := exercises.List("../fixtures/info.toml")

				Expect(err).NotTo(HaveOccurred())
				Expect(len(list)).To(Equal(2))
			})
		})
		When("info file does not exist", func() {
			It("returns an error", func() {
				list, err := exercises.List("../fixtures/info404.toml")

				Expect(err).To(HaveOccurred())
				Expect(len(list)).To(Equal(0))
			})
		})
	})
	Describe("Find an exercise", func() {
		When("exercise exists in the info file", func() {
			It("returns info about the exercise", func() {
				exercise, err := exercises.Find("compile1", "../fixtures/info.toml")

				Expect(err).NotTo(HaveOccurred())
				Expect(exercise.Name).To(Equal("compile1"))
				Expect(exercise.Path).To(Equal("compile/compile1.go"))
				Expect(exercise.Mode).To(Equal("compile"))
				Expect(exercise.Hint).To(Equal("hints are cool"))
			})
		})
		When("exercise does not exist in the info file", func() {
			It("returns an error", func() {
				_, err := exercises.Find("compile404", "../fixtures/info.toml")

				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(exercises.ErrExerciseNotFound))
			})
		})
	})
	Describe("Reporting progress", func() {
		When("half exercises pending", func() {
			It("reports 50%% progress", func() {
				progress, done, total, err := exercises.Progress("../fixtures/progress/info.toml")

				Expect(err).NotTo(HaveOccurred())
				Expect(done).To(Equal(1))
				Expect(total).To(Equal(2))
				Expect(progress).To(Equal(float32(0.5)))
			})
		})
	})
	Describe("Go version requirements", func() {
		It("reads go_version from the info file", func() {
			list, err := exercises.List("../fixtures/goversion/info.toml")

			Expect(err).NotTo(HaveOccurred())
			Expect(list[0].GoVersion).To(BeEmpty())
			Expect(list[1].GoVersion).To(Equal("99.0"))
		})
		DescribeTable("Supports",
			func(min, local string, want bool) {
				Expect(exercises.Exercise{GoVersion: min}.Supports(local)).To(Equal(want))
			},
			Entry("no minimum", "", "go1.20", true),
			Entry("same release", "1.26", "go1.26.0", true),
			Entry("patch release of minimum", "1.25", "go1.25.3", true),
			Entry("newer than minimum", "1.25", "go1.26.1", true),
			Entry("older than minimum", "1.26", "go1.25.9", false),
		)
		It("lists exercises the local Go can't run as skipped", func() {
			skipped, err := exercises.Skipped("../fixtures/goversion/pending/info.toml")

			Expect(err).NotTo(HaveOccurred())
			Expect(skipped).To(HaveLen(1))
			Expect(skipped[0].Name).To(Equal("future1"))
		})
		It("does not offer a skipped exercise as the next pending one", func() {
			_, err := exercises.NextPending("../fixtures/goversion/pending/info.toml")

			Expect(err).To(MatchError(exercises.ErrNoPendingExercises))
		})
		It("leaves skipped exercises out of the progress total", func() {
			progress, done, total, err := exercises.Progress("../fixtures/goversion/pending/info.toml")

			Expect(err).NotTo(HaveOccurred())
			Expect(done).To(Equal(1))
			Expect(total).To(Equal(1))
			Expect(progress).To(Equal(float32(1)))
		})
		It("does not run an exercise that needs a newer Go", func() {
			list, err := exercises.List("../fixtures/goversion/info.toml")
			Expect(err).NotTo(HaveOccurred())

			_, err = list[1].Run()

			Expect(err).To(MatchError(exercises.ErrGoVersionTooOld))
		})
	})
})
