package main

import (
	"testing"

	"github.com/google/uuid"
)

var sink uuid.UUID

func BenchmarkNewUUID(b *testing.B) {
	for b.Loop() {
		sink = uuid.New()
	}
}

func BenchmarkNewUUIDParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sink := uuid.New()
			_ = sink
		}
	})
}

func BenchmarkNewUUIDWithRandPool(b *testing.B) {
	uuid.EnableRandPool()
	defer uuid.DisableRandPool()
	for b.Loop() {
		sink = uuid.New()
	}
}

func BenchmarkNewUUIDParallelWithRandPool(b *testing.B) {
	uuid.EnableRandPool()
	defer uuid.DisableRandPool()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sink := uuid.New()
			_ = sink
		}
	})
}

func BenchmarkNewUUIDWithNewRand(b *testing.B) {
	uuid.SetRand(plainRander{})
	for b.Loop() {
		sink = uuid.New()
	}
}

func BenchmarkNewUUIDParallelWithNewRand(b *testing.B) {
	uuid.SetRand(plainRander{})
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sink := uuid.New()
			_ = sink
		}
	})
}

func BenchmarkNewUUIDWithNew16BytesRand(b *testing.B) {
	uuid.SetRand(sixteenBytesRander{})
	for b.Loop() {
		sink = uuid.New()
	}
}

func BenchmarkNewUUIDParallelWithNew16BytesRand(b *testing.B) {
	uuid.SetRand(sixteenBytesRander{})
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sink := uuid.New()
			_ = sink
		}
	})
}

func BenchmarkNewUUIDWithNewRandWithRandPool(b *testing.B) {
	uuid.SetRand(plainRander{})
	uuid.EnableRandPool()
	defer uuid.DisableRandPool()
	for b.Loop() {
		sink = uuid.New()
	}
}

func BenchmarkNewUUIDParallelWithNewRandWithRandPool(b *testing.B) {
	uuid.SetRand(plainRander{})
	uuid.EnableRandPool()
	defer uuid.DisableRandPool()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sink := uuid.New()
			_ = sink
		}
	})
}
