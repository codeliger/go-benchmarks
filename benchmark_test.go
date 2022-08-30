package main

import (
	"testing"
)

type Slice []int

func (slice *Slice) PushPointer(i int) {
	*slice = append(*slice, i)
}

func (slice Slice) Push(i int) []int {
	return append(slice, i)
}

func BenchmarkSlicePointerPush(b *testing.B) {
	slice := Slice{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice.PushPointer(i)
	}
}

func BenchmarkSlicePush(b *testing.B) {
	slice := Slice{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice = slice.Push(i)
	}
}

func BenchmarkSlicePointerPreallocatePush(b *testing.B) {
	slice := make(Slice, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice.PushPointer(i)
	}
}

func BenchmarkSlicePreallocatePush(b *testing.B) {
	slice := make(Slice, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice = slice.Push(i)
	}
}
