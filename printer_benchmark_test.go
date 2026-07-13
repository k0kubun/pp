package pp

import "testing"

var benchmarkSprintOutput string

type benchmarkNestedNode struct {
	Value int
	Next  *benchmarkNestedNode
}

func BenchmarkSprintShapes(b *testing.B) {
	benchmarks := []struct {
		name  string
		value interface{}
	}{
		{
			name:  "deep_256",
			value: newBenchmarkNestedNode(256),
		},
		{
			name:  "wide_scalars_512",
			value: make([]int, 512),
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			printer := New()
			printer.SetColoringEnabled(false)
			b.ReportAllocs()
			for b.Loop() {
				benchmarkSprintOutput = printer.Sprint(benchmark.value)
			}
		})
	}
}

func newBenchmarkNestedNode(depth int) *benchmarkNestedNode {
	var node *benchmarkNestedNode
	for range depth {
		node = &benchmarkNestedNode{
			Value: depth,
			Next:  node,
		}
	}
	return node
}
