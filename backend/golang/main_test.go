package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// func Benchmark_hitCache(b *testing.B) {
// 	var s []int64
// 	var i int64
// 	for i = 0; i < 1000; i++ {
// 		s = append(s, i)
// 	}

// 	for i := 0; i < b.N; i++ {
// 		sum2(s)
// 	}
// }

// func Benchmark_withoutHitCache(b *testing.B) {
// 	var s []int64
// 	var i int64
// 	for i = 0; i < 1000; i++ {
// 		s = append(s, i)
// 	}

// 	for i := 0; i < b.N; i++ {
// 		sum8(s)
// 	}
// }

func Benchmark_sumOfSliceStruct(b *testing.B) {
	foos := make([]Foo, 1000)
	var i int64
	for i = 0; i < 1000; i++ {
		foos[i] = Foo{
			a: i,
			b: i,
		}
	}

	for i := 0; i < b.N; i++ {
		sumFoo(foos)
	}
}

func Benchmark_sumOfStructSlice(b *testing.B) {
	a := make([]int64, 1000)
	a1 := make([]int64, 1000)
	var i int64
	for i = 0; i < 1000; i++ {
		a[i] = i
		a1[i] = i
	}
	bar := Bar{
		a: a,
		b: a1,
	}

	for i := 0; i < b.N; i++ {
		sumBar(bar)
	}
}

// Benchmark_hitCache-12    	20667152	        66.39 ns/op	       0 B/op	       0 allocs/op
// Benchmark_withoutHitCache-12    	16307011	        70.10 ns/op	       0 B/op	       0 allocs/op
