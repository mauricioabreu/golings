package exercises_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mauricioabreu/golings/src/exercises"
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
				file, err := ioutil.TempFile("/tmp", "exercise*.go")
				file.Write([]byte(`// exercise1.go
				// I AM NOT DONE
				package main

				func main() {

				}
				`))
				Expect(err).To(BeNil())

				defer os.Remove(file.Name())

				ex := exercises.Exercise{Path: file.Name()}

				Expect(ex.State()).To(Equal(exercises.Pending))
			})
		})

		When("'I AM NOT DONE' comment is not there", func() {
			It("has the Done state", func() {
				file, err := ioutil.TempFile("/tmp", "exercise*.go")
				Expect(err).To(BeNil())

				defer os.Remove(file.Name())

				ex := exercises.Exercise{Path: file.Name()}

				Expect(ex.State()).To(Equal(exercises.Done))
			})
		})
	})
})
