# UUID Benchmarks

When creating UUID's with googles UUID lib at a high concurrency, it can get really bogged down from getting a random number.

```
goos: darwin
goarch: arm64
pkg: tracing
cpu: Apple M1 Pro
BenchmarkNewUUID-8                                       4080936               292.0 ns/op            16 B/op          1 allocs/op
BenchmarkNewUUIDParallel-8                               1237551              1032 ns/op              16 B/op          1 allocs/op
BenchmarkNewUUIDWithRandPool-8                          30328686                40.51 ns/op            0 B/op          0 allocs/op
BenchmarkNewUUIDParallelWithRandPool-8                   8599550               142.5 ns/op             0 B/op          0 allocs/op
BenchmarkNewUUIDWithNewRand-8                           12024308                98.65 ns/op           16 B/op          1 allocs/op
BenchmarkNewUUIDParallelWithNewRand-8                   61836685                19.05 ns/op           16 B/op          1 allocs/op
BenchmarkNewUUIDWithNew16BytesRand-8                    37796614                31.29 ns/op           16 B/op          1 allocs/op
BenchmarkNewUUIDParallelWithNew16BytesRand-8            138289387                9.876 ns/op          16 B/op          1 allocs/op
BenchmarkNewUUIDWithNewRandWithRandPool-8               11863656               100.2 ns/op             0 B/op          0 allocs/op
BenchmarkNewUUIDParallelWithNewRandWithRandPool-8        7615442               163.9 ns/op             0 B/op          0 allocs/op
PASS
ok      tracing 15.669s
```

Goal is to just explore if we can get a faster UUID with googles UUID

Why would you care? UUID's are generated for a ton of reasons, especially for tracing/correlation, etc.

Fastest way I have right now is creating a new rander that just calls rand.Uint and fills the byte array with it. It's not crypto secure, but I'm going for perf right now.

Also using a randpool seems to slow it down, most likely because of the introduction of mutexes protecting the pool.

# References

<https://github.com/open-telemetry/opentelemetry-go/blob/main/sdk/trace/id_generator.go#L52>
