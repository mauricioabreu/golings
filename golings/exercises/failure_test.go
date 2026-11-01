package exercises_test

import (
	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Interpreting failures", func() {
	compile := exercises.Exercise{Name: "e", Mode: "compile"}
	test := exercises.Exercise{Name: "e", Mode: "test"}

	It("extracts compiler diagnostics and drops the package header", func() {
		result := exercises.Result{Exercise: compile, Err: "# command-line-arguments\nexercises/a/main.go:10:6: syntax error: unexpected =, expected name\nexercises/a/main.go:12:2: undefined: x\n"}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.CompileFailure))
		Expect(f.Diagnostics).To(Equal([]exercises.Diagnostic{
			{File: "exercises/a/main.go", Line: 10, Col: 6, Message: "syntax error: unexpected =, expected name"},
			{File: "exercises/a/main.go", Line: 12, Col: 2, Message: "undefined: x"},
		}))
		Expect(f.Output).To(BeEmpty())
	})
	It("keeps messages that are not diagnostics", func() {
		result := exercises.Result{Exercise: compile, Err: "# command-line-arguments\na/main.go:3:1: undefined: y\na/main.go:3:1: too many errors\n"}

		Expect(result.Failure().Diagnostics).To(HaveLen(2))

		result = exercises.Result{Exercise: compile, Err: "a/main.go:3:1: undefined: y\ntoo many errors\n"}

		f := result.Failure()
		Expect(f.Diagnostics).To(HaveLen(1))
		Expect(f.Output).To(Equal("too many errors"))
	})
	It("understands diagnostics without a column", func() {
		result := exercises.Result{Exercise: compile, Err: "vet: a/main.go:7: oops\n"}

		Expect(result.Failure().Diagnostics).To(ConsistOf(
			exercises.Diagnostic{File: "a/main.go", Line: 7, Message: "oops"},
		))
	})
	It("reports a compile failure of a test-mode exercise as such", func() {
		result := exercises.Result{Exercise: test, Out: "# command-line-arguments\na/main_test.go:5:2: undefined: Foo\nFAIL\tcommand-line-arguments [build failed]\n"}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.CompileFailure))
		Expect(f.Diagnostics).To(HaveLen(1))
	})
	It("does not mistake t.Log lines for compiler errors", func() {
		result := exercises.Result{Exercise: test, Out: "=== RUN   TestX\n    main_test.go:17: 2 is bigger than 1\n--- FAIL: TestX (0.00s)\nFAIL\nFAIL\tcommand-line-arguments\t0.260s\n"}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.TestFailure))
		Expect(f.Diagnostics).To(BeEmpty())
		Expect(f.Output).To(Equal("=== RUN   TestX\n    main_test.go:17: 2 is bigger than 1\n--- FAIL: TestX (0.00s)"))
	})
	It("reports a crash of a compiled program as a runtime failure", func() {
		result := exercises.Result{Exercise: compile, Err: "panic: boom\n\ngoroutine 1 [running]:\nexit status 2\n"}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.RuntimeFailure))
		Expect(f.Output).To(ContainSubstring("panic: boom"))
	})
})
