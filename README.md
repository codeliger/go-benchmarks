# Go Pointer Benchmarks

```
goos: linux
goarch: amd64
pkg: github.com/codeliger/go-benchmarks
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkSlicePush-24                           186407437                6.792 ns/op          48 B/op          0 allocs/op
BenchmarkSlicePointerPush-24                    168715802                7.215 ns/op          42 B/op          0 allocs/op
BenchmarkSliceCapacityPush-24                   819572942                1.467 ns/op           8 B/op          0 allocs/op
BenchmarkSlicePointerCapacityPush-24            447383870                2.509 ns/op           8 B/op          0 allocs/op
BenchmarkSlicePointerPreallocatedSet-24         830871080                1.571 ns/op           8 B/op          0 allocs/op
BenchmarkSlicePreallocatedSet-24                767708407                1.483 ns/op           8 B/op          0 allocs/op
PASS
ok      github.com/codeliger/go-benchmarks      9.404s
```
