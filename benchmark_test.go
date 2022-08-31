package main

import (
	"testing"
)

type Slice []int

func (slice Slice) Push(i int) []int {
	return append(slice, i)
}

func (slice *Slice) PushPointer(i int) {
	*slice = append(*slice, i)
}

func (slice *Slice) Set(i int) {
	(*slice)[i] = i
}

func BenchmarkSlicePush(b *testing.B) {
	slice := Slice{}
	for i := 0; i < b.N; i++ {
		slice = slice.Push(i)
	}
}

func BenchmarkSlicePointerPush(b *testing.B) {
	slice := Slice{}
	for i := 0; i < b.N; i++ {
		slice.PushPointer(i)
	}
}

func BenchmarkSliceCapacityPush(b *testing.B) {
	slice := make(Slice, 0, b.N)
	for i := 0; i < b.N; i++ {
		slice = slice.Push(i)
	}
}

func BenchmarkSlicePointerCapacityPush(b *testing.B) {
	slice := make(Slice, 0, b.N)
	for i := 0; i < b.N; i++ {
		slice.PushPointer(i)
	}
}

func BenchmarkSlicePointerPreallocatedSet(b *testing.B) {
	slice := make(Slice, b.N)
	for i := 0; i < b.N; i++ {
		slice.Set(i)
	}
}

func BenchmarkSlicePreallocatedSet(b *testing.B) {
	slice := make(Slice, b.N)
	for i := 0; i < b.N; i++ {
		slice[i] = i
	}
}
