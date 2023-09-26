package main

import "testing"

func BenchmarkAllocMapV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		allocMapV1(100000)
	}
}

func BenchmarkAllocMapV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		allocMapV2(100000)
	}
}
