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

func Benchmark_linkedlist(b *testing.B) {
	n := &node{
		next: &node{
			next: &node{
				next: &node{
					next: &node{
						value: 5,
					},
					value: 4,
				},
				value: 3,
			},
			value: 2,
		},
		value: 1,
	}
	for i := 0; i < b.N; i++ {
		linkedlist(n)
	}
}
func Benchmark_sum2(b *testing.B) {
	n := []int64{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		sum2(n)
	}
}

// Benchmark_hitCache-12    	20667152	        66.39 ns/op	       0 B/op	       0 allocs/op
// Benchmark_withoutHitCache-12    	16307011	        70.10 ns/op	       0 B/op	       0 allocs/op
