package main

import "testing"

func Benchmark_concat(b *testing.B) {
	s := []string{
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
	}
	for i := 0; i < b.N; i++ {
		concat(s...)
	}
}

func Benchmark_concat1(b *testing.B) {
	s := []string{
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
	}
	for i := 0; i < b.N; i++ {
		concat1(s...)
	}
}

func Benchmark_concat2(b *testing.B) {
	s := []string{
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
		"aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh", "aabcdefghabcdefghabcdefghbcdefgh",
	}
	for i := 0; i < b.N; i++ {
		concat2(s...)
	}
}
