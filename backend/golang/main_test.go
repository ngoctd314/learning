package main

import (
	"sync/atomic"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func BenchmarkAtomicStoreInt64(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&v, 1)
	}
}

func BenchmarkAtomicStoreInt32(b *testing.B) {
	var v int32
	for i := 0; i < b.N; i++ {
		atomic.StoreInt32(&v, 1)
	}
}

var global uint64

func Benchmark_popcnt(b *testing.B) {
	var v uint64
	for i := 0; i < b.N; i++ {
		v = popcnt(11)
	}
	global = v
}
