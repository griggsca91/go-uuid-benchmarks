# UUID Benchmarks

When creating UUID's with googles UUID lib at a high concurrency, it can get really bogged down from getting a random number.

```
goos: darwin
goarch: arm64
pkg: tracing
cpu: Apple M1 Pro
BenchmarkNewUUID-8                                       3960430               292.6 ns/op
BenchmarkNewUUIDParallel-8                               1240293               946.2 ns/op
BenchmarkNewUUIDWithRandPool-8                          29754400                39.97 ns/op
BenchmarkNewUUIDParallelWithRandPool-8                   8101573               146.8 ns/op
BenchmarkNewUUIDWithNewRand-8                           12160770                99.17 ns/op
BenchmarkNewUUIDParallelWithNewRand-8                   60879198                21.77 ns/op
BenchmarkNewUUIDWithNewRandWithRandPool-8               11870110               101.4 ns/op
BenchmarkNewUUIDParallelWithNewRandWithRandPool-8        7486146               159.2 ns/op
PASS
ok      tracing 11.339s
```

Goal is to just explore if we can get a faster UUID with different ways and configurations

Fastest way I have right now is creating a new rander that just calls rand.Uint and fills the byte array with it. It's not crypto secure, but I'm going for perf right now.

Also using a randpool seems to slow it down, most likely because of the introduction of mutexes protecting the pool.
