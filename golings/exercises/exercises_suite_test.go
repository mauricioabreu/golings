package exercises_test

import (
	"os"
	"testing"

	"github.com/mauricioabreu/golings/golings/exercises"
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

				defer os.Remove(file.Name())

				ex := exercises.Exercise{Path: file.Name()}

				Expect(ex.State()).To(Equal(exercises.Pending))
				Expect(ex.State().String()).To(Equal("Pending"))
			})
		})

		When("'I AM NOT DONE' comment is not there", func() {
			It("has the Done state", func() {
				file, err := os.CreateTemp("/tmp", "exercise*.go")
				Expect(err).NotTo(HaveOccurred())

				defer os.Remove(file.Name())

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
})
