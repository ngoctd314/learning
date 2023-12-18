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

func Benchmark_count1(b *testing.B) {
	var n int64 = 1000000
	var inputs []Input

	inputs = make([]Input, 0, n)
	for i := int64(0); i < n; i++ {
		inputs = append(inputs, Input{
			a: 1,
			b: 1,
		})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count1(inputs)
	}
}

func Benchmark_count2(b *testing.B) {
	var n int64 = 1000000
	var inputs []Input

	inputs = make([]Input, 0, n)
	for i := int64(0); i < n; i++ {
		inputs = append(inputs, Input{
			a: 1,
			b: 1,
		})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count2(inputs)
	}
}
func Benchmark_count3(b *testing.B) {
	var n int64 = 1000000
	var inputs []Input

	inputs = make([]Input, 0, n)
	for i := int64(0); i < n; i++ {
		inputs = append(inputs, Input{
			a: 1,
			b: 1,
		})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count3(inputs)
	}
}
