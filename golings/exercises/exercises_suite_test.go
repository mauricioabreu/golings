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
				file.Write([]byte(`// exercise1.go
				// I AM NOT DONE
				package main

				func main() {

				}
				`))
				Expect(err).NotTo(HaveOccurred())

				defer os.Remove(file.Name())

				ex := exercises.Exercise{Path: file.Name()}

				Expect(ex.State()).To(Equal(exercises.Pending))
			})
		})

		When("'I AM NOT DONE' comment is not there", func() {
			It("has the Done state", func() {
				file, err := os.CreateTemp("/tmp", "exercise*.go")
				Expect(err).NotTo(HaveOccurred())

				defer os.Remove(file.Name())

				ex := exercises.Exercise{Path: file.Name()}

				Expect(ex.State()).To(Equal(exercises.Done))
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
	})
})
