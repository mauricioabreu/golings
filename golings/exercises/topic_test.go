package exercises_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// writeTopicWorkspace creates exercises laid out as exercises/<topic>/<name>/main.go.
// Each spec is "topic/name" and a leading "!" leaves the exercise unsolved.
func writeTopicWorkspace(specs ...string) string {
	root := GinkgoT().TempDir()
	var info strings.Builder
	for _, spec := range specs {
		solved := !strings.HasPrefix(spec, "!")
		spec = strings.TrimPrefix(spec, "!")
		path := filepath.Join(root, "exercises", spec, "main.go")
		Expect(os.MkdirAll(filepath.Dir(path), 0o755)).To(Succeed())

		source := "package main\n\nfunc main() {}\n"
		if !solved {
			source = "// I AM NOT DONE\n" + source
		}
		Expect(os.WriteFile(path, []byte(source), 0o644)).To(Succeed())
		fmt.Fprintf(&info, "[[exercises]]\nname = %q\npath = %q\nmode = \"compile\"\nhint = \"\"\n\n", filepath.Base(spec), path)
	}
	file := filepath.Join(root, "info.toml")
	Expect(os.WriteFile(file, []byte(info.String()), 0o644)).To(Succeed())
	return file
}

var _ = Describe("Topics", func() {
	DescribeTable("Topic",
		func(path, want string) {
			Expect(exercises.Exercise{Path: path}.Topic()).To(Equal(want))
		},
		Entry("regular exercise", "exercises/variables/variables1/main.go", "variables"),
		Entry("topic with underscore", "exercises/anonymous_functions/anonymous_functions1/main.go", "anonymous_functions"),
		Entry("absolute path", "/home/me/golings/exercises/slices/slices1/main_test.go", "slices"),
		Entry("path without a topic folder", "compile/compile1.go", ""),
		Entry("exercises folder without a topic", "exercises/main.go", ""),
	)

	DescribeTable("SolutionPath",
		func(path, want string) {
			Expect(exercises.Exercise{Path: path}.SolutionPath()).To(Equal(want))
		},
		Entry("regular exercise", "exercises/slices/slices1/main.go", "solutions/slices/slices1/main.go"),
		Entry("test-mode exercise", "exercises/if/if1/main_test.go", "solutions/if/if1/main_test.go"),
		Entry("workspace below another folder", "../tmp/ws/exercises/a/a1/main.go", "solutions/a/a1/main.go"),
		Entry("no topic folder", "compile/compile1.go", ""),
	)

	Describe("ProgressByTopic", func() {
		It("counts done and total per topic in manifest order", func() {
			info := writeTopicWorkspace("a/a1", "a/a2", "!a/a3", "!b/b1", "c/c1")

			topics, err := exercises.ProgressByTopic(info)

			Expect(err).NotTo(HaveOccurred())
			Expect(topics).To(Equal([]exercises.TopicProgress{
				{Topic: "a", Done: 2, Total: 3},
				{Topic: "b", Done: 0, Total: 1},
				{Topic: "c", Done: 1, Total: 1},
			}))
		})
		It("adds up to the overall progress", func() {
			info := writeTopicWorkspace("a/a1", "!a/a2", "!b/b1")

			topics, err := exercises.ProgressByTopic(info)
			Expect(err).NotTo(HaveOccurred())
			_, done, total, err := exercises.Progress(info)
			Expect(err).NotTo(HaveOccurred())

			sumDone, sumTotal := 0, 0
			for _, topic := range topics {
				sumDone += topic.Done
				sumTotal += topic.Total
			}
			Expect([]int{sumDone, sumTotal}).To(Equal([]int{done, total}))
		})
		It("leaves out exercises the local Go can't run", func() {
			topics, err := exercises.ProgressByTopic("../fixtures/goversion/pending/info.toml")

			Expect(err).NotTo(HaveOccurred())
			Expect(topics).To(ConsistOf(exercises.TopicProgress{Topic: "", Done: 1, Total: 1}))
		})
		It("returns an error when the info file can't be read", func() {
			_, err := exercises.ProgressByTopic("../fixtures/info404.toml")

			Expect(err).To(HaveOccurred())
		})
		It("covers every exercise of the real manifest", func() {
			topics, err := exercises.ProgressByTopic("../../info.toml")
			Expect(err).NotTo(HaveOccurred())
			all, err := exercises.List("../../info.toml")
			Expect(err).NotTo(HaveOccurred())

			total := 0
			for _, topic := range topics {
				Expect(topic.Topic).NotTo(BeEmpty())
				total += topic.Total
			}
			Expect(total).To(Equal(len(all)))
			Expect(topics[0].Topic).To(Equal("variables"))
		})
	})
})
