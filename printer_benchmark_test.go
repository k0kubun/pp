package pp

import (
	"fmt"
	"testing"
)

type benchmarkNode struct {
	Value int
	Empty string
	Child *benchmarkNode
}

var benchmarkOutput string

func BenchmarkSprintNested(b *testing.B) {
	for _, depth := range []int{1, 75, 256} {
		value := newBenchmarkNode(depth)
		b.Run(fmt.Sprintf("depth=%d", depth), func(b *testing.B) {
			printer := New()
			printer.SetColoringEnabled(true)
			printer.SetOmitEmpty(true)
			printer.SetMaxDepth(256)
			b.ReportAllocs()

			for b.Loop() {
				benchmarkOutput = printer.Sprint(value)
			}
		})
	}
}

func newBenchmarkNode(depth int) *benchmarkNode {
	var child *benchmarkNode
	for value := depth; value > 0; value-- {
		child = &benchmarkNode{
			Value: value,
			Child: child,
		}
	}
	return child
}
