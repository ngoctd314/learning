package main

import "testing"

func Benchmark_assignOnly1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assignOnly1()
	}
}

func Benchmark_assignOnly2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assignOnly10()
	}
}
