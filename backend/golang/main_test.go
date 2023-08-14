package main

import (
	"sort"
	"testing"
)

func Benchmark_SortString(b *testing.B) {
	s := []string{"heart", "lungs", "brain", "kidneys", "pancreas"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sort.Strings(s)
	}
}

func BenchmarkSortStrings(b *testing.B) {
	s := []string{"heart", "lungs", "brain", "kidneys", "pancreas"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// var ss sort.StringSlice = s
		// var si sort.Interface = ss // allocation
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	}
}

func Test_fibo(t *testing.T) {
	fibo(5)
}
