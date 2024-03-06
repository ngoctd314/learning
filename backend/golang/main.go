package main

import (
	"testing"
)

func main() {
	n := 800
	var s []byte
	fg := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s = make([]byte, n)
		}
	})
	println(fg.AllocedBytesPerOp())
	_ = s
}
