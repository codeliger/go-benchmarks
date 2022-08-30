# Go Pointer Benchmarks

```
goos: linux
goarch: amd64
pkg: github.com/codeliger/go-benchmarks
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkSlicePointerPush-24                    162452185                7.086 ns/op          44 B/op          0 allocs/op
BenchmarkSlicePush-24                           214884709                5.769 ns/op          41 B/op          0 allocs/op
BenchmarkSlicePointerCapacityPush-24            609691533                1.899 ns/op           0 B/op          0 allocs/op
BenchmarkSliceCapacityPush-24                   1000000000               0.7002 ns/op          0 B/op          0 allocs/op
BenchmarkSliceSizeSet-24                        1000000000               0.6895 ns/op          0 B/op          0 allocs/op
BenchmarkSlicePointerSet-24                     1000000000               0.7383 ns/op          0 B/op          0 allocs/op
```
