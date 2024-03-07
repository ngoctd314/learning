package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var src = []int{1, 2, 3, 4, 5}

func Benchmark_Fn1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst := make([]int, len(src))
		copy(dst, src)
	}
}

func Benchmark_Fn2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst := make([]int, len(src))
		_ = copy(dst, src)
	}
}
