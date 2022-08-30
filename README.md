# Go Pointer Benchmarks

goos: linux
goarch: amd64
pkg: github.com/codeliger/go-benchmarks
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkSlicePointerPush-24               	150355866	         7.669 ns/op
BenchmarkSlicePush-24                      	202723568	         6.186 ns/op
BenchmarkSlicePointerPreallocatePush-24    	532548985	         1.961 ns/op
BenchmarkSlicePreallocatePush-24           	1000000000	         0.7064 ns/op
PASS
ok  	github.com/codeliger/go-benchmarks	7.195s
