# UUID Benchmarks

When creating UUID's with googles UUID lib at a high concurrency, it can get really bogged down from getting a random number.

```
goos: darwin
goarch: arm64
pkg: tracing
cpu: Apple M1 Pro
BenchmarkNewUUID-8                               3882837               295.1 ns/op
BenchmarkNewUUIDParallel-8                       1271830               933.5 ns/op
BenchmarkNewUUIDWithRandPool-8                  27917553                48.10 ns/op
BenchmarkNewUUIDParallelWithRandPool-8           8146291               145.1 ns/op
BenchmarkNewUUIDWithNewRand-8                   11691253               102.5 ns/op
BenchmarkNewUUIDParallelWithNewRand-8           54229219                23.61 ns/op
PASS
ok      tracing 8.653s
```

Goal is to just explore if we can get a faster UUID with different ways and configurations

Fastest way I have right now is creating a new rander that just calls rand.Uint and fills the byte array with it. It's not crypto secure, but I'm going for perf right now.
