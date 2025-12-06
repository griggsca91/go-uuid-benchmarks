package main

import (
	"math/rand/v2"
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

type plainRander struct{}

func (plainRander) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = byte(rand.Uint())
	}
	return len(b), nil
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
