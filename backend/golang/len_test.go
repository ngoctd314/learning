package main

import "testing"

func Benchmark_forlen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		forlen1()
	}
}

func Benchmark_forlen2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		forlen2()
	}
}
