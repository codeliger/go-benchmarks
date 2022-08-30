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

func (slice *Slice) Set(i int) {
	(*slice)[i] = i
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

func BenchmarkSlicePointerCapacityPush(b *testing.B) {
	slice := make(Slice, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice.PushPointer(i)
	}
}

func BenchmarkSliceCapacityPush(b *testing.B) {
	slice := make(Slice, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice = slice.Push(i)
	}
}

func BenchmarkSliceSizeSet(b *testing.B) {
	slice := make(Slice, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice[i] = i
	}
}

func BenchmarkSlicePointerSet(b *testing.B) {
	slice := make(Slice, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice.Set(i)
	}
}
