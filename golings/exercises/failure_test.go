package exercises_test

import (
	"os"
	"path/filepath"
	"time"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Interpreting failures", func() {
	compile := exercises.Exercise{Name: "e", Mode: "compile"}
	test := exercises.Exercise{Name: "e", Mode: "test"}

	It("extracts compiler diagnostics and drops the package header", func() {
		result := exercises.Result{Exercise: compile, BuildFailed: true, Err: "# command-line-arguments\nexercises/a/main.go:10:6: syntax error: unexpected =, expected name\nexercises/a/main.go:12:2: undefined: x\n"}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.CompileFailure))
		Expect(f.Diagnostics).To(Equal([]exercises.Diagnostic{
			{File: "exercises/a/main.go", Line: 10, Col: 6, Message: "syntax error: unexpected =, expected name"},
			{File: "exercises/a/main.go", Line: 12, Col: 2, Message: "undefined: x"},
		}))
		Expect(f.Output).To(BeEmpty())
	})
	It("keeps messages that are not diagnostics", func() {
		result := exercises.Result{Exercise: compile, BuildFailed: true, Err: "# command-line-arguments\na/main.go:3:1: undefined: y\na/main.go:3:1: too many errors\n"}

		Expect(result.Failure().Diagnostics).To(HaveLen(2))

		result = exercises.Result{Exercise: compile, BuildFailed: true, Err: "a/main.go:3:1: undefined: y\ntoo many errors\n"}

		f := result.Failure()
		Expect(f.Diagnostics).To(HaveLen(1))
		Expect(f.Output).To(Equal("too many errors"))
	})
	It("understands diagnostics without a column", func() {
		result := exercises.Result{Exercise: compile, BuildFailed: true, Err: "vet: a/main.go:7: oops\n"}

		Expect(result.Failure().Diagnostics).To(ConsistOf(
			exercises.Diagnostic{File: "a/main.go", Line: 7, Message: "oops"},
		))
	})
	It("reports a compile failure of a test-mode exercise as such", func() {
		result := exercises.Result{Exercise: test, BuildFailed: true, Err: "# command-line-arguments\na/main_test.go:5:2: undefined: Foo\n"}

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
	It("reports a timeout and keeps what the program printed before it", func() {
		result := exercises.Result{Exercise: compile, TimedOut: true, Timeout: 5 * time.Second, Out: "started\n"}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.TimeoutFailure))
		Expect(f.Timeout).To(Equal(5 * time.Second))
		Expect(f.Output).To(Equal("started"))
	})
	It("reports a build failure without diagnostics as a compile failure", func() {
		result := exercises.Result{Exercise: compile, BuildFailed: true, Err: "go: cannot find main module\n"}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.CompileFailure))
		Expect(f.Output).To(ContainSubstring("cannot find main module"))
	})
	It("reports a crash of a compiled program as a runtime failure", func() {
		result := exercises.Result{Exercise: compile, Err: "panic: boom\n\ngoroutine 1 [running]:\nexit status 2\n"}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.RuntimeFailure))
		Expect(f.Output).To(ContainSubstring("panic: boom"))
	})
})

var _ = Describe("Condensing panics", func() {
	test := exercises.Exercise{Name: "e", Mode: "test"}
	cwd, _ := os.Getwd()
	file := filepath.Join(cwd, "exercises", "pointers", "pointers4", "main_test.go")

	nilDereference := "panic: runtime error: invalid memory address or nil pointer dereference [recovered, repanicked]\n" +
		"[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x100689738]\n" +
		"\n" +
		"goroutine 8 [running]:\n" +
		"testing.tRunner.func1.2({0x10081d0a8, 0x10085f6b0})\n" +
		"\t/usr/local/go/src/testing/testing.go:2125 +0x2b8\n" +
		"testing.tRunner.func1()\n" +
		"\t/usr/local/go/src/testing/testing.go:2128 +0x438\n" +
		"panic({0x10081d0a8?, 0x10085f6b0?})\n" +
		"\t/usr/local/go/src/runtime/panic.go:859 +0x120\n" +
		"command-line-arguments_test.greeting(...)\n" +
		"\t" + file + ":15\n" +
		"command-line-arguments_test.TestGreetingWithoutAPerson(0xc00013a488)\n" +
		"\t" + file + ":27 +0x28\n" +
		"testing.tRunner(0xc00013a488, 0x10083b648)\n" +
		"\t/usr/local/go/src/testing/testing.go:2195 +0x168\n" +
		"created by testing.(*T).Run in goroutine 1\n" +
		"\t/usr/local/go/src/testing/testing.go:2260 +0x7c0\n"

	It("keeps the message and the frames in the learner's code", func() {
		result := exercises.Result{Exercise: test, Out: "=== RUN   TestGreeting\n--- PASS: TestGreeting (0.00s)\n=== RUN   TestGreetingWithoutAPerson\n--- FAIL: TestGreetingWithoutAPerson (0.00s)\nFAIL\n", Err: nilDereference}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.TestFailure))
		Expect(f.Output).To(Equal("=== RUN   TestGreeting\n" +
			"--- PASS: TestGreeting (0.00s)\n" +
			"=== RUN   TestGreetingWithoutAPerson\n" +
			"--- FAIL: TestGreetingWithoutAPerson (0.00s)\n" +
			"\n" +
			"panic: runtime error: invalid memory address or nil pointer dereference\n" +
			"\n" +
			"  at greeting (exercises/pointers/pointers4/main_test.go:15)\n" +
			"  at TestGreetingWithoutAPerson (exercises/pointers/pointers4/main_test.go:27)"))
	})
	It("leaves output without a panic alone", func() {
		result := exercises.Result{Exercise: test, Out: "=== RUN   TestX\n    main_test.go:17: 2 is bigger than 1\n--- FAIL: TestX (0.00s)\nFAIL\n"}

		Expect(result.Failure().Output).To(Equal("=== RUN   TestX\n    main_test.go:17: 2 is bigger than 1\n--- FAIL: TestX (0.00s)"))
	})
	It("condenses the stack of a deadlock too", func() {
		deadlock := "fatal error: all goroutines are asleep - deadlock!\n\n" +
			"goroutine 1 [chan receive]:\n" +
			"main.main()\n" +
			"\t" + filepath.Join(cwd, "exercises", "a", "main.go") + ":9 +0x3c\n"
		result := exercises.Result{Exercise: exercises.Exercise{Mode: "compile"}, Err: deadlock}

		f := result.Failure()

		Expect(f.Kind).To(Equal(exercises.RuntimeFailure))
		Expect(f.Output).To(Equal("fatal error: all goroutines are asleep - deadlock!\n\n  at main (exercises/a/main.go:9)"))
	})
})
