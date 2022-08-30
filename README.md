# Go Pointer Benchmarks

```
goos: linux
goarch: amd64
pkg: github.com/codeliger/go-benchmarks
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkSlicePush-24                           186898304                5.975 ns/op          47 B/op          0 allocs/op
BenchmarkSlicePointerPush-24                    166452036                6.660 ns/op          43 B/op          0 allocs/op
BenchmarkSliceCapacityPush-24                   1000000000               0.7106 ns/op          0 B/op          0 allocs/op
BenchmarkSlicePointerCapacityPush-24            538413109                1.989 ns/op           0 B/op          0 allocs/op
BenchmarkSlicePointerPreallocatedSet-24         1000000000               0.7260 ns/op          0 B/op          0 allocs/op
BenchmarkSlicePreallocatedSet-24                1000000000               0.6903 ns/op          0 B/op          0 allocs/op
```
