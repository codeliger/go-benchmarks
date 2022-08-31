# Go Pointer Benchmarks

```
goos: linux
goarch: amd64
pkg: github.com/codeliger/go-benchmarks
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkSlicePush-24                      	227928574	         6.673 ns/op	      49 B/op	       0 allocs/op
BenchmarkSlicePointerPush-24               	179129089	         6.506 ns/op	      40 B/op	       0 allocs/op
BenchmarkSliceCapacityPush-24              	819268741	         1.480 ns/op	       8 B/op	       0 allocs/op
BenchmarkSlicePointerCapacityPush-24       	462015600	         2.639 ns/op	       8 B/op	       0 allocs/op
BenchmarkSlicePointerPreallocatedSet-24    	749215220	         1.561 ns/op	       8 B/op	       0 allocs/op
BenchmarkSlicePreallocatedSet-24           	750478893	         1.467 ns/op	       8 B/op	       0 allocs/op
BenchmarkPreallocateSlice-24               	1000000000	         0.8004 ns/op	       8 B/op	       0 allocs/op
BenchmarkPreallocateSliceCapacity-24       	1000000000	         0.7391 ns/op	       8 B/op	       0 allocs/op
PASS
ok  	github.com/codeliger/go-benchmarks	12.149s
```
