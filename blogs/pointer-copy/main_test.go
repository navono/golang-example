package pointer_copy

import (
	"fmt"
	"testing"
)

// go get -u golang.org/x/tools/cmd/benchcmp

// go test ./... -bench=BenchmarkMemoryHeap -trace=heapTrace.out -benchmem -run=^$ -count=10
// go test ./... -bench=BenchmarkMemoryStack -trace=stackTrace.out -benchmem -run=^$ -count=10

func BenchmarkMemoryStack(b *testing.B) {
	var s S

	for i := 0; i < b.N; i++ {
		s = byCopy()
	}

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryStackWithFn(b *testing.B) {
	var s S
	var s1 S

	s = byCopy()
	s1 = byCopy()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			s.stack(s1)
		}
	}
}

func BenchmarkMemoryHeap(b *testing.B) {
	var s *S

	for i := 0; i < b.N; i++ {
		s = byPointer()
	}

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryHeapWithFn(b *testing.B) {
	var s *S
	var s1 *S

	s = byPointer()
	s1 = byPointer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			s.heap(s1)
		}
	}
}
