package bench

import (
	"gempellm/playground/internal/async_cache/mcache"
	"gempellm/playground/internal/async_cache/rwcache"
	"testing"
)

var parallelFactor = 100_000

func Benchmark_RWMutex_BalancedLoad(b *testing.B) {
	c := rwcache.NewRWCache()
	for i := 0; i < b.N; i++ {
		emulateLoad(b, c, parallelFactor)
	}
}

func Benchmark_Mutex_BalancedLoad(b *testing.B) {
	c := mcache.NewMCache()
	for i := 0; i < b.N; i++ {
		emulateLoad(b, c, parallelFactor)
	}
}

func Benchmark_RWMutex_IntenseRead(b *testing.B) {
	c := rwcache.NewRWCache()
	for i := 0; i < b.N; i++ {
		emulateLoadIntenseRead(b, c, parallelFactor)
	}
}

func Benchmark_Mutex_IntenseRead(b *testing.B) {
	c := mcache.NewMCache()
	for i := 0; i < b.N; i++ {
		emulateLoadIntenseRead(b, c, parallelFactor)
	}
}
